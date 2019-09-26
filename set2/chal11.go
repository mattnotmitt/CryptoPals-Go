package set2

import (
	"CryptoPals/util"
	"fmt"
	"math/rand"
	"time"
)

func AESOracle (pt []byte) []byte {
	var encrypted []byte
	key := util.RandBytes(16)
	pt = append(util.RandBytes(rand.Intn(6) + 5), append(pt, util.RandBytes(rand.Intn(6) + 5)...)...)
	choice := rand.Intn(2)
	switch choice {
	case 0:
		fmt.Println("Using CBC mode.")
		encrypted = util.AESCBCEncrypt(pt, key, []byte("\x00"))
	case 1:
		fmt.Println("Using ECB mode.")
		encrypted = util.AESECBEncrypt(pt, key)
	}
	return encrypted
}

func Chal11 () {
	rand.Seed(time.Now().UnixNano())
	enc := AESOracle([]byte("Now that you have ECB and CBC working:\n\nWrite a function to generate a random AES key; that's just 16 random bytes.\n\nWrite a function that encrypts data under an unknown key --- that is, a function that generates a random key and encrypts under it.\n\nThe function should look like:"))
	_, score := util.DetectECB(enc, 16)
	if score == 0 {
		fmt.Println("CBC Mode")
	} else {
		fmt.Println("ECB Mode")
	}
}