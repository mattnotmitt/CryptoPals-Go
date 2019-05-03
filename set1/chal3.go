package set1

import (
	"CryptoPals/util"
)

func Chal3 (inp string) (string, float64, string) {
	rawInp := util.FromHex(inp)
	var best []byte
	bestScore := 0.0
	bestKey := 0

	for b := 0; b < 256; b++ {
		res := util.XOR(rawInp, []byte{byte(b)})
		score := util.ScoreString(string(res))
		if score > bestScore {
			bestScore = score
			best = res
			bestKey = b
		}
		//fmt.Println(score, string(res))
	}

	return string(best), bestScore, string(bestKey)
}