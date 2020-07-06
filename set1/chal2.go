package set1

import "github.com/mattnotmitt/CryptoPals-go/util"

func chal2(inp, key string) string {
	rawInp := util.FromHex(inp)
	rawKey := util.FromHex(key)
	return util.ToHex(util.XOR(rawInp, rawKey))
}
