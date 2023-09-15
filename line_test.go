package ttgo

import "testing"

func getEmptyLine() line {
	return line{blank, blank, blank}
}

func getFullLine() line {
	return line{x, x, x}
}

func getOneValLine() line {
	return line{x, blank, blank}
}

func TestIsUniform_EmptyLine(t *testing.T) {
	assert(t, getEmptyLine().isUniform())
}

func TestIsNotUniform_WithSingleValue(t *testing.T) {
	assert(t, !getOneValLine().isUniform())
}

func TestIsNotUniform_WithDifferentValues(t *testing.T) {
	l := line{x, o, x}
	assert(t, !l.isUniform())
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
