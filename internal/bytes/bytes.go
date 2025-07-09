package bytes

import "fmt"

// Tested in aes_test.go
func XORByteSlice(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("length of byte slices are not equivalent: %d != %d", len(a), len(b))
	}
	c := make([]byte, len(a))
	for i := range a {
		c[i] = a[i] ^ b[i]
	}
	return c, nil
}
