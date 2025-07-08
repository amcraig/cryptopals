package aes

import (
	"math/rand"
)

func EBCCBCEncryptionOracleGenerator(plaintext []byte) ([]byte, bool, error) {
	var paddedPT []byte
	// Pad Before
	for range rand.Intn(6) + 5 {
		paddedPT = append(paddedPT, byte(rand.Intn(256)))
	}
	// Sandwich plaintext
	paddedPT = append(paddedPT, plaintext...)
	// Pad After
	for range rand.Intn(6) + 5 {
		paddedPT = append(paddedPT, byte(rand.Intn(256)))
	}

	key, err := GenerateAESKey(16)
	if err != nil {
		return nil, false, err
	}

	var ct []byte
	coinToss := rand.Intn(2) == 1
	if coinToss {
		// Encrypt EBC
		ct, err = EncryptAESECB(paddedPT, key)
		if err != nil {
			return nil, false, err
		}
	} else {
		// Encrypt CBC
		iv, err := GenerateAESKey(16)
		if err != nil {
			return nil, false, err
		}
		ct, err = EncryptAESCBC(paddedPT, key, iv)
		if err != nil {
			return nil, false, err
		}
	}
	return ct, coinToss, nil
}

func EBCCBCEncryptionOracle(ciphertext []byte) (bool, error) {
	// Returns true if an EBC block cipher was detected, else false indicates CBC
	// specifically for the scope of this challenge: https://cryptopals.com/sets/2/challenges/11
	val, err := DetectRedundantAESBlocks(ciphertext)
	if err != nil {
		return false, err
	}
	return val > 0, nil
}
