package xor

import (
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

func SingleByteXORDecode(buf []byte, spaceThreshold int) ([]byte, byte, float64, error) {
	var score float64
	var asciiByte byte
	var message []byte

	for i := range 256 {
		tmpbuf, _ := SingleByteXOREncode(buf, byte(i))
		curScore, _ := common.ScoreEnglishPhrase(tmpbuf, spaceThreshold)

		if curScore == 0. {
			continue
		}

		if curScore > score {
			score = curScore
			asciiByte = byte(i)
			message = tmpbuf
		}
	}

	return message, asciiByte, score, nil
}
