package set1

import (
	"github.com/mattnotmitt/CryptoPals-go/util"
	"sort"
)

type BFXOR struct {
	score float64
	prob  float64
	key   byte
	res   []byte
}

func chal3(inp []byte) (string, float64, byte) {
	bfs := make([]BFXOR, 256)

	for b := 0; b < 256; b++ {
		res := util.XOR(inp, []byte{byte(b)})
		score, prob := util.ScoreString(res)
		bfs[b] = BFXOR{
			score,
			prob,
			byte(b),
			res,
		}
	}

	sort.Slice(bfs, func(i, j int) bool { return bfs[i].score < bfs[j].score })
	return string(bfs[0].res), bfs[0].score, bfs[0].key
}
