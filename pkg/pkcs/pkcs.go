package pkcs

func AddPKCS7Padding(plaintext []byte, blocksize int) ([]byte, error) {
	paddingByte := 0x04
	isMultiple := len(plaintext)%blocksize == 0
	var underFlow int
	if isMultiple {
		underFlow = blocksize - len(plaintext)
	} else {
		underFlow = blocksize - len(plaintext)%blocksize
	}

	t := make([]byte, underFlow)
	for i := range t {
		t[i] = byte(paddingByte)
	}
	return append(plaintext, t...), nil
}
