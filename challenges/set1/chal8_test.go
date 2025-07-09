package set1

import (
	"testing"

	"github.com/amcraig/cryptopals-go/internal/fileio"
	"github.com/amcraig/cryptopals-go/pkg/aes"
)

func Test8DetectAESInECBMode(t *testing.T) {
	ciphertexts := fileio.ReadFileLinesIntoByteSlices("./testFiles/8.txt")

	score := 0
	scoringIdx := -1
	for idx, ct := range ciphertexts {
		val, _ := aes.DetectRedundantAESBlocks(ct)
		if val > score {
			score = val
			scoringIdx = idx
		}
	}

	if scoringIdx != 132 {
		t.Errorf("wrong ciphertext was picked, got: %d, wanted %d", scoringIdx, 132)
	}
}
