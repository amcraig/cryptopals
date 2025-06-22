package set1

import (
	hex "encoding/hex"
	"log"
	"testing"

	"github.com/amcraig/cryptopals-go/internal/common"
	"github.com/amcraig/cryptopals-go/pkg/xor"
)

func Test4DetectSingleCharacterXOR(t *testing.T) {
	var inputFile string = "./testFiles/4.txt"
	charStrings := common.ReadFileIntoStringSlice(inputFile)

	detectedMessage := ""
	var detectedByte byte
	detectedIndex := 0
	curScore := 0.
	for idx, input := range charStrings {
		buf, _ := hex.DecodeString(input)
		mes, asciiByte, score, err := xor.SingleByteXORDecode(buf, 3)
		if err != nil {
			log.Fatal(err)
		}
		if score > curScore {
			curScore = score
			detectedMessage = mes
			detectedByte = asciiByte
			detectedIndex = idx
		}
	}

	if string(detectedByte) != "5" &&
		detectedIndex != 170 &&
		detectedMessage != "Now that the party is jumping\n" {
		t.Error("Message found is not what is expected")
	}

	// fmt.Printf("%f %d %s %#v", curScore, detectedIndex, detectedChar, detectedMessage)
}
