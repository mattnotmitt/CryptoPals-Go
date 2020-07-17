package set3

import (
	"bytes"
	"github.com/mattnotmitt/CryptoPals-Go/util"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestChal18(t *testing.T) {
	data, err := ioutil.ReadFile("data/18.txt")
	util.Check(err)
	enc := util.FromBase64(string(data))
	pt, err := util.AESCTRDecrypt(enc, []byte("YELLOW SUBMARINE"), bytes.Repeat([]byte("\x00"), 8))
	assert.Nil(t, err)
	_, p := util.ScoreString(pt)
	assert.Greater(t, p,0.75)
}
