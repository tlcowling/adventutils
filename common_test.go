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
func TestParseCommaSeparatedInts(t *testing.T) {
	lines := ParseCommaSeparatedInts("1,3,4,5,6")
	assert.Equal(t, []int{1, 3,4, 5,6}, lines, "parse comma separated ints'")
}
func TestParseCommaSeparatedIntsFromFile(t *testing.T) {
	lines := ParseCommaSeparatedIntsFromFile("./testdata/csv_ints.txt")
	assert.Equal(t, []int{14, 3,4,34,354}, lines, "parse comma separated ints from file'")
}
