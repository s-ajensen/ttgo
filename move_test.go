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

func TestDeterminesWinByDiags(t *testing.T) {
	winBoardX := board{x, o, o, b, x, b, b, b, x}
	winBoardO := board{x, x, o, x, o, b, o, b, b}

	assert(t, winBoardX.isWon())
	assert(t, winBoardO.isWon())
}

func TestCalcsNextPiecePlayed_BlankBoard(t *testing.T) {
	var b board
	assertEquals(t, b.curPiece(), x)
}

func TestCalcsNextPiecePlayed_SingleMovePlayed(t *testing.T) {
	b := board{x, b, b, b, b, b, b, b, b}
	assertEquals(t, b.curPiece(), o)
}

func TestCalcsNextPiecePlayed_ManyMovesPlayed(t *testing.T) {
	xBoard := board{x, o, b, b, b, b, b, b, b}
	oBoard := board{x, b, x, b, b, b, b, b, b}

	assertEquals(t, xBoard.curPiece(), x)
	assertEquals(t, oBoard.curPiece(), o)
}
