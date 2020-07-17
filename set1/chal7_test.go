package set1

import (
	"github.com/mattnotmitt/CryptoPals-Go/util"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestChal7(t *testing.T) {
	data, err := ioutil.ReadFile("data/7.txt")
	util.Check(err)
	enc, err := base64.StdEncoding.DecodeString(string(data))
	key := []byte("YELLOW SUBMARINE")
	plaintext := chal7(enc, key)
	fmt.Println(string(plaintext))
}
