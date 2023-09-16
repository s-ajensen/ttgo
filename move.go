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

func (b *board) eval(bound int, compFn func(n int, m int) int, isMaximizing bool) int {
	eval := bound
	for _, space := range b.getOpenSpaces() {
		newBoard := b.move(space, b.curPiece())
		eval = compFn(eval, minimax(newBoard, isMaximizing))
	}
	return eval
}

func minimax(b board, isMaximizing bool) int {
	if b.isWon() || b.isTied() {
		return b.staticEval(isMaximizing)
	}
	if isMaximizing {
		return b.eval(math.MinInt, maxInt, false)
	}
	return b.eval(math.MaxInt, minInt, true)
}
