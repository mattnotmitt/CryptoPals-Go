package set2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChal9(t *testing.T) {
	expected := "YELLOW SUBMARINE\x04\x04\x04\x04"
	assert.Equal(t, expected, chal9("YELLOW SUBMARINE", 20))
}
