package aes

import (
	"reflect"
	"slices"
	"testing"

	"github.com/amcraig/cryptopals-go/internal/constants"
)

func TestReadBlockIntoState(t *testing.T) {
	input := []byte("0123456789abcdef")
	want := "048c159d26ae37bf"

	state, _ := readBlockIntoState(input)
	got := slices.Concat(state...)

	if want != string(got) {
		t.Errorf("AES state is not read in as expected: %v", string(got))
	}
}

func TestStateToWords(t *testing.T) {
	state, _ := readBlockIntoState([]byte("0123456789abcdef"))
	want := word([]byte("0123"))
	words := state.Words()
	if reflect.DeepEqual(words[0], want) == false {
		t.Errorf("first word did not match expected. Got: %s Want: %s", words[0], want)
	}
}

func TestSetStateFromWords(t *testing.T) {
	words := []word{
		{0x00, 0x01, 0x02, 0x03},
		{0x04, 0x05, 0x06, 0x07},
		{0x08, 0x09, 0x0a, 0x0b},
		{0x0c, 0x0d, 0x0e, 0x0f},
	}

	var s state
	s.SetStateFromWords(words)
}

func TestKeyExpansion(t *testing.T) {
	key := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
	lastFour := []byte{0x13, 0x11, 0x1d, 0x7f, 0xe3, 0x94, 0x4a, 0x17, 0xf3, 0x07, 0xa7, 0x8b, 0x4d, 0x2b, 0x30, 0xc5}
	keySchedule, err := keyExpansion(key)
	if err != nil {
		t.Errorf("key expansion failed: %v", err)
	}
	if len(keySchedule) != 44 {
		t.Errorf("key schedule is not right length for a 128 bit key")
	}
	if reflect.DeepEqual([]byte(slices.Concat(keySchedule[:4]...)), key) == false {
		t.Errorf("first four words do not match initial key")
	}
	if reflect.DeepEqual([]byte(slices.Concat(keySchedule[len(keySchedule)-4:]...)), lastFour) == false {
		t.Errorf("last four words do not match expected")
	}
}

func TestWordXORWord(t *testing.T) {
	a1 := word{1, 0, 1, 0}
	b1 := word{0, 1, 1, 0}
	want1 := word{1, 1, 0, 0}
	got1 := XORWord(a1, b1)
	if reflect.DeepEqual(got1, want1) == false {
		t.Errorf("XOR of %v and %v is not equal to %v: %v", a1, b1, want1, got1)
	}
	a2 := word{1, 0, 1, 0}
	b2 := word{0, 1, 1, 0}
	want2 := word{1, 1, 0, 0}
	got2 := XORWord(a2, b2)
	if reflect.DeepEqual(got2, want2) == false {
		t.Errorf("XOR of %v and %v is not equal to %v: %v", a2, b2, want2, got2)
	}
}

func TestWordSubWord(t *testing.T) {
	w := word{0x00, 0x11, 0x22, 0x33}
	want := word{0x63, 0x82, 0x93, 0xC3}
	got := SubWord(w)
	if reflect.DeepEqual(got, want) == false {
		t.Errorf("SubWord of %v is not equal to %v", w, want)
	}
}

func TestWordRot(t *testing.T) {
	w := word{0x00, 0x11, 0x22, 0x33}
	want := word{0x11, 0x22, 0x33, 0x00}
	got := RotWord(w, 1)
	if reflect.DeepEqual(got, want) == false {
		t.Errorf("rotation result is not expected, got: %v wanted: %v", got, want)
	}
	RotWord(w, 1)
	RotWord(w, 1) // checking methods do not mutate owner
	if reflect.DeepEqual(got, want) == false {
		t.Errorf("rotation result is not expected, got: %v wanted: %v", got, want)
	}
}

func TestMethodChaining(t *testing.T) {
	w := word{0x00, 0x11, 0x22, 0x33}
	want := SubWord(word{0x11, 0x22, 0x33, 0x00})
	got := SubWord(RotWord(w, 1))
	if reflect.DeepEqual(got, want) == false {
		t.Errorf("method chaining did not work as expected, got: %v, wanted: %v", got, want)
	}

	want2 := XORWord(SubWord(word{0x11, 0x22, 0x33, 0x00}), word{0x01, 0x00, 0x00, 0x00})
	got = XORWord(got, word{constants.RoundConstant[1], 0x00, 0x00, 0x00})
	if reflect.DeepEqual(got, want2) == false {
		t.Errorf("method chaining did not work as expected, got: %v, wanted: %v", got, want)
	}
}

func TestMixColumns(t *testing.T) {
	s := state{
		{0xd4, 0x01, 0xd4, 0xf2},
		{0xbf, 0x01, 0xd4, 0x0a},
		{0x5d, 0x01, 0xd4, 0x22},
		{0x30, 0x01, 0xd5, 0x5c},
	}

	want := state{
		{0x04, 0x01, 0xd5, 0x9f},
		{0x66, 0x01, 0xd5, 0xdc},
		{0x81, 0x01, 0xd7, 0x58},
		{0xe5, 0x01, 0xd6, 0x9d},
	}
	got := s.MixColumns()

	if reflect.DeepEqual(got, want) == false {
		t.Errorf("generated state mix was not expected, got: %v wanted: %v", got, want)
	}
}

func TestAddRoundKey(t *testing.T) {
	s, _ := readBlockIntoState(
		[]byte{
			0x32, 0x43, 0xf6, 0xa8, 0x88, 0x5a, 0x30, 0x8d,
			0x31, 0x31, 0x98, 0xa2, 0xe0, 0x37, 0x07, 0x34,
		})

	key := []byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c}
	ks, _ := keyExpansion(key)

	want := state{
		{0x19, 0xa0, 0x9a, 0xe9},
		{0x3d, 0xf4, 0xc6, 0xf8},
		{0xe3, 0xe2, 0x8d, 0x48},
		{0xbe, 0x2b, 0x2a, 0x08},
	}

	got := s.AddRoundKey(ks, 0)

	if reflect.DeepEqual(got, want) == false {
		t.Errorf("round key added to state was not expected, \ngot: %#v \n\nwanted: %#v", got, want)
	}

}

func TestCipher(t *testing.T) {
	in := []byte{0x32, 0x43, 0xf6, 0xa8, 0x88, 0x5a, 0x30, 0x8d, 0x31, 0x31, 0x98, 0xa2, 0xe0, 0x37, 0x07, 0x34}
	key := []byte{0x2b, 0x7e, 0x15, 0x16, 0x28, 0xae, 0xd2, 0xa6, 0xab, 0xf7, 0x15, 0x88, 0x09, 0xcf, 0x4f, 0x3c}
	want := []byte{0x39, 0x25, 0x84, 0x1d, 0x02, 0xdc, 0x09, 0xfb, 0xdc, 0x11, 0x85, 0x97, 0x19, 0x6a, 0x0b, 0x32}

	ks, _ := keyExpansion(key)
	got := Cipher(in, 10, ks)

	if reflect.DeepEqual(got, want) == false {
		t.Errorf("cipher output is unexpected. Got: %#v Wanted:%#v", got, want)
	}
}
