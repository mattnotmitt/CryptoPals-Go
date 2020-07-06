package set1

import "github.com/mattnotmitt/CryptoPals-go/util"

func chal5(inp string, key string) string {
	return util.ToHex(util.XOR([]byte(inp), []byte(key)))
}
