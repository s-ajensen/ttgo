package ttgo

type piece byte

const (
	blank = piece(0)
	x     = piece(1)
	o     = piece(2)
)

func (b *board) getWinner() string {
	if !b.isGameOver() {
		panic("Game is not over and has no winner!")
	}
	if b.isTied() {
		return "Tie"
	}
	xCount, oCount := b.pieceCount(x), b.pieceCount(o)
	if xCount == oCount {
		return "O"
	}
	return "X"
}
