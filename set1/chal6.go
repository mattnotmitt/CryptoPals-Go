package set1

import (
	"CryptoPals/util"
	"fmt"
	"sort"
)

// oh godddddddddd

type Keysize struct {
	ks  int
	nhd float64
}

func chal6(enc []byte, keys int) {
	fmt.Println("Bruteforcing Vigenere...")

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
	kss = kss[:keys]

	fmt.Printf("Optimal keysize: %v\n", kss[0].ks)
	if keys > 1 {
		fmt.Printf("Also checking: %v\n", kss[1:])
	}

	for _, ks := range kss {
		bls := util.ChunkByteArray(enc, ks.ks, false)
		tbls := make([][]byte, ks.ks)

		for i := 0; i < len(bls); i++ {
			for j := 0; j < len(bls[i]); j++ {
				tbls[j] = append(tbls[j], bls[i][j])
			}
		}

		keys := make([]byte, ks.ks)
		for i, tbl := range tbls {
			_, _, k := chal3(tbl)
			keys[i] = k
		}
		fmt.Printf("Best match found, with key: \"%v\":\n", string(keys))
		dec := util.XOR(enc, keys)
		//fmt.Println(enc)
		fmt.Println(string(dec))
	}
}
