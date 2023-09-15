package ttgo

import "testing"

func emptyLine() line {
	return line{blank, blank, blank}
}

func fullLine() line {
	return line{x, x, x}
}

func oneValLine() line {
	return line{x, blank, blank}
}

func TestIsUniform_EmptyLine(t *testing.T) {
	assert(t, emptyLine().isUniform())
}

func TestIsNotUniform_WithSingleValue(t *testing.T) {
	assert(t, !oneValLine().isUniform())
}

func TestIsNotUniform_WithDifferentValues(t *testing.T) {
	l := line{x, o, x}
	assert(t, !l.isUniform())
}

func TestIsUniform_WithFullLine(t *testing.T) {
	assert(t, fullLine().isUniform())
}

func TestIsEmpty_EmptyLine(t *testing.T) {
	assert(t, emptyLine().isEmpty())
}

func TestIsNotEmpty_WithSingleValue(t *testing.T) {
	assert(t, !oneValLine().isEmpty())
}

func TestIsNotEmpty_FullLine(t *testing.T) {
	assert(t, !fullLine().isEmpty())
}
