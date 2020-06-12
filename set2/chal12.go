package set2

import (
	"CryptoPals/util"
	"bytes"
	"io/ioutil"
	"math/rand"
)

type Oracle func (pt []byte) []byte

func AESOracleStatic (pt []byte) []byte {
	var encrypted []byte
	data, err := ioutil.ReadFile("data/12.txt")
	util.Check(err)
	unknown := util.FromBase64(string(data))

	rand.Seed(1)
	key := util.RandBytes(16)
	pt = append(pt, unknown...)
	encrypted = util.AESECBEncrypt(pt, key)
	return encrypted
}

func determineBlockSize(oracle Oracle) int {
	prevLen := 0
	for i := 1;; i++ {
		result := oracle(bytes.Repeat([]byte("A"), i))
		if len(result) > prevLen {
			if prevLen != 0 {
				return len(result) - prevLen
			}
			prevLen = len(result)
		}
	}
}

func generateByteLookup(oracle Oracle, prefix []byte, start, end int) map[string]byte {
	chunkFreq := make(map[string]byte)
	for b := byte(0); b < 128; b++ {
		chunkFreq[string(oracle(append(prefix, b))[start:end])] = b
	}
	return chunkFreq
}

func Chal12 () []byte {
	var decrypted []byte
	result := AESOracleStatic(bytes.Repeat([]byte("A"), 128))
	size := determineBlockSize(AESOracleStatic)
	_, score := util.DetectECB(result, size)
	if score == 0 {
		panic("Not ECB!")
	}
	emptyCipherLen := len(AESOracleStatic([]byte{}))

	for len(decrypted) < emptyCipherLen {
		blockStart := len(decrypted)
		blockEnd := blockStart + size

		for i := size - 1; i >= 0; i-- {
			pref := bytes.Repeat([]byte("A"), i)
			knownPrefix := append(pref, decrypted...)
			lookup := generateByteLookup(AESOracleStatic, knownPrefix, blockStart, blockEnd)

			block := AESOracleStatic(pref)[blockStart:blockEnd]
			decrypted = append(decrypted, lookup[string(block)])
		}
	}

	return decrypted
}