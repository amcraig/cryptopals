package aes

import (
	"fmt"
	"slices"

	"github.com/amcraig/cryptopals-go/internal/bytes"
	"github.com/amcraig/cryptopals-go/pkg/pkcs"
)

func EncryptAESCBC(plaintext []byte, key []byte, iv []byte) ([]byte, error) {
	if len(iv) != 16 {
		return nil, fmt.Errorf("iv given was not the correct length, was given %d bytes, requires 16 bytes", len(iv))
	}
	blocks := slices.Chunk(plaintext, AESBlockSize)
	var ciphertext []byte
	for block := range blocks {
		if len(block) != AESBlockSize {
			block, _ = pkcs.AddPKCS7Padding(block, AESBlockSize)
		}
		block, err := bytes.XORByteSlice(block, iv)
		if err != nil {
			return nil, err
		}
		ct, err := Cipher(block, key)
		if err != nil {
			return nil, fmt.Errorf("this block could not be encrypted: %#v", block)
		}
		iv = ct
		ciphertext = append(ciphertext, ct...)
	}
	return ciphertext, nil
}

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
		pt, err = bytes.XORByteSlice(pt, iv)
		if err != nil {
			return nil, err
		}
		iv = block
		plaintext = append(plaintext, pt...)
	}
	return plaintext, nil
}
