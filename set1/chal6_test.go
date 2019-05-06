package set1

import (
	"CryptoPals/util"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestChal6(t *testing.T) {
	hd := util.HammingDistance([]byte("this is a test"), []byte("wokka wokka!!!"))
	expectedHd := 37
	assert.Equal(t, expectedHd, hd)

	data, err := ioutil.ReadFile("data/6.txt")
	//fmt.Println(string(data[:len(data)-1]))
	if err != nil {
		panic(err)
	}
	enc := util.FromBase64(string(data))
	Chal6(enc, 1)
}