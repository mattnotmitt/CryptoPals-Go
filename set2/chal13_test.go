package set2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChal13(t *testing.T) {
	// KVParser Tests
	expected := map[string]string{
		"foo": "bar",
		"baz": "qux",
		"zap": "zazzle",
	}
	actual := KVParser("foo=bar&baz=qux&zap=zazzle")
	assert.Equal(t, expected, actual)

	// ProfileFrom Tests
}