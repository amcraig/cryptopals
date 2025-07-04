package aes

import (
	"fmt"
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
