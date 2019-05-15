package set1

import (
	"CryptoPals/util"
)

func Chal7 (enc, key []byte) []byte {
	return util.AESECBDecrypt(enc, key)
}