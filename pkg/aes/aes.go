package aes

import (
	"fmt"
	"math/bits"
	"slices"

	"github.com/amcraig/cryptopals-go/internal/constants"
	"github.com/amcraig/cryptopals-go/internal/math/matrix"
)

// Advanced Encryption Standard (AES)
// https://nvlpubs.nist.gov/nistpubs/FIPS/NIST.FIPS.197-upd1.pdf

// word type definition:

type word = []byte

func XORWord(a, b word) word {
	for i := range a {
		a[i] = a[i] ^ b[i]
	}
	return a
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

func (s state) ShiftRows() state {
	for idx, row := range s {
		mod := idx % 4
		s[idx] = matrix.RotateVector(row, mod)
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

func (s state) AddRoundKey(keySchedule []word, i uint8) state {
	roundKey := matrix.Transpose(keySchedule[i*4 : i*4+4])
	for i := range roundKey {
		for j := range roundKey[i] {
			s[i][j] = s[i][j] ^ roundKey[i][j]
		}
	}
	return s
}

// Public

func AES128Cipher(in []byte, key []byte) ([]byte, error) {
	const NumRounds uint8 = 10
	keySchedule, err := keyExpansion(key)
	if err != nil {
		return nil, fmt.Errorf("failed to generate the key schedule: %v", err)
	}

	return Cipher(in, NumRounds, keySchedule), nil
}

func Cipher(in []byte, numRounds uint8, keySchedule []word) []byte {
	state, _ := readBlockIntoState(in)
	state = state.AddRoundKey(keySchedule, 0)
	var round uint8
	for round = 1; round < numRounds; round++ {
		state = state.SubBytes()
		state = state.ShiftRows()
		state = state.MixColumns()
		state = state.AddRoundKey(keySchedule, round)
	}
	state = state.SubBytes()
	state = state.ShiftRows()
	state.AddRoundKey(keySchedule, 10)
	return slices.Concat(state.Words()...)
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
	if numWords < 4 || numWords > 6 {
		panic("the incoming key is does not match: 128, 192, or 256 bits")
	}
	var words []word
	keyChunks := slices.Chunk(key, 4)
	for chunk := range keyChunks {
		words = append(words, chunk)
	}
	return words
}

func keyExpansion(key []byte) ([]word, error) {
	var numRounds int

	switch len(key) {
	default:
		return nil, fmt.Errorf("the incoming key is does not match: 128, 192, or 256 bits")
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
	return keySchedule, nil
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
