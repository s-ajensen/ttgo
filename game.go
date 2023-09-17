package ttgo

type Piece byte

const (
	blank = Piece(0)
	x     = Piece(1)
	o     = Piece(2)
)

func (board *Board) getWinner() string {
	if !board.isGameOver() {
		panic("Game is not over and has no winner!")
	}
	if board.isTied() {
		return "Tie"
	}
	xCount, oCount := board.pieceCount(x), board.pieceCount(o)
	if xCount == oCount {
		return "O"
	}
	return "X"
}
