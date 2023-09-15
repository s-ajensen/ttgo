package ttgo

import "testing"

func getTestBoard() *board {
	return &board{x, blank, blank, o, blank, blank, blank, x, o}
}

func TestIdentifiesRows(t *testing.T) {
	rows := getTestBoard().getRows()

	assertSliceEquals(t, line{x, blank, blank}, rows[0])
	assertSliceEquals(t, line{o, blank, blank}, rows[1])
	assertSliceEquals(t, line{blank, x, o}, rows[2])
}

func TestIdentifiesCols(t *testing.T) {
	cols := getTestBoard().getCols()

	assertSliceEquals(t, line{x, o, blank}, cols[0])
	assertSliceEquals(t, line{blank, blank, x}, cols[1])
	assertSliceEquals(t, line{blank, blank, o}, cols[2])
}

func TestIdentifiesDiags(t *testing.T) {
	diags := getTestBoard().getDiags()

	assertSliceEquals(t, line{x, blank, o}, diags[0])
	assertSliceEquals(t, line{blank, blank, blank}, diags[1])
}

func TestIsUniformLine(t *testing.T) {
	uniformLine := line{x, x, x}
	nonUniformLine := line{x, o, x}

	assert(t, uniformLine.isUniform())
	assert(t, !nonUniformLine.isUniform())
}

func TestIsEmptyLine(t *testing.T) {
	emptyLine := line{blank, blank, blank}
	populatedLine := line{x, blank, blank}
	fullLine := line{x, x, o}

	assert(t, emptyLine.isEmpty())
	assert(t, !populatedLine.isEmpty())
	assert(t, !fullLine.isEmpty())
}
