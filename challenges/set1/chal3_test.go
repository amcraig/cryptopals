package set1

import (
	hex "encoding/hex"
	"fmt"
	"testing"

	"github.com/amcraig/cryptopals-go/pkg/xor"
)

func Test3SingleByteXORCipher(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	buf, _ := hex.DecodeString(input)
	message, asciiChar, score, _ := xor.SingleByteXORDecode(buf, 3)

	if message != "Cooking MC's like a pound of bacon" {
		t.Error("Message is incorrect:", message)
	}

	if string(asciiChar) != "X" {
		t.Error("Byte for XOR Cypher is incorrect:", asciiChar)
	}

	fmt.Println(score, "-", asciiChar, "-", message)
}
