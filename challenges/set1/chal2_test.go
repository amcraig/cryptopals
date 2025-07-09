package set1

import (
	hex "encoding/hex"
	"testing"

	"github.com/amcraig/cryptopals-go/internal/bytes"
)

func Test2FixedXOR(t *testing.T) {
	input1 := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"
	want := "746865206b696420646f6e277420706c6179"

	// Decodes hex strings into their actual bytes
	decBuf1, _ := hex.DecodeString(input1)
	decBuf2, _ := hex.DecodeString(input2)

	// XORs each byte in each slice with eachother
	xorBuf, _ := bytes.XORByteSlice(decBuf1, decBuf2)

	// Encode bytes back to hex string
	resp := hex.EncodeToString(xorBuf)

	if resp != want {
		t.Error(resp)
	}
}
