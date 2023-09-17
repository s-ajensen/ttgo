package ttgo

import "testing"

func getBlankBoard() *Board {
	var board Board
	return &board
}

func getTestBoard() *Board {
	return &Board{x, blank, blank, o, blank, blank, blank, x, o}
}

func getTieBoard() *Board {
	return &Board{x, o, x, x, o, x, o, x, o}
}

func getWinBoard() *Board {
	return &Board{x, x, x, o, o, x, o, o, x}
}

func TestIdentifiesRows(t *testing.T) {
	rows := getTestBoard().getRows()

	assertSliceEquals(t, Line{x, blank, blank}, rows[0])
	assertSliceEquals(t, Line{o, blank, blank}, rows[1])
	assertSliceEquals(t, Line{blank, x, o}, rows[2])
}

func TestIdentifiesCols(t *testing.T) {
	cols := getTestBoard().getCols()

	assertSliceEquals(t, Line{x, o, blank}, cols[0])
	assertSliceEquals(t, Line{blank, blank, x}, cols[1])
	assertSliceEquals(t, Line{blank, blank, o}, cols[2])
}

func TestIdentifiesDiags(t *testing.T) {
	diags := getTestBoard().getDiags()

	assertSliceEquals(t, Line{x, blank, o}, diags[0])
	assertSliceEquals(t, Line{blank, blank, blank}, diags[1])
}

func TestDeterminesWinByRow(t *testing.T) {
	var unfinishedBoard Board
	winBoard := Board{x, x, x, blank, o, o, o, blank, blank}

	assert(t, !unfinishedBoard.isWon())
	assert(t, winBoard.isWon())
}

func TestDeterminesWinByColumn(t *testing.T) {
	winBoard := Board{x, o, o, x, blank, blank, x, blank, blank}
	assert(t, winBoard.isWon())
}

func TestDeterminesWinByDiags(t *testing.T) {
	winBoardX := Board{x, o, o, blank, x, blank, blank, blank, x}
	winBoardO := Board{x, x, o, x, o, blank, o, blank, blank}

	assert(t, winBoardX.isWon())
	assert(t, winBoardO.isWon())
}

func TestDoesNotCount_NewGame_AsTie(t *testing.T) {
	assert(t, !getBlankBoard().isTied())
}

func TestDoesNotCount_UnfinishedGame_AsTie(t *testing.T) {
	board := Board{x, blank, blank, blank, blank, blank, blank, blank, blank}
	assert(t, !board.isTied())
}

func TestDeterminesTie(t *testing.T) {
	assert(t, getTieBoard().isTied())
}

func TestDoesNotCount_FullWin_AsTie(t *testing.T) {
	board := Board{x, o, x, o, x, o, x, o, x}
	assert(t, !board.isTied())
}

func TestCalcsNextPiecePlayed_BlankBoard(t *testing.T) {
	assertEquals(t, getBlankBoard().curPiece(), x)
}

func TestCalcsNextPiecePlayed_SingleMovePlayed(t *testing.T) {
	board := Board{x, blank, blank, blank, blank, blank, blank, blank, blank}
	assertEquals(t, board.curPiece(), o)
}

func TestCalcsNextPiecePlayed_ManyMovesPlayed(t *testing.T) {
	xBoard := Board{x, o, blank, blank, blank, blank, blank, blank, blank}
	oBoard := Board{x, o, x, blank, blank, blank, blank, blank, blank}

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
	board := Board{x, blank, blank, blank, blank, blank, blank, blank, blank}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}

	assertSliceEquals(t, expected, board.getOpenSpaces())
}

func TestOpenPositiong_ReturnsOpen_WithManyMoves(t *testing.T) {
	board := Board{x, o, x, blank, x, blank, blank, blank, o}
	expected := []int{3, 5, 6, 7}

	assertSliceEquals(t, expected, board.getOpenSpaces())
}
