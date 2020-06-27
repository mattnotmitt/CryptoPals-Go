package set2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChal9(t *testing.T) {
	expected0 := "YELLOW SUBMARINE\x04\x04\x04\x04"
	assert.Equal(t, expected0, chal9("YELLOW SUBMARINE", 20))
	expected1 := "YELLOW SUBMARINE\x10\u0010\u0010\u0010\u0010\u0010\u0010\u0010\u0010\u0010\u0010\u0010\u0010\u0010\u0010\u0010"
	assert.Equal(t, expected1, chal9("YELLOW SUBMARINE", 16))
	expected2 := "YELLOW SUBMARIN\x01"
	assert.Equal(t, expected2, chal9("YELLOW SUBMARIN", 16))
}
