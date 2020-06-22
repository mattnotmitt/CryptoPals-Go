package set2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChal11(t *testing.T) {
	det, choice := chal11()
	assert.Equal(t, choice, det)
}
