package common

import (
	"fmt"
	"math/bits"

	"github.com/amcraig/cryptopals-go/internal/constants"
)

func ScoreEnglishPhrase(cyphertext []byte, spaceThreshold int) (float64, error) {

	// Calculates the score based on letter frequencies
	score := 0.
	spaces := 0
	for _, b := range cyphertext {
		// If the current char is a carriage return or line feed just continue without accounting for score
		if b == byte(10) || b == byte(12) {
			continue
		}

		// If the phrase contains non-printable ASCII characters return score of 0
		if b < byte(32) || b > byte(127) {
			return 0., nil
		}

		if b == byte(32) {
			spaces++
			continue
		}

		// Convert to lowercase to uppercase to score with map
		if b >= 'a' && b <= 'z' {
			b = b - ('a' - 'A')
		}

		freq, ok := constants.EnglishLetterFreqMap[byte(b)]
		if ok {
			score += freq * float64(len(cyphertext))
		}
	}

	if spaces < spaceThreshold {
		return 0., nil
	}

	return score, nil
}

// The (bytewise) Hamming distance is just the number of differing bits.
func HammingDistance(a, b []byte) (int, error) {
	if len(a) != len(b) {
		return -1, fmt.Errorf("inputs lengths do not match")
	}
	score := 0
	for i := range a {
		ai := a[i]
		bi := b[i]
		if ai != bi {
			diff := ai ^ bi
			// https://en.wikipedia.org/wiki/Hamming_weight
			// Also known as the population count or sideways add of a bitstring
			bitcount := bits.OnesCount8(diff)
			score += bitcount
		}
	}
	return score, nil
}
