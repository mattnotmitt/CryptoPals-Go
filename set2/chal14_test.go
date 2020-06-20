package set2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChal14(t *testing.T) {
	expected := "Rollin' in my 5.0\nWith my rag-top down so my hair can blow\nThe girlies on standby waving just to say hi\nDid you stop? No, I just drove by\n"
	assert.Equal(t, expected, string(Chal14()))
}