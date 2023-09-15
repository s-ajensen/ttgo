package ttgo

import "testing"

func TestIdentifiesRows(t *testing.T) {
	board := board{x, b, b, o, b, b, b, x, o}
	rows := board.getRows()

	assertSliceEquals(t, line{x, b, b}, rows[0])
	assertSliceEquals(t, line{o, b, b}, rows[1])
	assertSliceEquals(t, line{b, x, o}, rows[2])
}
