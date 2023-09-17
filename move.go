package ttgo

import (
	"math"
)

func (board Board) move(i int, p Piece) Board {
	board[i] = p
	return board
}

func (board *Board) evalStatic(isMaximizing bool, depth int) int {
	scalar := 1
	if !isMaximizing {
		scalar = -1
	}
	if board.isWon() {
		return (10 - depth) * scalar
	}
	return 0
}

func getCompFn(isMaximizing bool) func(n int, m int) int {
	if isMaximizing {
		return maxInt
	}
	return minInt
}

func getBound(isMaximizing bool) int {
	if isMaximizing {
		return math.MinInt
	}
	return math.MaxInt
}

func (board *Board) eval(depth int, isMaximizing bool) int {
	eval := getBound(isMaximizing)
	compFn := getCompFn(isMaximizing)
	for _, space := range board.getOpenSpaces() {
		newBoard := board.move(space, board.curPiece())
		eval = compFn(eval, minimax(newBoard, depth+1, isMaximizing))
	}
	return eval
}

func minimax(b Board, depth int, isMaximizing bool) int {
	if b.isGameOver() {
		return b.evalStatic(isMaximizing, depth)
	}
	if isMaximizing {
		return b.eval(depth, false)
	}
	return b.eval(depth, true)
}

func NextBoard(b *Board) Board {
	bestBoard, moveWeight := *b, math.MinInt
	for _, space := range b.getOpenSpaces() {
		possibleBoard := b.move(space, b.curPiece())
		possibleWeight := minimax(possibleBoard, 0, true)
		if possibleWeight > moveWeight {
			bestBoard, moveWeight = possibleBoard, possibleWeight
		}
	}
	return bestBoard
}
