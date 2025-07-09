package set1

import (
	hex "encoding/hex"
	"log"
	"testing"

	"github.com/amcraig/cryptopals-go/internal/fileio"
	"github.com/amcraig/cryptopals-go/pkg/xor"
)

func Test4DetectSingleCharacterXOR(t *testing.T) {
	var inputFile string = "./testFiles/4.txt"
	charStrings := fileio.ReadFileLinesIntoStringSlices(inputFile)

	var detectedMessage []byte
	var detectedByte byte
	detectedIndex := 0
	curScore := 0.
	for idx, input := range charStrings {
		buf, _ := hex.DecodeString(input)
		mesBuf, asciiByte, score, err := xor.SingleByteXORDecode(buf, 3)
		if err != nil {
			log.Fatal(err)
		}
		if score > curScore {
			curScore = score
			detectedMessage = mesBuf
			detectedByte = asciiByte
			detectedIndex = idx
		}
	}

	if string(detectedByte) != "5" &&
		detectedIndex != 170 &&
		string(detectedMessage) != "Now that the party is jumping\n" {
		t.Error("Message found is not what is expected")
	}

	// fmt.Printf("%f %d %s %#v", curScore, detectedIndex, detectedChar, detectedMessage)
}
