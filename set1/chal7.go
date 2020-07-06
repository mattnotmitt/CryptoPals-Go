package set1

import (
	"github.com/mattnotmitt/CryptoPals-go/util"
)

func chal7(enc, key []byte) []byte {
	return util.AESECBDecrypt(enc, key)
}
