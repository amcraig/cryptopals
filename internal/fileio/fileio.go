package fileio

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func ReadFileIntoByteSlice(filepath string) (sliceStr [][]byte) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fscanner := bufio.NewScanner(file)

	for fscanner.Scan() {
		textinput := strings.TrimSpace(fscanner.Text())
		sliceStr = append(sliceStr, []byte(textinput))
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
