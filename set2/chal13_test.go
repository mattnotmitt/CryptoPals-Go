package set2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChal13(t *testing.T) {
	// KVParser Tests
	kvExpected := map[string]string{
		"foo": "bar",
		"baz": "qux",
		"zap": "zazzle",
	}
	kvActual := kvParser("foo=bar&baz=qux&zap=zazzle")
	assert.Equal(t, kvExpected, kvActual)

	// ProfileFor Tests
	pfExpected := map[string]string{
		"email": "foo@bar.comroleadmin",
		"uid":   "10",
		"role":  "user",
	}
	pfActual := profileFor("foo@bar.com&role=admin")
	assert.Equal(t, pfExpected, pfActual)

	// Login Tests
	encProf := login("foo@bar.com") // Compare 2 runs
	encProf2 := login("foo@bar.com")
	assert.Equal(t, encProf, encProf2)

	// VerifyCookie Tests
	vcExpected := map[string]string{
		"email": "foo@bar.com",
		"uid":   "10",
		"role":  "user",
	}
	decProf := verifyCookie(encProf2)
	assert.Equal(t, vcExpected, decProf)
}

func TestBreakECB(t *testing.T) {
	// blocks are 16 bytes long - so we want to generate two encrypted profiles
	// - one where the second block is "admin" padded out to 16 bytes
	// - one where the final block contains "user" padded out to 16 bytes
	// then we simply remove the last block of the second and replace it with the
	// second block of the first profile
	normalProf := login("a@aaaaa.co.uk")
	weirdProf := login("a@aaaaa.euadmin\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b")
	adminProfEnc := append(normalProf[:32], weirdProf[16:32]...)
	adminProf := verifyCookie(adminProfEnc)
	assert.Equal(t, "admin", adminProf["role"])
}
