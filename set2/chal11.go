package set2

import (
	"github.com/mattnotmitt/CryptoPals-go/util"
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

func aesOracle(pt []byte) ([]byte, int) {
	var encrypted []byte
	key := util.RandBytes(16)
	pt = append(util.RandBytes(rand.Intn(6)+5), append(pt, util.RandBytes(rand.Intn(6)+5)...)...)
	choice := rand.Intn(2)
	switch choice {
	case 0:
		fmt.Println("Using CBC mode.")
		encrypted = util.AESCBCEncrypt(pt, key, []byte("\x00"))
	case 1:
		fmt.Println("Using ECB mode.")
		encrypted = util.AESECBEncrypt(pt, key)
	}
	return encrypted, choice
}

func chal11() (int, int) {
	rand.Seed(time.Now().UnixNano())
	enc, choice := aesOracle(bytes.Repeat([]byte("A"), 80))

	_, score := util.DetectECB(enc, 16)
	if score == 0 {
		return 0, choice
	}
	return 1, choice
}
