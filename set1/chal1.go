package set1

import (
	"github.com/mattnotmitt/CryptoPals-Go/util"
)

func chal1(inp string) string {
	rawInp := util.FromHex(inp)
	return util.ToBase64(rawInp)
}
