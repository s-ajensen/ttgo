package ttgo

import "testing"

func TestIdentifiesRows(t *testing.T) {
	board := board{x, b, b, o, b, b, b, x, o}
	rows := board.getRows()

	assertSliceEquals(t, line{x, b, b}, rows[0])
	assertSliceEquals(t, line{o, b, b}, rows[1])
	assertSliceEquals(t, line{b, x, o}, rows[2])
}

func TestIdentifiesCols(t *testing.T) {
	board := board{x, b, b, o, b, b, b, x, o}
	cols := board.getCols()

	assertSliceEquals(t, line{x, o, b}, cols[0])
	assertSliceEquals(t, line{b, b, x}, cols[1])
	assertSliceEquals(t, line{b, b, o}, cols[2])
}
