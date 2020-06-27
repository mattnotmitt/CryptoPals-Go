package set2

import (
	"CryptoPals/util"
	"bytes"
)

func chal9(inp string, size int) string {
	return string(bytes.Join(util.PKCS7Pad([]byte(inp), size), []byte("")))
}
