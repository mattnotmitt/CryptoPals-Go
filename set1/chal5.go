package set1

import "CryptoPals/util"

func Chal5(inp string, key string) string {
	return util.ToHex(util.XOR([]byte(inp), []byte(key)))
}