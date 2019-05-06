package set1

import (
	"CryptoPals/util"
	"bufio"
	"os"
)

func detectECB (inp []byte) (map[string]int, float64) {
	chunks := util.ChunkByteArray(inp, 16)
	chunkFreq := make(map[string]int)
	repeats := 0.0
	for _, chunk := range chunks {
		if _, ok := chunkFreq[string(chunk)]; ok {
			chunkFreq[string(chunk)]++
			repeats++
		} else {
			chunkFreq[string(chunk)] = 1
		}
	}
	return chunkFreq, repeats
}

func Chal8 (inp string) ([]byte, int, float64) {
	file, err := os.Open(inp)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var best []byte
	bestLine := 0
	bestScore := 0.0

	scanner := bufio.NewScanner(file)
	for line := 0; scanner.Scan(); line++ {
		data := scanner.Text()
		_, score := detectECB(util.FromHex(data))
		if score > bestScore {
			best = []byte(data)
			bestLine = line
			bestScore = score
		}
	}

	return best, bestLine, bestScore
}