package set2

import "CryptoPals/util"

func chal9(inp string, size int) string {
	return string(util.PKCS7Pad([]byte(inp), size))
}
