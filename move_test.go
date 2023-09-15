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

func TestDeterminesWinByRow(t *testing.T) {
	var unfinishedBoard board
	winBoard := board{x, x, x, b, o, o, o, b, b}

	assert(t, !unfinishedBoard.isWon())
	assert(t, winBoard.isWon())
}

func TestDeterminesWinByColumn(t *testing.T) {
	winBoard := board{x, o, o, x, b, b, x, b, b}

	assert(t, winBoard.isWon())
}
