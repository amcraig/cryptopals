package pkcs

import (
	"reflect"
	"strings"
	"testing"
)

// Basic usage tested in Test9ImplementPKCS7Padding

// Test if padding block size is a multiple of the input len
func TestPKCS7PaddingEdgeCase(t *testing.T) {
	in := "YELLOW SUBMARINE" // len = 16
	want := []byte(in + strings.Repeat("\x10", len(in)))

	got, _ := AddPKCS7Padding([]byte(in), 2*len(in))

	if reflect.DeepEqual(want, got) == false {
		t.Errorf("padding result is not expected, got: %s, want: %s", got, want)
	}

	if got[len(got)-1] != 16 {
		t.Errorf("padding byte is not expected, got: %#v, want: %#v", got[len(got)-1], 16)
	}
}
