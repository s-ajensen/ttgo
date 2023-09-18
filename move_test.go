package ttgo

import (
	"testing"
	. "ttgo/assert"
)

func canLose(board Board, playingAs string) bool {
	resp := NextBoard(&board)
	if resp.isGameOver() {
		return playingAs != resp.getWinner() && "Tie" != resp.getWinner()
	}
	isLoss := false
	for _, space := range resp.getOpenSpaces() {
		newBoard := resp.Move(space, resp.CurPiece())
		if newBoard.isGameOver() {
			return playingAs != newBoard.getWinner() && "Tie" != newBoard.getWinner()
		}
		isLoss = isLoss || canLose(newBoard, playingAs)
	}
	return isLoss
}

func TestMovePlayedOnBlankBoard(t *testing.T) {
	var blankBoard Board
	var playedBoard = blankBoard.Move(0, X)

	AssertEquals(t, X, playedBoard[0])
	AssertNotEquals(t, X, blankBoard[0])
}

func TestEvaluatesTieAsZero_WhenMaximizing(t *testing.T) {
	AssertEquals(t, 0, getTieBoard().evalStatic(true, 0))
}

func TestEvaluatesTieAsZero_WhenMinimizing(t *testing.T) {
	AssertEquals(t, 0, getTieBoard().evalStatic(false, 0))
}

func TestEvaluatesWinAsTen_WhenMaximizing(t *testing.T) {
	AssertEquals(t, 10, getWinBoard().evalStatic(true, 0))
}

func TestEvaluatesWinAsMinusTen_WhenMinimizing(t *testing.T) {
	AssertEquals(t, -10, getWinBoard().evalStatic(false, 0))
}

func TestMinimaxReturnsStaticEval_WhenGameOver(t *testing.T) {
	AssertEquals(t, 0, minimax(*getTieBoard(), 0, true))
	AssertEquals(t, 10, minimax(*getWinBoard(), 0, true))
	AssertEquals(t, -10, minimax(*getWinBoard(), 0, false))
}

func TestMinimaxPrefersWin_ToTie(t *testing.T) {
	board := Board{X, O, X, X, O, X, O, Blank, Blank}
	winMove := board.Move(7, O)
	tieMove := board.Move(8, O)

	Assert(t, minimax(winMove, 0, true) > minimax(tieMove, 0, true))
}

func TestMinimaxIdentifies_PlayerWin(t *testing.T) {
	board := Board{X, O, X, X, O, X, O, Blank, Blank}
	playerWinMove := board.Move(7, O)
	playerTieMove := board.Move(8, O)

	Assert(t, minimax(playerWinMove, 0, false) < minimax(playerTieMove, 0, false))
}

func TestMinimaxPrefersTie_ToPlayerWin(t *testing.T) {
	board := Board{X, O, Blank, X, X, Blank, O, X, O}
	blockMove := board.Move(5, O)
	loseMove := board.Move(2, O)

	Assert(t, minimax(blockMove, 0, true) > minimax(loseMove, 0, true))
}

func TestMinimaxPrefersFastWin_ToSlow(t *testing.T) {
	board := Board{X, O, O, Blank, Blank, O, Blank, X, X}
	expected := board.Move(4, X)

	AssertSliceEquals(t, expected, NextBoard(&board))
}

func TestMakesMove_EmptyBoard(t *testing.T) {
	board := NextBoard(getBlankBoard())
	Assert(t, !(&board).isEmpty())
}

func TestMakesMove_PopulatedBoard(t *testing.T) {
	board := getBlankBoard().Move(2, X)
	board = NextBoard(&board)

	AssertEquals(t, 7, len(board.getOpenSpaces()))
}

func TestNextBoard_TakesCornerFirstMove(t *testing.T) {
	expected := getBlankBoard().Move(0, X)
	AssertSliceEquals(t, expected, NextBoard(getBlankBoard()))
}

func TestBlocksPlayerWin(t *testing.T) {
	board := Board{X, O, X, Blank, Blank, Blank, Blank, O, Blank}
	expected := board.Move(4, X)

	AssertSliceEquals(t, expected, NextBoard(&board))
}

func TestNextBoard_BlocksCornerStrategy(t *testing.T) {
	board := getBlankBoard().Move(0, X)
	expected := board.Move(4, O)

	AssertSliceEquals(t, expected, NextBoard(&board))
}

func TestAlwaysWinsX(t *testing.T) {
	Assert(t, !canLose(*getBlankBoard(), "X"))
}

func TestAlwaysWinsO(t *testing.T) {
	didLose := false
	board := getBlankBoard()
	for _, space := range board.getOpenSpaces() {
		newBoard := board.Move(space, board.CurPiece())
		didLose = didLose || canLose(newBoard, "O")
	}
	Assert(t, !didLose)
}
