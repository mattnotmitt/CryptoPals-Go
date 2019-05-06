package set1

import (
	"fmt"
	"testing"
)

func TestChal8(t *testing.T) {
	bd, bl, bs := Chal8("data/8.txt")
	fmt.Printf("Line most likely to be ECB is %v, with %v repeats:\n%v\n", bl, bs, string(bd))
}