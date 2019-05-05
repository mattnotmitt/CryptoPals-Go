package set1

import (
	"CryptoPals/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChal6(t *testing.T) {
	hd := util.HammingDistance([]byte("this is a test"), []byte("wokka wokka!!!"))
	expectedHd := 37
	Chal6("data/6.txt")
	assert.Equal(t, expectedHd, hd)
}