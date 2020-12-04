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
func TestStringsToInts(t *testing.T) {
	lines := StringArrayToIntArray([]string{"1", "4", "5", "2049"})
	assert.Equal(t, []int{1, 4, 5, 2049}, lines, "strings should be ints'")
}
