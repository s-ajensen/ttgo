package ttgo

import "testing"

func getBlankBoard() *board {
	var b board
	return &b
}

func getTestBoard() *board {
	return &board{x, blank, blank, o, blank, blank, blank, x, o}
}

func getTieBoard() *board {
	return &board{x, o, x, x, o, x, o, x, o}
}

func getWinBoard() *board {
	return &board{x, x, x, o, o, x, o, o, x}
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

func TestDeterminesWinByRow(t *testing.T) {
	var unfinishedBoard board
	winBoard := board{x, x, x, blank, o, o, o, blank, blank}

	assert(t, !unfinishedBoard.isWon())
	assert(t, winBoard.isWon())
}

func TestDeterminesWinByColumn(t *testing.T) {
	winBoard := board{x, o, o, x, blank, blank, x, blank, blank}
	assert(t, winBoard.isWon())
}

func TestDeterminesWinByDiags(t *testing.T) {
	winBoardX := board{x, o, o, blank, x, blank, blank, blank, x}
	winBoardO := board{x, x, o, x, o, blank, o, blank, blank}

	assert(t, winBoardX.isWon())
	assert(t, winBoardO.isWon())
}

func TestDoesNotCount_NewGame_AsTie(t *testing.T) {
	assert(t, !getBlankBoard().isTied())
}

func TestDoesNotCount_UnfinishedGame_AsTie(t *testing.T) {
	b := board{x, blank, blank, blank, blank, blank, blank, blank, blank}
	assert(t, !b.isTied())
}

func TestDeterminesTie(t *testing.T) {
	assert(t, getTieBoard().isTied())
}

func TestDoesNotCount_FullWin_AsTie(t *testing.T) {
	b := board{x, o, x, o, x, o, x, o, x}
	assert(t, !b.isTied())
}

func TestCalcsNextPiecePlayed_BlankBoard(t *testing.T) {
	assertEquals(t, getBlankBoard().curPiece(), x)
}

func TestCalcsNextPiecePlayed_SingleMovePlayed(t *testing.T) {
	b := board{x, blank, blank, blank, blank, blank, blank, blank, blank}
	assertEquals(t, b.curPiece(), o)
}

func TestCalcsNextPiecePlayed_ManyMovesPlayed(t *testing.T) {
	xBoard := board{x, o, blank, blank, blank, blank, blank, blank, blank}
	oBoard := board{x, o, x, blank, blank, blank, blank, blank, blank}

	assertEquals(t, xBoard.curPiece(), x)
	assertEquals(t, oBoard.curPiece(), o)
}

func TestOpenPositions_ReturnsAllIndices_BlankBoard(t *testing.T) {
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	assertSliceEquals(t, expected, getBlankBoard().getOpenSpaces())
}

func TestOpenPositiong_ReturnsEmptyArray_FullBoard(t *testing.T) {
	expected := []int{}
	assertSliceEquals(t, expected, getTieBoard().getOpenSpaces())
}

func TestOpenPositions_ReturnsOpen_WithSingleMove(t *testing.T) {
	b := board{x, blank, blank, blank, blank, blank, blank, blank, blank}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}

	assertSliceEquals(t, expected, b.getOpenSpaces())
}

func TestOpenPositiong_ReturnsOpen_WithManyMoves(t *testing.T) {
	b := board{x, o, x, blank, x, blank, blank, blank, o}
	expected := []int{3, 5, 6, 7}

	assertSliceEquals(t, expected, b.getOpenSpaces())
}
