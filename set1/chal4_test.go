package set1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChal4(t *testing.T) {
	str := chal4("data/4.txt")
	assert.Equal(t, "3a5557340be6f5315c35112912393503320f54065f0e275a3b5853352008", str)
}
