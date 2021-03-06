package set2

import (
	"github.com/mattnotmitt/CryptoPals-Go/util"
	"bytes"
	"io/ioutil"
	"math/rand"
	"sync"
	"time"
)

var setup14 sync.Once
var key14 []byte
var randPref []byte

func aesOracleStaticRand(pt []byte) []byte {
	var encrypted []byte
	data, err := ioutil.ReadFile("data/12.txt")
	util.Check(err)
	unknown := util.FromBase64(string(data))

	setup14.Do(func() {
		key14 = util.RandBytes(16) // Generate key on first run of program and persist
		rand.Seed(time.Now().UnixNano())
		byteNum := rand.Intn(256)
		randPref = util.RandBytes(byteNum) // Generate a random number of random bytes
	})
	pt = append(randPref, pt...) // Add random prefix
	pt = append(pt, unknown...)  // Add unknown bytes
	encrypted = util.AESECBEncrypt(pt, key14)
	return encrypted
}

/*
func byteSliceIndex(chunkedSlice [][]byte, chunkMatch []byte) int {
	for n, chunk := range chunkedSlice {
		if bytes.Equal(chunk, chunkMatch) {
			return n
		}
	}
	return -1
}*/

func determineRandLen(oracle Oracle, size int) int {
	//var lastIndex = 0
	var lastPos int
	for i := size * 3; i >= 0; i-- {
		result := oracle(bytes.Repeat([]byte("A"), i))
		chunkResult := util.ChunkByteArray(result, size, false)
		lastIndex := 0
		for n := 1; n < len(chunkResult); n++ {
			if bytes.Equal(chunkResult[lastIndex], chunkResult[n]) {
				lastPos = lastIndex
				break
			}
			lastIndex++
		}
		if lastIndex == len(chunkResult)-1 {
			return lastPos*size + (32 - (i + 1))
		}
	}
	return -1
}

func chal14() []byte {
	var decrypted []byte
	result := aesOracleStaticRand(bytes.Repeat([]byte("A"), 128))
	size := determineBlockSize(aesOracleStaticRand)
	randSize := determineRandLen(aesOracleStaticRand, size)
	if randSize == -1 {
		panic("Couldn't determine length of random string")
	}
	_, score := util.DetectECB(result, size)
	if score == 0 {
		panic("Not ECB!")
	}
	emptyCipherLen := len(aesOracleStaticRand([]byte{}))
	randPad := size - randSize % size
	targetLen :=  emptyCipherLen - randSize
	for len(decrypted) < targetLen {
		blockStart := len(decrypted) + randSize + randPad
		blockEnd := blockStart + size
		for i := size - 1; i >= 0; i-- {
			pref := bytes.Repeat([]byte("A"), i + randPad)
			knownPrefix := append(pref, decrypted...)

			lookupC := make(chan map[string]byte)
			go generateByteLookup(aesOracleStaticRand, knownPrefix, blockStart, blockEnd, lookupC)
			prefEnc := aesOracleStaticRand(pref)
			block := prefEnc[blockStart:blockEnd]
			lookup := <-lookupC
			decrypted = append(decrypted, lookup[string(block)])
			if len(decrypted) == targetLen {
				break
			}
		}
	}
	decrypted = bytes.TrimRight(decrypted, "\x00")
	return util.UnPKCS7(decrypted)
}
