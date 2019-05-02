package util

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"log"
	"math"
)

func ToHex (inp []byte) string {
	out := make([]byte, hex.EncodedLen(len(inp)))
	n := hex.Encode(out, inp)
	return string(out[:n])
}

func FromHex (inp string) []byte {
	rawInp := []byte(inp)
	out := make([]byte, hex.DecodedLen(len(rawInp)))
	n, err := hex.Decode(out, rawInp)
	if err != nil {
		log.Fatal(err)
	}
	return out[:n]
}

func ToBase64 (inp []byte) string {
	buf := new(bytes.Buffer)

	encoder := base64.NewEncoder(base64.StdEncoding, buf)
	defer encoder.Close()
	_, err := encoder.Write(inp)
	if err != nil {
		log.Fatal(err)
	}

	return buf.String()
}

func XOR (inp, key []byte) []byte {
	diff := len(key) - len(inp)

	if diff < 0 {
		key = append(key, bytes.Repeat([]byte("\x00"), int(math.Abs(float64(diff))))...)
	}
	res := make([]byte, int(math.Max(float64(len(inp)), float64(len(key)))))
	for i := 0; i < len(inp); i++ {
		res[i] = inp[i] ^ key[i]
	}

	return res
}
