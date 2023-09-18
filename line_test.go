package ttgo

import (
	"testing"
	. "ttgo/assert"
)

func getEmptyLine() Line {
	return Line{Blank, Blank, Blank}
}

func getFullLine() Line {
	return Line{X, X, X}
}

func getOneValLine() Line {
	return Line{X, Blank, Blank}
}

func TestIsUniform_EmptyLine(t *testing.T) {
	Assert(t, getEmptyLine().isUniform())
}

func TestIsNotUniform_WithSingleValue(t *testing.T) {
	Assert(t, !getOneValLine().isUniform())
}

func TestIsNotUniform_WithDifferentValues(t *testing.T) {
	line := Line{X, O, X}
	Assert(t, !line.isUniform())
}

func TestIsUniform_WithFullLine(t *testing.T) {
	Assert(t, getFullLine().isUniform())
}

func TestIsEmpty_EmptyLine(t *testing.T) {
	Assert(t, getEmptyLine().isEmpty())
}

func TestIsNotEmpty_WithSingleValue(t *testing.T) {
	Assert(t, !getOneValLine().isEmpty())
}

func TestIsNotEmpty_FullLine(t *testing.T) {
	Assert(t, !getFullLine().isEmpty())
}
