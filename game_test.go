package ttgo

import (
	"testing"
	. "ttgo/assert"
)

func (board *Board) isEmpty() bool {
	for _, c := range board {
		if c != 0 {
			return false
		}
	}
	return true
}

func TestInitializesBoardOfSizeNine(t *testing.T) {
	var b Board
	AssertEquals(t, 9, len(b))
}

func TestInitializesBoardBlank(t *testing.T) {
	board := new(Board)
	Assert(t, board.isEmpty())
}

func TestTokensAssigned(t *testing.T) {
	AssertEquals(t, Piece(0), Blank)
	AssertEquals(t, Piece(1), X)
	AssertEquals(t, Piece(2), O)
}

func TestWinnerReturnsX(t *testing.T) {
	AssertEquals(t, "X", getWinBoard().getWinner())
}

func TestWinnerReturnsO(t *testing.T) {
	oWinBoard := Board{O, O, O, X, X, Blank, X, Blank, Blank}
	AssertEquals(t, "O", oWinBoard.getWinner())
}

func TestWinnerReturnsTie(t *testing.T) {
	AssertEquals(t, "Tie", getTieBoard().getWinner())
}

func TestWinnerReturnsPanicsForUnfinishedGame(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			Assert(t, true)
		}
	}()
	getBlankBoard().getWinner()
	Assert(t, false)
}
