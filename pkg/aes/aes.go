package aes

import (
	"fmt"
	"math/bits"
	"slices"

	"github.com/amcraig/cryptopals-go/internal/bytes"
	"github.com/amcraig/cryptopals-go/internal/constants"
	"github.com/amcraig/cryptopals-go/internal/math/matrix"
)

// Advanced Encryption Standard (AES)
// https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.197-upd1.pdf

const AESBlockSize = 16 // Bytes

// word type definition:

type word = []byte

func XORWord(a, b word) word {
	return bytes.XORBytes(a, b)
}

func SubWord(w word) word {
	for i := range w {
		w[i] = constants.SBox[w[i]]
	}
	return w
}

func RotWord(w word, pos int) word {
	return matrix.RotateVector(w, pos)
}

// state type definition

type state [][]byte

func (s state) Words() []word {
	var words []word
	for j := range 4 {
		words = append(words, word{s[0][j], s[1][j], s[2][j], s[3][j]})
	}
	return words
}

func (s *state) SetStateFromWords(w []word) {
	if len(w) != 4 {
		panic("tried to set a non-128 bit state")
	}
	*s, _ = readBlockIntoState([]byte(slices.Concat(w...)))
}

func (s state) SubBytes() state {
	for i := range s {
		for j := range s[i] {
			s[i][j] = constants.SBox[s[i][j]]
		}
	}
	return s
}

func (s state) InvSubBytes() state {
	for i := range s {
		for j := range s[i] {
			s[i][j] = constants.InvSBox[s[i][j]]
		}
	}
	return s
}

func (s state) ShiftRows() state {
	for idx, row := range s {
		mod := idx % 4
		s[idx] = matrix.RotateVector(row, mod)
	}
	return s
}

func (s state) InvShiftRows() state {
	for idx, row := range s {
		mod := idx % 4
		s[idx] = matrix.RotateVector(row, -mod)
	}
	return s
}

func (s state) MixColumns() state {
	t := make([][]byte, 4)
	for wi, word := range s.Words() {
		tempWord := make([]byte, 4)
		for ri, row := range constants.MixColumnMatrix {
			var acc byte
			for bi := range word {
				switch mult := row[bi]; mult {
				case 2:
					acc ^= xtimes(word[bi])
				case 3:
					acc ^= xtimes(word[bi]) ^ word[bi]
				default:
					acc ^= word[bi]
				}
			}
			tempWord[ri] = acc
		}
		t[wi] = tempWord
	}
	return matrix.Transpose(t)
}

func (s state) InvMixColumns() state {
	t := make([][]byte, 4)
	for wi, word := range s.Words() {
		tempWord := make([]byte, 4)
		for ri, row := range constants.InvMixColumnMatrix {
			var acc byte
			for bi := range word {
				switch mult := row[bi]; mult {
				case 0x09:
					// 9 = 8 ^ 1
					acc ^= xtimes(xtimes(xtimes(word[bi]))) ^ word[bi]
				case 0x0b:
					// 11 = 8 ^ 2 ^ 1
					acc ^= xtimes(xtimes(xtimes(word[bi]))) ^ xtimes(word[bi]) ^ word[bi]
				case 0x0d:
					// 13 = 8 ^ 4 ^ 1
					acc ^= xtimes(xtimes(xtimes(word[bi]))) ^ xtimes(xtimes(word[bi])) ^ word[bi]
				case 0x0e:
					// 14 = 8 ^ 4 ^ 2
					acc ^= xtimes(xtimes(xtimes(word[bi]))) ^ xtimes(xtimes(word[bi])) ^ xtimes(word[bi])
				}
			}
			tempWord[ri] = acc
		}
		t[wi] = tempWord
	}
	return matrix.Transpose(t)
}

func (s state) AddRoundKey(rk [][]byte) state {
	for i := range rk {
		for j := range rk[i] {
			s[i][j] = s[i][j] ^ rk[i][j]
		}
	}
	return s
}

// Public

func Cipher(in []byte, key []byte) ([]byte, error) {
	if len(in) != AESBlockSize {
		return nil, fmt.Errorf("input block was not 128 bits")
	}
	ks, numRounds, _ := keyExpansion(key)
	state, _ := readBlockIntoState(in)
	state = state.AddRoundKey(matrix.Transpose(ks[:4]))
	for round := 1; round < numRounds; round++ {
		state = state.SubBytes()
		state = state.ShiftRows()
		state = state.MixColumns()
		state = state.AddRoundKey(matrix.Transpose(ks[4*round : 4*round+4]))
	}
	state = state.SubBytes()
	state = state.ShiftRows()
	state.AddRoundKey(matrix.Transpose(ks[4*numRounds : 4*numRounds+4]))
	return slices.Concat(state.Words()...), nil
}

func InvCipher(ct []byte, key []byte) ([]byte, error) {
	if len(ct) != AESBlockSize {
		return nil, fmt.Errorf("input block was not 128 bits")
	}
	ks, numRounds, _ := keyExpansion(key)
	state, _ := readBlockIntoState(ct)
	state = state.AddRoundKey(matrix.Transpose(ks[4*numRounds:])) // since the schedule is applied via XOR, redoing the operation unapplies
	for round := numRounds - 1; round > 0; round-- {
		state = state.InvShiftRows()
		state = state.InvSubBytes()
		state = state.AddRoundKey(matrix.Transpose(ks[4*round : 4*round+4]))
		state = state.InvMixColumns()
	}
	state = state.InvShiftRows()
	state = state.InvSubBytes()
	state = state.AddRoundKey(matrix.Transpose(ks[:4]))
	return slices.Concat(state.Words()...), nil
}

// Private

func xtimes(b byte) byte {
	if bits.LeadingZeros8(b) == 0 {
		// If byte has a leading bit, then doubling it would cause an overflow
		return (b << 1) ^ 0x1b
	}
	return b << 1 // Doubles byte
}

func readKeyIntoWords(key []byte) []word {
	numWords := len(key) / 4
	if numWords < 4 || numWords > 8 {
		panic("the incoming key is does not match: 128, 192, or 256 bits")
	}
	var words []word
	keyChunks := slices.Chunk(key, 4)
	for chunk := range keyChunks {
		words = append(words, chunk)
	}
	return words
}

func keyExpansion(key []byte) ([]word, int, error) {
	var numRounds int

	switch len(key) {
	default:
		return nil, 0, fmt.Errorf("the incoming key is does not match: 128, 192, or 256 bits")
	case 16:
		numRounds = 10
	case 24:
		numRounds = 12
	case 32:
		numRounds = 14
	}

	keyWords := readKeyIntoWords(key)
	nk := len(keyWords)

	var keySchedule []word
	// The first Nk words of the expanded key are the key itself.
	keySchedule = append(keySchedule, keyWords...)

	//Every subsequent word w[i] is generated recursively from the preceding word,
	// w[i − 1], and the word Nk positions earlier, w[i − Nk]
	for i := nk; i < 4*(numRounds+1); i++ {
		temp := slices.Clone(keySchedule[i-1])
		if i%nk == 0 {
			temp = RotWord(temp, 1)
			temp = SubWord(temp)
			temp = XORWord(temp, word{constants.RoundConstant[i/nk], 0, 0, 0})
		} else if nk > 6 && i%nk == 4 {
			temp = SubWord(temp)
		}
		keySchedule = append(keySchedule, XORWord(slices.Clone(keySchedule[i-nk]), temp))
	}
	return keySchedule, numRounds, nil
}

func readBlockIntoState(block []byte) (state, error) {
	if len(block) > 16 {
		return nil, fmt.Errorf("the incoming block is larger than the size requirements to load into the state: %v", len(block))
	}

	state := make(state, 4)
	for i := range len(state) {
		state[i] = make([]byte, 4)
	}

	for idx, b := range block {
		state[idx%4][idx/4] = b
	}

	return state, nil
}
