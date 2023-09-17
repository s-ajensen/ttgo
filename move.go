package ttgo

import (
	"math"
)

func (b board) move(i int, p piece) board {
	b[i] = p
	return b
}

func (b *board) evalStatic(isMaximizing bool, depth int) int {
	scalar := 1
	if !isMaximizing {
		scalar = -1
	}
	if b.isWon() {
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

func (b *board) eval(depth int, isMaximizing bool) int {
	eval := getBound(isMaximizing)
	compFn := getCompFn(isMaximizing)
	for _, space := range b.getOpenSpaces() {
		newBoard := b.move(space, b.curPiece())
		eval = compFn(eval, minimax(newBoard, depth+1, isMaximizing))
	}
	return eval
}

func minimax(b board, depth int, isMaximizing bool) int {
	if b.isGameOver() {
		return b.evalStatic(isMaximizing, depth)
	}
	if isMaximizing {
		return b.eval(depth, false)
	}
	return b.eval(depth, true)
}

func nextBoard(b *board) board {
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
