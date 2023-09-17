package ttgo

import "testing"

func (board Board) isEmpty() bool {
	for _, c := range board {
		if c != 0 {
			return false
		}
	}
	return true
}

func TestInitializesBoardOfSizeNine(t *testing.T) {
	var b Board
	assertEquals(t, 9, len(b))
}

func TestInitializesBoardBlank(t *testing.T) {
	var b Board
	assert(t, b.isEmpty())
}

func TestTokensAssigned(t *testing.T) {
	assertEquals(t, Piece(0), blank)
	assertEquals(t, Piece(1), x)
	assertEquals(t, Piece(2), o)
}

func TestWinnerReturnsX(t *testing.T) {
	assertEquals(t, "X", getWinBoard().getWinner())
}

func TestWinnerReturnsO(t *testing.T) {
	oWinBoard := Board{o, o, o, x, x, blank, x, blank, blank}
	assertEquals(t, "O", oWinBoard.getWinner())
}

func TestWinnerReturnsTie(t *testing.T) {
	assertEquals(t, "Tie", getTieBoard().getWinner())
}

func TestWinnerReturnsPanicsForUnfinishedGame(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert(t, true)
		}
	}()
	getBlankBoard().getWinner()
	assert(t, false)
}
