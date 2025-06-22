package common_test

import (
	"testing"

	"github.com/amcraig/cryptopals-go/internal/common"
)

func TestHammingDistance(t *testing.T) {
	a := "this is a test"
	b := "wokka wokka!!!"

	dist, _ := common.HammingDistance([]byte(a), []byte(b))

	if dist != 37 {
		t.Error(dist)
	}
}
