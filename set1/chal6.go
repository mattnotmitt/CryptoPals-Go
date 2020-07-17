package set1

import (
	"github.com/mattnotmitt/CryptoPals-Go/util"
	"sort"
)

// oh godddddddddd

type Keysize struct {
	ks  int
	nhd float64
}

func BreakRepeatingXOR(enc []byte, keyLen int) ([]byte, []byte) {
		blocks := util.ChunkByteArray(enc, keyLen, false)
		transposedBlocks := make([][]byte, keyLen)

		for i := 0; i < len(blocks); i++ {
			for j := 0; j < len(blocks[i]); j++ {
				transposedBlocks[j] = append(transposedBlocks[j], blocks[i][j])
			}
		}

		keys := make([]byte, keyLen)
		for i, tbl := range transposedBlocks {
			_, _, k := chal3(tbl)
			keys[i] = k
		}
		dec := util.XOR(enc, keys)
		return dec, keys
}

func determineKeysize(enc []byte, noKeys int) []int {
	kss := make([]Keysize, 39)

	for ks := 2; ks < 41; ks++ {
		avgDist := 0.0
		iter := 0.0

		for i := 0; i+2*ks < len(enc); i += ks {
			lb, ub := enc[i:i+ks], enc[i+ks:i+2*ks]
			avgDist += float64(util.HammingDistance(lb, ub)) / float64(ks)
			iter++
		}

		nhd := avgDist / iter
		kss[ks-2] = Keysize{ks, nhd}
	}
	sort.Slice(kss, func(i, j int) bool { return kss[i].nhd < kss[j].nhd })
	kss = kss[:noKeys]
	keyLens := make([]int, noKeys)
	for i, v := range kss {
		keyLens[i] = v.ks
	}
	return keyLens
}

func chal6(enc []byte, noKeys int) ([]byte, []byte) {
	keyLens := determineKeysize(enc, noKeys)
	dec, keys := BreakRepeatingXOR(enc, keyLens[0])
	return dec, keys
}
