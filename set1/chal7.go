package set1

import (
	"github.com/mattnotmitt/CryptoPals-Go/util"
)

func chal7(enc, key []byte) []byte {
	return util.AESECBDecrypt(enc, key)
}
