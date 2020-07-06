package set2

import (
	"github.com/mattnotmitt/CryptoPals-go/util"
	"regexp"
	"sync"
)

var keySetup16 sync.Once
var key16 []byte
func vanillaIceify(arb string) []byte {
	keySetup16.Do(func() { key16 = util.RandBytes(16) }) // Generate key on first run of program and persist
	re := regexp.MustCompile(`[;=]`)
	arbClean := []byte(re.ReplaceAllString(arb, ""))
	ice := append([]byte("comment1=cooking%20MCs;userdata="), append(arbClean, ";comment2=%20like%20a%20pound%20of%20bacon"...)...)
	return util.AESCBCEncrypt(ice, key16, []byte("\x00"))
}

func checkAdmin(enc []byte) bool {
	dec := util.AESCBCDecrypt(enc, key16, []byte("\x00"), true)
	re := regexp.MustCompile(`;admin=true;`)
	return re.Match(dec)
}

