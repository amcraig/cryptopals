package matrix

import (
	"reflect"
	"testing"
)

func TestTranspose(t *testing.T) {
	input := [][]byte{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	want := [][]byte{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	}

	got := Transpose(input)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestRotateVector(t *testing.T) {
	input := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
	want := []byte{3, 4, 5, 6, 7, 8, 9, 1, 2}

	// Test shift left
	n := 2
	got := RotateVector(input, n)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

	// Test shift right
	n = -7
	got = RotateVector(input, n)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

	// Test greater than len(vector) shift to left
	n = 20
	got = RotateVector(input, n)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

	// Test greater than len(vector) shift to right
	n = -16
	got = RotateVector(input, n)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}

	// Test no shift
	n = 0
	got = RotateVector(input, n)
	if !reflect.DeepEqual(got, input) {
		t.Errorf("got %v, want %v", got, want)
	}
}
