package set1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChal1(t *testing.T) {
	out := Chal1("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	fmt.Println(out)
	assert.Equal(t, out, expected)
}