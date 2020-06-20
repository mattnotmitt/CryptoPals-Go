package set2

import (
	"CryptoPals/util"
	"github.com/stretchr/testify/assert"
	"testing"
)



func TestChal15(t *testing.T) {
	// Test padding stripping
	input0 := "ICE ICE BABY\x04\x04\x04\x04"
	expected0 := "ICE ICE BABY"
	resp0, err0 := util.VerifyPadding(input0, 16)
	assert.Nil(t, err0)
	assert.Equal(t, expected0, resp0)
	// Rest of tests in ../util/utils_test.go
}
