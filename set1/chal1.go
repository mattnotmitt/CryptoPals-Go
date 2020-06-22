package set1

import (
	"CryptoPals/util"
)

func chal1(inp string) string {
	rawInp := util.FromHex(inp)
	return util.ToBase64(rawInp)
}
