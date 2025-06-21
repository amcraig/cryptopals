package set1

import (
	"log"
	"testing"

	"github.com/amcraig/cryptopals-go/internal/common"
	"github.com/amcraig/cryptopals-go/pkg/xor"
)

func TestDetectSingleCharacterXOR(t *testing.T) {
	var inputFile string = "./testFiles/4.txt"
	charStrings := common.ReadFileIntoStringSlice(inputFile)

	detectedMessage := ""
	detectedChar := ""
	detectedIndex := 0
	curScore := 0.
	for idx, input := range charStrings {
		mes, char, score, err := xor.SingleByteXORDecode(input)
		if err != nil {
			log.Fatal(err)
		}
		if score > curScore {
			curScore = score
			detectedMessage = mes
			detectedChar = char
			detectedIndex = idx
		}
	}

	if detectedChar != "5" &&
		detectedIndex != 170 &&
		detectedMessage != "Now that the party is jumping\n" {
		t.Error("Message found is not what is expected")
	}

	// fmt.Printf("%f %d %s %#v", curScore, detectedIndex, detectedChar, detectedMessage)
}
