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
