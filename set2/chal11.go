package set2

import (
	"CryptoPals/util"
	"fmt"
	"math/rand"
	"time"
)

func Chal11 () {
	rand.Seed(time.Now().UnixNano())
	enc := util.AESOracle([]byte("Now that you have ECB and CBC working:\n\nWrite a function to generate a random AES key; that's just 16 random bytes.\n\nWrite a function that encrypts data under an unknown key --- that is, a function that generates a random key and encrypts under it.\n\nThe function should look like:"))
	_, score := util.DetectECB(enc)
	if score == 0 {
		fmt.Println("CBC Mode")
	} else {
		fmt.Println("ECB Mode")
	}
}