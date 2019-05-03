package set1

import (
	"bufio"
	"log"
	"os"
)

func Chal4 (inp string) string {
	file, err := os.Open(inp)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var best string
	bestScore := 0.0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bestDec, locScore, _ := Chal3(scanner.Text())
		if locScore > bestScore {
			best = bestDec
			bestScore = locScore
		}
	}

	return best
}