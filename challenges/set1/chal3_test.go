package set1

import (
	"testing"

	"github.com/amcraig/cryptopals-go/pkg/xor"
)

func TestSingleByteXORCipher(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	message, asciiChar, _, _ := xor.SingleByteXORDecode(input)

	if message != "Cooking MC's like a pound of bacon" {
		t.Error("Message is incorrect:", message)
	}

	if asciiChar != "X" {
		t.Error("Byte for XOR Cypher is incorrect:", asciiChar)
	}

	// fmt.Println(score, "-", asciiChar, "-", message)
}
