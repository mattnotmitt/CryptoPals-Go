package set1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChal2(t *testing.T) {
	out := chal2("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	expected := "746865206b696420646f6e277420706c6179"
	fmt.Println(out)
	assert.Equal(t, out, expected)
}
