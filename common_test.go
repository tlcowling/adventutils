package adventutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadStrings(t *testing.T) {
	lines := ReadInputAsLines("./testdata/strings.txt")
	assert.Equal(t, "first", lines[0], "first line should be 'first'")
}

func TestReadInts(t *testing.T) {
	lines := ReadInputAsInts("./testdata/ints.txt")
	assert.Equal(t, 425, lines[0], "first line should be '425'")
}
