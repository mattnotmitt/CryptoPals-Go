package set3

import (
	"CryptoPals/set2"
	"CryptoPals/util"
	"bytes"
	"io/ioutil"
	"math/rand"
	"strings"
)

var key17 = util.RandBytes(16)

func encRandom () ([]byte, []byte) {
	data, err := ioutil.ReadFile("data/12.txt")
	util.Check(err)
	strArray := strings.Split(string(data), "\n")
	randStr := strArray[rand.Intn(len(strArray))]
	enc := util.AESCBCEncrypt([]byte(randStr), key17, []byte("\x00"))
	return enc, bytes.Repeat([]byte("\x00"), 16)
}

func paddingValid (enc []byte) bool {
	pt := util.AESCBCDecrypt(enc, key17, []byte("\x00"))
	_, err := set2.VerifyPadding(string(pt), 16)
	return err != nil
}