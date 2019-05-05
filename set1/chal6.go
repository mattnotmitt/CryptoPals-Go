package set1

import (
	"CryptoPals/util"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
)

// oh godddddddddd

type Keysize struct {
	ks int
	nhd float64
}

func Chal6 (fn string) {
	fmt.Println("Bruteforcing Vigenere...")

	data, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}
	enc, err := base64.StdEncoding.DecodeString(string(data))


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
		kss[ks - 2] = Keysize{ks,nhd}

	}
	sort.Slice(kss, func(i,j int) bool { return kss[i].nhd < kss[j].nhd })
	kss = kss[:1]

	fmt.Printf("Optimal keysize: %v\n", kss[0].ks)

	for _, ks := range kss {
		var bls [][]byte

		for i := 0; i < len(enc); i += ks.ks {
			end := i + ks.ks
			if end > len(enc) {
				end = len(enc)
			}

			bls = append(bls, enc[i:end])
		}
		tbls := make([][]byte, ks.ks)

		for i := 0; i < len(bls); i++ {
			for j := 0; j < len(bls[i]); j++ {
				tbls[j] = append(tbls[j], bls[i][j])
			}
		}

		keys := make([]byte, ks.ks)
		for i, tbl := range tbls {
			_, _, k := Chal3(tbl)
			keys[i] = k
		}
		fmt.Printf("Best match found, with key: \"%v\":\n", string(keys))
		fmt.Println(string(util.XOR(enc, keys)))
	}
}