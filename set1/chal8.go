package set1

import (
	"CryptoPals/util"
	"bufio"
	"os"
)

func chal8(inp string) ([]byte, int, float64) {
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
		_, score := util.DetectECB(util.FromHex(data), 16)
		if score > bestScore {
			best = []byte(data)
			bestLine = line
			bestScore = score
		}
	}

	return best, bestLine, bestScore
}
