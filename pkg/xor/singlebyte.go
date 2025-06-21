package xor

import (
	"encoding/hex"

	"github.com/amcraig/cryptopals-go/internal/common"
)

// XOR encodes a Byte Slice by single byte
func SingleByteXOREncode(a []byte, b byte) ([]byte, error) {
	buf := make([]byte, len(a))
	for i := range a {
		buf[i] = a[i] ^ b
	}
	return buf, nil
}

func SingleByteXORDecode(cyphertext string) (string, string, float64, error) {
	buf, _ := hex.DecodeString(cyphertext)

	score := 0.
	asciiChar := ""
	message := ""

	for i := range 256 {
		tmpbuf, _ := SingleByteXOREncode(buf, byte(i))
		curScore, _ := common.ScoreEnglishPhrase(tmpbuf)

		if curScore == 0. {
			continue
		}

		if curScore > score {
			score = curScore
			asciiChar = string(byte(i))
			message = string(tmpbuf)
		}
	}

	return message, asciiChar, score, nil
}
