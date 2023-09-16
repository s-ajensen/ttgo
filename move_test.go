package ttgo

import (
	"testing"
)

func TestMovePlayedOnBlankBoard(t *testing.T) {
	var blankBoard board
	var playedBoard = blankBoard.move(0, x)

	assertEquals(t, x, playedBoard[0])
	assertNotEquals(t, x, blankBoard[0])
}

func TestEvaluatesTieAsZero_WhenMaximizing(t *testing.T) {
	assertEquals(t, 0, getTieBoard().eval(true))
}

func TestEvaluatesTieAsZero_WhenMinimizing(t *testing.T) {
	assertEquals(t, 0, getTieBoard().eval(false))
}

func TestEvaluatesWinAsTen_WhenMaximizing(t *testing.T) {
	assertEquals(t, 10, getWinBoard().eval(true))
}

func TestEvaluatesWinAsMinusTen_WhenMinimizing(t *testing.T) {
	assertEquals(t, -10, getWinBoard().eval(false))
}

func TestMinimaxReturnsStaticEval_WhenGameOver(t *testing.T) {
	assertEquals(t, 0, minimax(*getTieBoard(), true))
	assertEquals(t, 10, minimax(*getWinBoard(), true))
	assertEquals(t, -10, minimax(*getWinBoard(), false))
}

func TestMinimaxPrefersWin_ToTie(t *testing.T) {
	b := board{x, o, x, x, o, x, o, blank, blank}
	winMove := b.move(7, o)
	tieMove := b.move(8, o)

	assert(t, minimax(winMove, true) > minimax(tieMove, true))
}

func TestMinimaxIdentifies_PlayerWin(t *testing.T) {
	b := board{x, o, x, x, o, x, o, blank, blank}
	playerWinMove := b.move(7, o)
	playerTieMove := b.move(8, o)

	assert(t, minimax(playerWinMove, false) < minimax(playerTieMove, false))
}

func TestMinimaxPrefersTie_ToPlayerWin(t *testing.T) {
	b := board{x, o, blank, x, x, blank, o, x, o}
	blockMove := b.move(5, o)
	loseMove := b.move(2, o)

	println(minimax(blockMove, true))
	println(minimax(loseMove, true))

	assert(t, minimax(blockMove, true) > minimax(loseMove, true))
}
