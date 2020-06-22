package set2

import (
	"CryptoPals/util"
	"io/ioutil"
)

func chal10(fp string) []byte {
	data, err := ioutil.ReadFile(fp)
	//fmt.Println(string(data[:len(data)-1]))
	if err != nil {
		panic(err)
	}
	enc := util.FromBase64(string(data))
	return util.AESCBCDecrypt(enc, []byte("YELLOW SUBMARINE"), []byte("\x00"))
}
