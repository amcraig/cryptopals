package aes

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateAESKey(keysize int) ([]byte, error) {
	switch keysize {
	case 16, 24, 32:
	default:
		return nil, fmt.Errorf("the incoming key is does not match: 16, 24, 32 bits")
	}

	key := make([]byte, keysize)
	for i := range keysize {
		val, err := rand.Int(rand.Reader, big.NewInt(256))
		if err != nil {
			return nil, err
		}
		key[i] = byte(val.Int64())
	}
	return key, nil
}
