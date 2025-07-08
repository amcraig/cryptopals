package set2

import (
	"testing"

	"github.com/amcraig/cryptopals-go/internal/fileio"
	"github.com/amcraig/cryptopals-go/pkg/aes"
)

func Test11AnECBCBCDetectionOracle(t *testing.T) {
	ciphertext := fileio.ReadBase64FileIntoByteSlice("./testFiles/10.txt")

	ct, isECB, err := aes.EBCCBCEncryptionOracleGenerator(ciphertext)
	if err != nil {
		t.Errorf("ECBCBCEncryptionOracleGenerator failed: %s", err)
	}

	prediction, err := aes.EBCCBCEncryptionOracle(ct)
	if err != nil {
		t.Errorf("ECBCBCEncryptionOracle failed: %s", err)
	}

	if prediction != isECB {
		t.Errorf("ECB CBC Detection Oracle Failed, got: %v, want: %v", prediction, isECB)
	}

}
