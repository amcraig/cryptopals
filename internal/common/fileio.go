package common

import (
	"bufio"
	"log"
	"os"
	"strings"
)

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
