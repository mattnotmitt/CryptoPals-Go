package set3

import (
	"bytes"
	"errors"
	"github.com/mattnotmitt/CryptoPals-go/set2"
	"github.com/mattnotmitt/CryptoPals-go/util"
	"io/ioutil"
	"math/rand"
	"strings"
)

var key17 = util.RandBytes(16)

func encRandom () ([]byte, []byte, []byte) {
	data, err := ioutil.ReadFile("data/17.txt")
	util.Check(err)
	strArray := strings.Split(string(data), "\n")
	randStr := util.FromBase64(strArray[rand.Intn(len(strArray))])
	enc := util.AESCBCEncrypt(randStr, key17, []byte("\x00"))
	return enc, bytes.Repeat([]byte("\x00"), 16), randStr
}

func paddingValid (enc []byte) bool {
	pt := util.AESCBCDecrypt(enc, key17, []byte("\x00"), false)
	_, err := set2.VerifyPadding(string(pt), 16)
	return err == nil
}
// Oracle Implementation
// ==========================================
// Exploit Implementation

func breakOracle(enc []byte, iv []byte) ([]byte, error) {
	var pt []byte
	if !paddingValid(enc) {
		panic("unknown text is invalid")
	}
	chunkEnc := util.ChunkByteArray(enc, 16, false)
	lastBlock := iv
	for _, chunk := range chunkEnc {
		ptBlock, err  := breakBlock(chunk, lastBlock)
		if err != nil {
			return []byte(""), err
		}
		pt = append(pt, ptBlock...)
		lastBlock = chunk
	}
	return pt, nil
}

func breakBlock(block, lastBlock []byte) ([]byte, error) {
	var pt []byte
	for n := range block {
		ptByte, err := breakChar(block, lastBlock, pt, n+1)
		if err != nil {
			return []byte(""), err
		}
		pt = append([]byte{ptByte}, pt...)
	}
	return pt, nil
}

func breakChar(block, lastBlock, knownPt []byte, byteN int) (byte, error) {
	offset := len(lastBlock) - byteN
	encByte := lastBlock[offset]
	blockPref := make([]byte, offset)
	copy(blockPref, lastBlock[:offset])
	var blockSuff []byte
	for i, v := range knownPt {
		blockSuff = append(blockSuff, v ^ byte(byteN) ^ lastBlock[(len(lastBlock)) - (len(knownPt) - i)])
	}
	for b := byte(0); ; b++ {
		nBlock := append(append(append(blockPref, b), blockSuff...), block...)
		if paddingValid(nBlock) {
			if b == encByte && byteN == 1 {
				continue
			}

			return b ^ encByte ^ byte(byteN), nil
		}
		if b == 255 {
			break
		}
	}
	return 0, errors.New("no valid byte found")
}