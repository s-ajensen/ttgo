package ttgo

import (
	"math"
)

func (b board) move(i int, p piece) board {
	b[i] = p
	return b
}

func (b *board) staticEval(isMaximizing bool) int {
	scalar := 1
	if !isMaximizing {
		scalar = -1
	}
	if b.isWon() {
		return 10 * scalar
	}
	return 0
}

func (b *board) eval(bound int, depth int, compFn func(n int, m int) int, isMaximizing bool) int {
	eval := bound
	for _, space := range b.getOpenSpaces() {
		newBoard := b.move(space, b.curPiece())
		eval = compFn(eval, minimax(newBoard, depth+1, isMaximizing)) - depth
	}
	return eval
}

func minimax(b board, depth int, isMaximizing bool) int {
	if b.isGameOver() {
		return b.staticEval(isMaximizing) - depth
	}
	if isMaximizing {
		return b.eval(math.MinInt, depth, maxInt, false)
	}
	return b.eval(math.MaxInt, depth, minInt, true)
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
