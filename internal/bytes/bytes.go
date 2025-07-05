package bytes

// Tested in aes_test.go
func XORBytes(a, b []byte) []byte {
	for i := range a {
		a[i] = a[i] ^ b[i]
	}
	return a
}
