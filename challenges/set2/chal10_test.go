package set2

import (
	"reflect"
	"strings"
	"testing"

	"github.com/amcraig/cryptopals-go/internal/fileio"
	"github.com/amcraig/cryptopals-go/pkg/aes"
)

func Test10ImplementCBCMode(t *testing.T) {
	ciphertext := fileio.ReadBase64FileIntoByteSlice("./testFiles/10.txt")

	key := []byte("YELLOW SUBMARINE")
	iv := []byte(strings.Repeat("\x00", aes.AESBlockSize))

	pt, _ := aes.DecryptAESCBC(ciphertext, key, iv)

	want := []byte("I'm back and I'm")
	got := pt[:16]

	if reflect.DeepEqual(want, got) == false {
		t.Errorf("AES CBC decrypt did not give expected results, want: %s, got %s", want, got)
	}
}
