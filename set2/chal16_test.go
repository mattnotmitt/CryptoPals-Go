package set2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChal16 (t *testing.T) {
	notAdmin := vanillaIceify("xx:admin<true:xx")
	assert.Equal(t, false, checkAdmin(notAdmin))
	// Turn on LSB of bits aligned with the colons
	// and lt sign in previous block to change to
	// semicolon and equals sign respectively
	// Completely corrupts previous block but ¯\_(ツ)_/¯
	notAdmin[18] ^= 0b01
	notAdmin[24] ^= 0b01
	notAdmin[29] ^= 0b01
	assert.Equal(t, true, checkAdmin(notAdmin))
}

// comment1=cooking
// %20MCs;userdata=
// xx:user<admin:xx
// ;comment2=%20lik
// e%20a%20pound%20
// of%20bacon