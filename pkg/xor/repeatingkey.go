package xor

import "fmt"

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
