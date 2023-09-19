package ttgo

import (
	"math"
)

func validateMove(move int) (int, error) {
	if move < 0 || move > boardSize {
		return move, inputErr
	}
	return move, nil
}

func (board Board) Move(i int, p Piece) (Board, error) {
	space, err := validateMove(i)
	if err != nil {
		return board, err
	}
	board[space] = p
	return board, nil
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
		newBoard, _ := board.Move(space, board.CurPiece())
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
		possibleBoard, _ := b.Move(space, b.CurPiece())
		possibleWeight := minimax(possibleBoard, 0, true)
		if possibleWeight > moveWeight {
			bestBoard, moveWeight = possibleBoard, possibleWeight
		}
	}
	return bestBoard
}
