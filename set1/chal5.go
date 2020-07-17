package set1

import "github.com/mattnotmitt/CryptoPals-Go/util"

func chal5(inp string, key string) string {
	return util.ToHex(util.XOR([]byte(inp), []byte(key)))
}
