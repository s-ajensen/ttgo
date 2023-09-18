package ttgo

import (
	"testing"
	. "ttgo/assert"
)

func getBlankBoard() *Board {
	var board Board
	return &board
}

func getTestBoard() *Board {
	return &Board{X, Blank, Blank, O, Blank, Blank, Blank, X, O}
}

func getTieBoard() *Board {
	return &Board{X, O, X, X, O, X, O, X, O}
}

func getWinBoard() *Board {
	return &Board{X, X, X, O, O, X, O, O, X}
}

func TestStringsEmptyBoard(t *testing.T) {
	boardStr := getBlankBoard().String()
	AssertEquals(t, "- - -\n- - -\n- - -\n", boardStr)
}

func TestStringsBoardWithSingleMove(t *testing.T) {
	board := getBlankBoard().Move(0, X)
	boardStr := (&board).String()

	AssertEquals(t, "X - -\n- - -\n- - -\n", boardStr)
}

func TestStringsBoardWithManyMoves(t *testing.T) {
	boardStr := getTestBoard().String()
	AssertEquals(t, "X - -\nO - -\n- X O\n", boardStr)
}

func TestIdentifiesRows(t *testing.T) {
	rows := getTestBoard().getRows()

	AssertSliceEquals(t, Line{X, Blank, Blank}, rows[0])
	AssertSliceEquals(t, Line{O, Blank, Blank}, rows[1])
	AssertSliceEquals(t, Line{Blank, X, O}, rows[2])
}

func TestIdentifiesCols(t *testing.T) {
	cols := getTestBoard().getCols()

	AssertSliceEquals(t, Line{X, O, Blank}, cols[0])
	AssertSliceEquals(t, Line{Blank, Blank, X}, cols[1])
	AssertSliceEquals(t, Line{Blank, Blank, O}, cols[2])
}

func TestIdentifiesDiags(t *testing.T) {
	diags := getTestBoard().getDiags()

	AssertSliceEquals(t, Line{X, Blank, O}, diags[0])
	AssertSliceEquals(t, Line{Blank, Blank, Blank}, diags[1])
}

func TestDeterminesWinByRow(t *testing.T) {
	var unfinishedBoard Board
	winBoard := Board{X, X, X, Blank, O, O, O, Blank, Blank}

	Assert(t, !unfinishedBoard.isWon())
	Assert(t, winBoard.isWon())
}

func TestDeterminesWinByColumn(t *testing.T) {
	winBoard := Board{X, O, O, X, Blank, Blank, X, Blank, Blank}
	Assert(t, winBoard.isWon())
}

func TestDeterminesWinByDiags(t *testing.T) {
	winBoardX := Board{X, O, O, Blank, X, Blank, Blank, Blank, X}
	winBoardO := Board{X, X, O, X, O, Blank, O, Blank, Blank}

	Assert(t, winBoardX.isWon())
	Assert(t, winBoardO.isWon())
}

func TestDoesNotCount_NewGame_AsTie(t *testing.T) {
	Assert(t, !getBlankBoard().isTied())
}

func TestDoesNotCount_UnfinishedGame_AsTie(t *testing.T) {
	board := Board{X, Blank, Blank, Blank, Blank, Blank, Blank, Blank, Blank}
	Assert(t, !board.isTied())
}

func TestDeterminesTie(t *testing.T) {
	Assert(t, getTieBoard().isTied())
}

func TestDoesNotCount_FullWin_AsTie(t *testing.T) {
	board := Board{X, O, X, O, X, O, X, O, X}
	Assert(t, !board.isTied())
}

func TestCalcsNextPiecePlayed_BlankBoard(t *testing.T) {
	AssertEquals(t, getBlankBoard().CurPiece(), X)
}

func TestCalcsNextPiecePlayed_SingleMovePlayed(t *testing.T) {
	board := Board{X, Blank, Blank, Blank, Blank, Blank, Blank, Blank, Blank}
	AssertEquals(t, board.CurPiece(), O)
}

func TestCalcsNextPiecePlayed_ManyMovesPlayed(t *testing.T) {
	xBoard := Board{X, O, Blank, Blank, Blank, Blank, Blank, Blank, Blank}
	oBoard := Board{X, O, X, Blank, Blank, Blank, Blank, Blank, Blank}

	AssertEquals(t, xBoard.CurPiece(), X)
	AssertEquals(t, oBoard.CurPiece(), O)
}

func TestOpenPositions_ReturnsAllIndices_BlankBoard(t *testing.T) {
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	AssertSliceEquals(t, expected, getBlankBoard().getOpenSpaces())
}

func TestOpenPositiong_ReturnsEmptyArray_FullBoard(t *testing.T) {
	expected := []int{}
	AssertSliceEquals(t, expected, getTieBoard().getOpenSpaces())
}

func TestOpenPositions_ReturnsOpen_WithSingleMove(t *testing.T) {
	board := Board{X, Blank, Blank, Blank, Blank, Blank, Blank, Blank, Blank}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8}

	AssertSliceEquals(t, expected, board.getOpenSpaces())
}

func TestOpenPositiong_ReturnsOpen_WithManyMoves(t *testing.T) {
	board := Board{X, O, X, Blank, X, Blank, Blank, Blank, O}
	expected := []int{3, 5, 6, 7}

	AssertSliceEquals(t, expected, board.getOpenSpaces())
}
