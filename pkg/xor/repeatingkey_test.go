package xor

import "testing"

func TestChunkTextByKeyIndex(t *testing.T) {
	const keysize = 3

	inputStr := "secret_text"
	wantStr := "sr_xeettcte"

	input := []byte(inputStr)

	blocks, _ := chunkTextByKeyIndex(input, keysize)

	gotStr := ""
	for _, v := range blocks {
		gotStr += string(v)
	}

	if wantStr != gotStr {
		t.Errorf("the correct chunk result was not given: %v", gotStr)
	}
}
