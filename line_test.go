package ttgo

import "testing"

func getEmptyLine() Line {
	return Line{blank, blank, blank}
}

func getFullLine() Line {
	return Line{x, x, x}
}

func getOneValLine() Line {
	return Line{x, blank, blank}
}

func TestIsUniform_EmptyLine(t *testing.T) {
	assert(t, getEmptyLine().isUniform())
}

func TestIsNotUniform_WithSingleValue(t *testing.T) {
	assert(t, !getOneValLine().isUniform())
}

func TestIsNotUniform_WithDifferentValues(t *testing.T) {
	line := Line{x, o, x}
	assert(t, !line.isUniform())
}

func TestIsUniform_WithFullLine(t *testing.T) {
	assert(t, getFullLine().isUniform())
}

func TestIsEmpty_EmptyLine(t *testing.T) {
	assert(t, getEmptyLine().isEmpty())
}

func TestIsNotEmpty_WithSingleValue(t *testing.T) {
	assert(t, !getOneValLine().isEmpty())
}

func TestIsNotEmpty_FullLine(t *testing.T) {
	assert(t, !getFullLine().isEmpty())
}
