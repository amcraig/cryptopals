package fileio

import (
	"bufio"
	"encoding/base64"
	"log"
	"os"
	"strings"
)

func ReadBase64FileIntoByteSlice(filepath string) (sliceByte []byte) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fscanner := bufio.NewScanner(file)

	for fscanner.Scan() {
		textinput := strings.TrimSpace(fscanner.Text())
		b64DecodedBytes, _ := base64.StdEncoding.DecodeString(textinput)
		sliceByte = append(sliceByte, b64DecodedBytes...)
	}

	return
}

func ReadFileIntoByteSlice(filepath string) (sliceByte [][]byte) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fscanner := bufio.NewScanner(file)

	for fscanner.Scan() {
		textinput := strings.TrimSpace(fscanner.Text())
		sliceByte = append(sliceByte, []byte(textinput))
	}

	return
}

func ReadFileIntoStringSlice(filepath string) (sliceStr []string) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fscanner := bufio.NewScanner(file)

	for fscanner.Scan() {
		textinput := strings.TrimSpace(fscanner.Text())
		sliceStr = append(sliceStr, textinput)
	}

	return
}

func ReadFileIntoString(filepath string) string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fscanner := bufio.NewScanner(file)

	// It's ready to use from the get-go.
	// You don't need to initialize it.
	var strBuilder strings.Builder

	for fscanner.Scan() {
		textinput := strings.TrimSpace(fscanner.Text())
		strBuilder.WriteString(textinput)
	}

	return strBuilder.String()
}
