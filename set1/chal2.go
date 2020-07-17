package set1

import "github.com/mattnotmitt/CryptoPals-Go/util"

func chal2(inp, key string) string {
	rawInp := util.FromHex(inp)
	rawKey := util.FromHex(key)
	return util.ToHex(util.XOR(rawInp, rawKey))
}
