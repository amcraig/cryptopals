package pkcs

func AddPKCS7Padding(plaintext []byte, blocksize int) ([]byte, error) {
	paddingByte := 0x04
	isMultiple := len(plaintext)%blocksize == 0
	var underflow int
	if isMultiple {
		underflow = blocksize - len(plaintext)
	} else {
		underflow = blocksize - len(plaintext)%blocksize
	}

	t := make([]byte, underflow)
	for i := range t {
		t[i] = byte(paddingByte)
	}
	return append(plaintext, t...), nil
}
