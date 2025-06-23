package xor

import (
	"fmt"
	"slices"
	"sort"

	"github.com/amcraig/cryptopals-go/internal/common"
)

const NumBlocksToAverage = 4
const PickNKeysizes = 3

func RepeatingKeyXOREncode(input []byte, key []byte) ([]byte, error) {
	lenKey := len(key)
	if lenKey == 0 {
		return nil, fmt.Errorf("encoding key provided is of length: 0")
	}

	buf := make([]byte, len(input))
	idxKey := 0

	for i := range input {
		buf[i] = input[i] ^ key[idxKey]
		// Roll key to next byte, wrap at the end
		idxKey = (idxKey + 1) % lenKey
	}

	return buf, nil
}

func RepeatingKeyXORDecode(input []byte, key []byte) ([]byte, error) {
	buf, err := RepeatingKeyXOREncode(input, key)
	if err != nil {
		return []byte{}, fmt.Errorf("something went wrong xor decoding")
	}
	return buf, nil
}

func BreakRepeatingKeyXOR(buf []byte, maxKeysize int) ([]byte, error) {
	bestKeysizes, _ := guessKeysize(buf, maxKeysize)
	var solvedKey []byte
	bestScore := 0.

	for _, keysize := range bestKeysizes {
		blocks, _ := chunkTextByKeyIndex(buf, keysize)

		// Solve each block as if it was single-character XOR. You already have code to do this.
		// For each block, the single-byte XOR key that produces the best looking histogram is the repeating-key XOR key byte for that block. Put them together and you have the key.
		var key []byte
		for _, tb := range blocks {
			// Decode the block for the single pos of the key, do not care about spaces
			_, asciiByte, _, _ := SingleByteXORDecode(tb, -1)
			key = append(key, asciiByte)
		}

		// Judge to see if the string is human-writable
		score, _ := common.ScoreEnglishPhrase(key, -1)
		if score > bestScore {
			bestScore = score
			solvedKey = key
		}
	}
	return solvedKey, nil
}

// The KEYSIZE with the smallest normalized edit distance is probably the key.
// You could proceed perhaps with the smallest 2-3 KEYSIZE values.
// Or take 4 KEYSIZE blocks instead of 2 and average the distances.
func guessKeysize(buf []byte, maxKeysize int) ([PickNKeysizes]int, error) {
	type KeysizeResult struct {
		dist    float64
		keysize int
	}

	var keysizeResults []KeysizeResult
	var results [PickNKeysizes]int

	// For each KEYSIZE, take the first KEYSIZE worth of bytes, and the second KEYSIZE worth of bytes, and find the edit distance between them. Normalize this result by dividing by KEYSIZE.
	for keysize := 2; keysize < maxKeysize+1; keysize++ {

		avgDist := 0.

		firstBite := buf[:keysize]
		for biteN := 2; biteN <= NumBlocksToAverage; biteN++ {
			nBite := buf[(biteN-1)*keysize : biteN*keysize]
			nDist, err := common.HammingDistance(firstBite, nBite)
			if err != nil {
				return [PickNKeysizes]int{}, fmt.Errorf("hamming distance could not be calculated for: %#v %#v", firstBite, nBite)
			}
			avgDist += float64(nDist)
		}
		avgDist /= float64(NumBlocksToAverage)
		normDist := avgDist / float64(keysize)

		keysizeResults = append(keysizeResults, KeysizeResult{normDist, keysize})
	}

	sort.Slice(keysizeResults, func(i, j int) bool {
		return keysizeResults[i].dist < keysizeResults[j].dist
	})

	for i := range PickNKeysizes {
		results[i] = keysizeResults[i].keysize
	}

	return results, nil
}

func chunkTextByKeyIndex(buf []byte, keysize int) ([][]byte, error) {
	// Now that you probably know the KEYSIZE: break the ciphertext into blocks of KEYSIZE length.
	iterBlocks := slices.Chunk(buf, keysize)

	// Now transpose the blocks: make a block that is the first byte of every block, and a block that is the second byte of every block, and so on.
	// Initializes a 2D slice containing slices of type: []byte
	transBlocks := make([][]byte, keysize)
	// For each keysize chunk
	for ib := range iterBlocks {
		// Loop through each byte and index
		for i, v := range ib {
			// Append to the subslice corresponding to that index of the keysize
			transBlocks[i] = append(transBlocks[i], v)
		}
	}
	return transBlocks, nil
}
