package set1

import (
	"CryptoPals/util"
	"fmt"
	"testing"
)

func TestChal3(t *testing.T) {
	str, _, key := chal3(util.FromHex("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"))
	fmt.Printf("Best match: \"%s\" with key \"%s\"\n", str, string(key))
}
