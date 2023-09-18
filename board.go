package ttgo

type Board [9]Piece

const boardSize int = 3

func (board *Board) String() string {
	var boardStr string
	for i := 0; i < len(board); i += boardSize {
		pieces := make(Line, boardSize)
		pieces = board[i:(i + boardSize)]
		boardStr = boardStr + pieces.String() + "\n"
	}
	return boardStr
}

func (board *Board) pieceCount(p Piece) int {
	count := 0
	for i := range board {
		if board[i] == p {
			count++
		}
	}
	return count
}

func (board *Board) CurPiece() Piece {
	xCount, oCount := board.pieceCount(X), board.pieceCount(O)
	if xCount == oCount {
		return X
	}
	return O
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
	isFull := board.pieceCount(X)+board.pieceCount(O) == len(board)
	return !board.isWon() && isFull
}

func (board *Board) isGameOver() bool {
	return board.isWon() || board.isTied()
}

func (board *Board) getOpenSpaces() []int {
	spaces := make([]int, 0, 9)
	for i, piece := range board {
		if piece == Blank {
			spaces = append(spaces, i)
		}
	}
	return spaces
}
