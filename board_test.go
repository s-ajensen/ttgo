package ttgo

import "testing"

func getTestBoard() *board {
	return &board{x, b, b, o, b, b, b, x, o}
}

func TestIdentifiesRows(t *testing.T) {
	rows := getTestBoard().getRows()

	assertSliceEquals(t, line{x, b, b}, rows[0])
	assertSliceEquals(t, line{o, b, b}, rows[1])
	assertSliceEquals(t, line{b, x, o}, rows[2])
}

func TestIdentifiesCols(t *testing.T) {
	cols := getTestBoard().getCols()

	assertSliceEquals(t, line{x, o, b}, cols[0])
	assertSliceEquals(t, line{b, b, x}, cols[1])
	assertSliceEquals(t, line{b, b, o}, cols[2])
}

func TestIdentifiesDiags(t *testing.T) {
	diags := getTestBoard().getDiags()

	assertSliceEquals(t, line{x, b, o}, diags[0])
	assertSliceEquals(t, line{b, b, b}, diags[1])
}

func TestIsUniformLine(t *testing.T) {
	uniformLine := line{x, x, x}
	nonUniformLine := line{x, o, x}

	assert(t, uniformLine.isUniform())
	assert(t, !nonUniformLine.isUniform())
}

func TestIsEmptyLine(t *testing.T) {
	emptyLine := line{b, b, b}
	populatedLine := line{x, b, b}
	fullLine := line{x, x, o}

	assert(t, emptyLine.isEmpty())
	assert(t, !populatedLine.isEmpty())
	assert(t, !fullLine.isEmpty())
}
