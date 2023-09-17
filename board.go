package ttgo

type Board [9]Piece

const boardSize int = 3

func (board *Board) pieceCount(p Piece) int {
	count := 0
	for i := range board {
		if board[i] == p {
			count++
		}
	}
	return count
}

func (board *Board) curPiece() Piece {
	xCount, oCount := board.pieceCount(x), board.pieceCount(o)
	if xCount == oCount {
		return x
	}
	return o
}

func (board *Board) getRows() []Line {
	return []Line{
		board[:3], board[3:6], board[6:9],
	}
}

func (board *Board) getCols() []Line {
	cols := [3]Line{}
	for i, piece := range board {
		cols[i%boardSize] = append(cols[i%boardSize], piece)
	}
	return cols[:]
}

func (board *Board) getDiags() []Line {
	diags := [2]Line{}
	for i := 0; i < len(board); i += boardSize + 1 {
		diags[0] = append(diags[0], board[i])
	}
	for i := 2; i < len(board)-1; i += boardSize - 1 {
		diags[1] = append(diags[1], board[i])
	}
	return diags[:]
}

func (board *Board) isWon() bool {
	lines := append(board.getRows(),
		append(board.getCols(), board.getDiags()...)...)
	for _, line := range lines {
		if line.isUniform() && !line.isEmpty() {
			return true
		}
	}
	return false
}

func (board *Board) isTied() bool {
	isFull := board.pieceCount(x)+board.pieceCount(o) == len(board)
	return !board.isWon() && isFull
}

func (board *Board) isGameOver() bool {
	return board.isWon() || board.isTied()
}

func (board *Board) getOpenSpaces() []int {
	spaces := make([]int, 0, 9)
	for i, piece := range board {
		if piece == blank {
			spaces = append(spaces, i)
		}
	}
	return spaces
}
