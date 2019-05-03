package util

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
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
		key = append(key, bytes.Repeat(key, int(math.Ceil(math.Abs(float64(diff))/float64(len(key)))))...)
	}
	res := make([]byte, int(math.Max(float64(len(inp)), float64(len(key)))))
	for i := 0; i < len(inp); i++ {
		res[i] = inp[i] ^ key[i]
	}

	return res
}

func ScoreString (inp string) float64 {
	rawInp := []byte(inp)
	counts := make([]int, 256)
	for _, b := range rawInp {
		counts[b]++
	}

	engFreq := []float64{0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.755, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 15.843, 0.004, 0.375, 0.002, 0.008, 0.019, 0.008, 0.134, 0.137, 0.137, 0.001, 0.001, 0.972, 0.19, 0.857, 0.017, 0.334, 0.421, 0.246, 0.108, 0.104, 0.112, 0.103, 0.1, 0.127, 0.237, 0.04, 0.027, 0.004, 0.003, 0.004, 0.002, 0.0001, 0.338, 0.218, 0.326, 0.163, 0.121, 0.149, 0.133, 0.192, 0.232, 0.107, 0.082, 0.148, 0.248, 0.134, 0.103, 0.195, 0.012, 0.162, 0.368, 0.366, 0.077, 0.061, 0.127, 0.009, 0.03, 0.015, 0.004, 0.0001, 0.004, 0.0001, 0.003, 0.0001, 6.614, 1.039, 2.327, 2.934, 9.162, 1.606, 1.415, 3.503, 5.718, 0.081, 0.461, 3.153, 1.793, 5.723, 5.565, 1.415, 0.066, 5.036, 4.79, 6.284, 1.992, 0.759, 1.176, 0.139, 1.162, 0.102, 0.0001, 0.002, 0.0001, 0.0001, 0.0001, 0.06, 0.004, 0.003, 0.002, 0.001, 0.001, 0.001, 0.002, 0.001, 0.001, 0.0001, 0.001, 0.001, 0.003, 0.0001, 0.0001, 0.001, 0.001, 0.001, 0.031, 0.006, 0.001, 0.001, 0.001, 0.002, 0.014, 0.001, 0.001, 0.005, 0.005, 0.001, 0.002, 0.017, 0.007, 0.002, 0.003, 0.004, 0.002, 0.001, 0.002, 0.002, 0.012, 0.001, 0.002, 0.001, 0.004, 0.001, 0.001, 0.003, 0.003, 0.002, 0.005, 0.001, 0.001, 0.003, 0.001, 0.003, 0.001, 0.002, 0.001, 0.004, 0.001, 0.002, 0.001, 0.0001, 0.0001, 0.02, 0.047, 0.009, 0.009, 0.0001, 0.0001, 0.001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.003, 0.001, 0.004, 0.002, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.001, 0.001, 0.001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.005, 0.002, 0.061, 0.001, 0.0001, 0.002, 0.001, 0.001, 0.001, 0.001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001, 0.0001}
	freqDist := make([]float64, 256)

	for i, c := range counts {
		if c > 0 {
			freqDist[i] = float64(c) / float64(len(rawInp))
		}
	}

	score := stat.ChiSquare(freqDist, engFreq)
	df := float64(len(rawInp) - 1)

	return 1 - distuv.ChiSquared{K: df}.CDF(score)
}
