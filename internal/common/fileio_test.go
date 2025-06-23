package common_test

import (
	"testing"

	"github.com/amcraig/cryptopals-go/internal/common"
)

func TestReadFileIntoString(t *testing.T) {
	filepath := "../../challenges/set1/testFiles/6_file_input_test.txt"
	want := "HUIfTQsPAh9PE048GmllH0kcDk4TAQsHThsBFkU2AB4BSWQgVB0dQzNTTmVSBgBHVBwNRU0HBAxTEjwMHghJGgkRTxRMIRpHKwAFHUdZEQQJAGQmB1MANxYG"

	got := common.ReadFileIntoString(filepath)

	if want != got {
		t.Errorf("File content processed does not match expected: %v", got)
	}
}
