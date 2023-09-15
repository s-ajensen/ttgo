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
	winBoard := board{x, x, x, blank, o, o, o, blank, blank}

	assert(t, !unfinishedBoard.isWon())
	assert(t, winBoard.isWon())
}

func TestDeterminesWinByColumn(t *testing.T) {
	winBoard := board{x, o, o, x, blank, blank, x, blank, blank}
	assert(t, winBoard.isWon())
}

func TestDeterminesWinByDiags(t *testing.T) {
	winBoardX := board{x, o, o, blank, x, blank, blank, blank, x}
	winBoardO := board{x, x, o, x, o, blank, o, blank, blank}

	assert(t, winBoardX.isWon())
	assert(t, winBoardO.isWon())
}

func TestCalcsNextPiecePlayed_BlankBoard(t *testing.T) {
	var b board
	assertEquals(t, b.curPiece(), x)
}

func TestCalcsNextPiecePlayed_SingleMovePlayed(t *testing.T) {
	b := board{x, blank, blank, blank, blank, blank, blank, blank, blank}
	assertEquals(t, b.curPiece(), o)
}

func TestCalcsNextPiecePlayed_ManyMovesPlayed(t *testing.T) {
	xBoard := board{x, o, blank, blank, blank, blank, blank, blank, blank}
	oBoard := board{x, blank, x, blank, blank, blank, blank, blank, blank}

	assertEquals(t, xBoard.curPiece(), x)
	assertEquals(t, oBoard.curPiece(), o)
}
