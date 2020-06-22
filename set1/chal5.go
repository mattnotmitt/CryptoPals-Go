package set1

import "CryptoPals/util"

func chal5(inp string, key string) string {
	return util.ToHex(util.XOR([]byte(inp), []byte(key)))
}
