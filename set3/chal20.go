package set3

import (
	"bytes"
	"github.com/mattnotmitt/CryptoPals-Go/set1"
	"github.com/mattnotmitt/CryptoPals-Go/util"
	"io/ioutil"
	"sort"
	"strings"
	"sync"
)

var key19 = util.RandBytes(16)

func generateCipherText() ([][]byte, []byte) {
	data, err := ioutil.ReadFile("data/20.txt")
	util.Check(err)
	strArray := strings.Split(string(data), "\n")
	bArrArr := make([][]byte, len(strArray))

	var wg sync.WaitGroup
	wg.Add(len(strArray))
	for i, v := range strArray {
		go func(val string, i int) {
			defer wg.Done()
			enc, err := util.AESCTREncrypt(util.FromBase64(val), key19, bytes.Repeat([]byte("\x00"), 8))
			util.Check(err)
			bArrArr[i] = enc
		}(v, i)
	}
	wg.Wait()
	sort.Slice(bArrArr, func(i, j int) bool { return len(bArrArr[i]) < len(bArrArr[j]) })
	return bArrArr, key19
}

func breakFixNonceCTR(encs [][]byte) ([]string, []byte) {
	var truncEncs []byte
	minLen := len(encs[0])
	for _, v := range encs {
		truncEncs = append(truncEncs, v[:minLen]...)
	}
	dec, keyStream := set1.BreakRepeatingXOR(truncEncs, minLen)
	splitDec := make([]string, len(encs))
	start := 0
	end := minLen
	for i := 0; i < len(encs); i++ {
		splitDec[i] = string(dec[start:end])
		start = end
		end = start + minLen
	}
	return splitDec, keyStream
}