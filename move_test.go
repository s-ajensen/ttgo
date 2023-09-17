package ttgo

import (
	"testing"
)

func canLose(board Board, playingAs string) bool {
	resp := NextBoard(&board)
	if resp.isGameOver() {
		return playingAs != resp.getWinner() && "Tie" != resp.getWinner()
	}
	isLoss := false
	for _, space := range resp.getOpenSpaces() {
		newBoard := resp.move(space, resp.curPiece())
		if newBoard.isGameOver() {
			return playingAs != newBoard.getWinner() && "Tie" != newBoard.getWinner()
		}
		isLoss = isLoss || canLose(newBoard, playingAs)
	}
	return isLoss
}

func TestMovePlayedOnBlankBoard(t *testing.T) {
	var blankBoard Board
	var playedBoard = blankBoard.move(0, x)

	assertEquals(t, x, playedBoard[0])
	assertNotEquals(t, x, blankBoard[0])
}

func TestEvaluatesTieAsZero_WhenMaximizing(t *testing.T) {
	assertEquals(t, 0, getTieBoard().evalStatic(true, 0))
}

func TestEvaluatesTieAsZero_WhenMinimizing(t *testing.T) {
	assertEquals(t, 0, getTieBoard().evalStatic(false, 0))
}

func TestEvaluatesWinAsTen_WhenMaximizing(t *testing.T) {
	assertEquals(t, 10, getWinBoard().evalStatic(true, 0))
}

func TestEvaluatesWinAsMinusTen_WhenMinimizing(t *testing.T) {
	assertEquals(t, -10, getWinBoard().evalStatic(false, 0))
}

func TestMinimaxReturnsStaticEval_WhenGameOver(t *testing.T) {
	assertEquals(t, 0, minimax(*getTieBoard(), 0, true))
	assertEquals(t, 10, minimax(*getWinBoard(), 0, true))
	assertEquals(t, -10, minimax(*getWinBoard(), 0, false))
}

func TestMinimaxPrefersWin_ToTie(t *testing.T) {
	board := Board{x, o, x, x, o, x, o, blank, blank}
	winMove := board.move(7, o)
	tieMove := board.move(8, o)

	assert(t, minimax(winMove, 0, true) > minimax(tieMove, 0, true))
}

func TestMinimaxIdentifies_PlayerWin(t *testing.T) {
	board := Board{x, o, x, x, o, x, o, blank, blank}
	playerWinMove := board.move(7, o)
	playerTieMove := board.move(8, o)

	assert(t, minimax(playerWinMove, 0, false) < minimax(playerTieMove, 0, false))
}

func TestMinimaxPrefersTie_ToPlayerWin(t *testing.T) {
	board := Board{x, o, blank, x, x, blank, o, x, o}
	blockMove := board.move(5, o)
	loseMove := board.move(2, o)

	assert(t, minimax(blockMove, 0, true) > minimax(loseMove, 0, true))
}

func TestMinimaxPrefersFastWin_ToSlow(t *testing.T) {
	board := Board{x, o, o, blank, blank, o, blank, x, x}
	expected := board.move(4, x)

	assertSliceEquals(t, expected, NextBoard(&board))
}

func TestMakesMove_EmptyBoard(t *testing.T) {
	assert(t, !NextBoard(getBlankBoard()).isEmpty())
}

func TestMakesMove_PopulatedBoard(t *testing.T) {
	board := getBlankBoard().move(2, x)
	board = NextBoard(&board)

	assertEquals(t, 7, len(board.getOpenSpaces()))
}

func TestNextBoard_TakesCornerFirstMove(t *testing.T) {
	expected := getBlankBoard().move(0, x)
	assertSliceEquals(t, expected, NextBoard(getBlankBoard()))
}

func TestBlocksPlayerWin(t *testing.T) {
	board := Board{x, o, x, blank, blank, blank, blank, o, blank}
	expected := board.move(4, x)

	assertSliceEquals(t, expected, NextBoard(&board))
}

func TestNextBoard_BlocksCornerStrategy(t *testing.T) {
	board := getBlankBoard().move(0, x)
	expected := board.move(4, o)

	assertSliceEquals(t, expected, NextBoard(&board))
}

func TestAlwaysWinsX(t *testing.T) {
	assert(t, !canLose(*getBlankBoard(), "X"))
}

func TestAlwaysWinsO(t *testing.T) {
	didLose := false
	board := getBlankBoard()
	for _, space := range board.getOpenSpaces() {
		newBoard := board.move(space, board.curPiece())
		didLose = didLose || canLose(newBoard, "O")
	}
	assert(t, !didLose)
}
