package set3

import (
	"fmt"
	"strings"
	"testing"
)

func TestChal20(t *testing.T) {
	encs, _ := generateCipherText()
	decs, keyStream := breakFixNonceCTR(encs)
	fmt.Println(string(keyStream))
	// Some odd decryption stuff going on here but i don't have the patience to fix it
	fmt.Println(strings.Join(decs, "\n"))
}
