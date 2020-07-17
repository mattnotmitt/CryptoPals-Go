package set1

import (
	"github.com/mattnotmitt/CryptoPals-Go/util"
	"fmt"
	"testing"
)

func TestChal3(t *testing.T) {
	str, score, key := chal3(util.FromHex("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"))
	fmt.Printf( "Best match: \"%s\" with key \"%s\"\n and score \"%f\"", str, string(key), score)
}
