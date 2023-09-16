package ttgo

import "testing"

func (b board) isEmpty() bool {
	res := true
	for _, c := range b {
		if c != 0 {
			res = false
		}
	}
	return res
}

func TestInitializesBoardOfSizeNine(t *testing.T) {
	var b board
	assertEquals(t, 9, len(b))
}

func TestInitializesBoardBlank(t *testing.T) {
	var b board
	assert(t, b.isEmpty())
}

func TestTokensAssigned(t *testing.T) {
	assertEquals(t, piece(0), blank)
	assertEquals(t, piece(1), x)
	assertEquals(t, piece(2), o)
}

func TestWinnerReturnsX(t *testing.T) {
	assertEquals(t, x, getWinBoard().getWinner())
}

func TestWinnerReturnsO(t *testing.T) {
	oWinBoard := board{o, o, o, x, x, blank, x, blank, blank}
	assertEquals(t, o, oWinBoard.getWinner())
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
