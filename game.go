package ttgo

type piece byte

const (
	blank = piece(0)
	x     = piece(1)
	o     = piece(2)
)

func (b *board) getWinner() piece {
	if !b.isGameOver() {
		panic("Game is not over and has no winner!")
	}
	xCount, oCount := b.pieceCount(x), b.pieceCount(o)
	if xCount == oCount {
		return o
	}
	return x
}
