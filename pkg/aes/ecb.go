package aes

import (
	"fmt"
	"maps"
	"slices"
)

// Recommendation for Block Cipher Modes of Operation: Methods and Techniques
// https://nvlpubs.nist.gov/nistpubs/Legacy/SP/nistspecialpublication800-38a.pdf

func DecryptAESECB(ciphertext []byte, key []byte) ([]byte, error) {
	blocks := slices.Chunk(ciphertext, AESBlockSize)
	var plaintext []byte
	for block := range blocks {
		pt, err := InvCipher(block, key)
		if err != nil {
			return nil, fmt.Errorf("this block could not be decrypted: %#v", block)
		}
		plaintext = append(plaintext, pt...)
	}
	return plaintext, nil
}

func DetectRedundantAESBlocks(ciphertext []byte) (int, error) {
	freqBlocks := make(map[string]int)
	blocks := slices.Chunk(ciphertext, AESBlockSize)
	for block := range blocks {
		freqBlocks[string(block)]++
	}
	val := slices.Sorted(maps.Values(freqBlocks))[len(freqBlocks)-1]
	if val == 1 {
		return 0, nil
	}
	return val, nil
}
