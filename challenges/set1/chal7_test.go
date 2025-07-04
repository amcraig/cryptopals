package set1

import (
	"encoding/base64"
	"reflect"
	"testing"

	"github.com/amcraig/cryptopals-go/internal/fileio"
	"github.com/amcraig/cryptopals-go/pkg/aes"
)

func Test7AESInECBMode(t *testing.T) {
	/*
		Use AES-128-ECB as the cipher
		In ECB(Electronic codebook) mode, each block of plaintext is separately enciphered and each block of the ciphertext is separately deciphered.
		A block cipher by itself is only suitable for the secure cryptographic transformation (encryption or decryption) of one
		fixed-length group of bits called a block.
		The simplest of the encryption modes is the electronic codebook (ECB) mode (named after conventional physical codebooks).
		The message is divided into blocks, and each block is encrypted separately. ECB is not recommended for use in cryptographic protocols:
		the disadvantage of this method is a lack of diffusion, wherein it fails to hide data patterns when it encrypts identical plaintext
		blocks into identical ciphertext blocks.
		The Advanced Encryption Standard (AES), also known by its original name Rijndael is a specification for the encryption
		of electronic data established by the U.S. National Institute of Standards and Technology (NIST) in 2001.
		---
		AES operates on a 4 × 4 column-major order array of 16 bytes b0, b1, ..., b15 termed the state
		The key size used for an AES cipher specifies the number of transformation rounds that convert the input, called the plaintext,
		into the final output, called the ciphertext.
		The number of rounds are as follows:
		* 10 rounds for 128-bit keys.
		* 12 rounds for 192-bit keys.
		* 14 rounds for 256-bit keys.
		Each round consists of several processing steps, including one that depends on the encryption key itself. A set of
		reverse rounds are applied to transform ciphertext back into the original plaintext using the same encryption key.
		---
		Here we have a 128 bit (16 byte) key with a fixed block size also of 128 bits.

	*/
	const key = "YELLOW SUBMARINE"
	inputStrBase64 := fileio.ReadFileIntoString("./testFiles/7.txt")
	ciphertext, _ := base64.StdEncoding.DecodeString(inputStrBase64)
	plaintext, err := aes.DecryptAESECB(ciphertext, []byte(key))
	if err != nil {
		t.Errorf("could not decrypt ciphertext.")
	}

	want := []byte("I'm back and I'm")
	if reflect.DeepEqual(plaintext[:16], want) == false {
		t.Errorf("plaintext is not as expected. Want: %#v Got: %#v", want, plaintext[:16])
	}
}
