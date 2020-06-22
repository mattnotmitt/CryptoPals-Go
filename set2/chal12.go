package set2

import (
	"CryptoPals/util"
	"bytes"
	"io/ioutil"
	"sync"
)

type Oracle func(pt []byte) []byte

var setup12 sync.Once
var key12 []byte
var unknown12 []byte

func aesOracleStatic(pt []byte) []byte {
	var encrypted []byte

	setup12.Do(func() {
		data, err := ioutil.ReadFile("data/12.txt")
		util.Check(err)
		unknown12 = util.FromBase64(string(data))
		key12 = util.RandBytes(16)
	}) // Generate key on first run of program and persist
	// Load key from file
	pt = append(pt, unknown12...)
	encrypted = util.AESECBEncrypt(pt, key12)
	return encrypted
}

func determineBlockSize(oracle Oracle) int {
	prevLen := 0
	for i := 1; ; i++ {
		result := oracle(bytes.Repeat([]byte("A"), i))
		if len(result) > prevLen {
			if prevLen != 0 {
				return len(result) - prevLen
			}
			prevLen = len(result)
		}
	}
}

func generateByteLookup(oracle Oracle, prefix []byte, start, end int, c chan map[string]byte) {
	chunkFreq := make(map[string]byte)
	for b := byte(0); b < 128; b++ {
		chunkFreq[string(oracle(append(prefix, b))[start:end])] = b
	}
	c <- chunkFreq
}

func chal12() []byte {
	var decrypted []byte
	result := aesOracleStatic(bytes.Repeat([]byte("A"), 128))
	size := determineBlockSize(aesOracleStatic)
	_, score := util.DetectECB(result, size)
	if score == 0 {
		panic("Not ECB!")
	}
	emptyCipherLen := len(aesOracleStatic([]byte{}))

	for len(decrypted) < emptyCipherLen {
		blockStart := len(decrypted)
		blockEnd := blockStart + size

		for i := size - 1; i >= 0; i-- {
			pref := bytes.Repeat([]byte("A"), i)
			knownPrefix := append(pref, decrypted...)

			lookupC := make(chan map[string]byte)
			go generateByteLookup(aesOracleStatic, knownPrefix, blockStart, blockEnd, lookupC)

			block := aesOracleStatic(pref)[blockStart:blockEnd]
			lookup := <-lookupC
			decrypted = append(decrypted, lookup[string(block)])
		}
	}

	return bytes.TrimRight(decrypted, "\x04")
}
