package ttgo

type Piece byte

const (
	Blank = Piece(0)
	X     = Piece(1)
	O     = Piece(2)
)

func (board *Board) getWinner() string {
	if !board.isGameOver() {
		panic("Game is not over and has no winner!")
	}
	if board.isTied() {
		return "Tie"
	}
	xCount, oCount := board.pieceCount(X), board.pieceCount(O)
	if xCount == oCount {
		return "O"
	}
	return "X"
}
