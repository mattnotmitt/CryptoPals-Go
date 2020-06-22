package set1

import "CryptoPals/util"

func chal2(inp, key string) string {
	rawInp := util.FromHex(inp)
	rawKey := util.FromHex(key)
	return util.ToHex(util.XOR(rawInp, rawKey))
}
