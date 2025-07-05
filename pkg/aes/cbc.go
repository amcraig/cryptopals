package aes

import (
	"fmt"
	"slices"

	"github.com/amcraig/cryptopals-go/internal/bytes"
)

func DecryptAESCBC(ciphertext []byte, key []byte, iv []byte) ([]byte, error) {
	if len(iv) != 16 {
		return nil, fmt.Errorf("iv given was not the correct length, was given %d bytes, requires 16 bytes", len(iv))
	}

	blocks := slices.Chunk(ciphertext, AESBlockSize)
	var plaintext []byte
	for block := range blocks {
		pt, err := InvCipher(block, key)
		if err != nil {
			return nil, fmt.Errorf("this block could not be decrypted: %#v", block)
		}
		pt = bytes.XORBytes(pt, iv)
		iv = block
		plaintext = append(plaintext, pt...)
	}
	return plaintext, nil
}
