package set1

import (
	"encoding/hex"
	"testing"

	"github.com/amcraig/cryptopals-go/pkg/xor"
)

func Test5RepeatingKeyXOR(t *testing.T) {
	// Unsure if there is a newline char here or should just be a space
	input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	encryptionKey := "ICE"
	want := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	bufIn := []byte(input)
	bufKey := []byte(encryptionKey)

	encodedBuf, _ := xor.RepeatingKeyXOREncode(bufIn, bufKey)
	resp := hex.EncodeToString(encodedBuf)

	if resp != want {
		t.Error(resp)
	}

}
