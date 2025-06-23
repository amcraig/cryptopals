package set1

import (
	"encoding/base64"
	"testing"

	"github.com/amcraig/cryptopals-go/internal/common"
	"github.com/amcraig/cryptopals-go/pkg/xor"
)

func Test6BreakRepeatingKeyXOR(t *testing.T) {
	// Input file had been Base64'd after being encrypt with repeating-key XOR
	inputStrBase64 := common.ReadFileIntoString("./testFiles/6.txt")
	encryptedBuf, _ := base64.StdEncoding.DecodeString(inputStrBase64)
	want := "Terminator X: Bring the noise"

	got, _ := xor.BreakRepeatingKeyXOR(encryptedBuf, 40)

	if want != string(got) {
		t.Errorf("key generated does not match expected: %v", string(got))
	}

	// Optional: Decode message (it's long)
	// mes, _ := xor.RepeatingKeyXORDecode(encryptedBuf, []byte(got))
	// fmt.Println(string(mes))
}
