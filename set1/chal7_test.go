package set1

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestChal7(t *testing.T) {

	data, err := ioutil.ReadFile("data/7.txt")
	//fmt.Println(string(data[:len(data)-1]))
	if err != nil {
		panic(err)
	}
	enc, err := base64.StdEncoding.DecodeString(string(data))
	key := []byte("YELLOW SUBMARINE")
	plaintext := Chal7(enc, key)
	fmt.Println(string(plaintext))
}