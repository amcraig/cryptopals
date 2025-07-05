package set2

import (
	"reflect"
	"testing"

	"github.com/amcraig/cryptopals-go/pkg/pkcs"
)

func Test9ImplementPKCS7Padding(t *testing.T) {
	in := "YELLOW SUBMARINE"
	want := []byte(in + "\x04\x04\x04\x04")

	got, _ := pkcs.AddPKCS7Padding([]byte(in), 20)

	if reflect.DeepEqual(want, got) == false {
		t.Errorf("padding result is not expected, got: %s, want: %s", got, want)
	}
}
