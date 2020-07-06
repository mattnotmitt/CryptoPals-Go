package set3

import (
	"github.com/mattnotmitt/CryptoPals-go/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaddingValid (t *testing.T) {
	enc, _, _ := encRandom()
	valid := paddingValid(enc)
	assert.Equal(t, true, valid)
}

func TestCBCPaddingOracle (t *testing.T) {
	enc, iv, pt := encRandom()
	dec, err := breakOracle(enc, iv)
	assert.Nil(t, err)
	assert.Equal(t, pt, util.UnPKCS7(dec))
	// fmt.Println(string(breakOracle(enc, iv)))
}