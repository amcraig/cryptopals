package bytes

// Tested in aes_test.go
func XORBytes(a, b []byte) []byte {
	if len(a) != len(b) {
		panic("the byte slices are not equal length")
	}
	c := make([]byte, len(a))
	for i := range a {
		c[i] = a[i] ^ b[i]
	}
	return c
}
