package ttgo

import "math"

func (b board) move(i int, p piece) board {
	b[i] = p
	return b
}

func (b *board) eval(isMaximizing bool) int {
	scalar := 1
	if !isMaximizing {
		scalar = -1
	}
	if b.isWon() {
		return 10 * scalar
	}
	return 0
}

func minimax(b board, isMaximizing bool) int {
	if b.isWon() || b.isTied() {
		return b.eval(isMaximizing)
	}
	if isMaximizing {
		maxEval := math.MinInt
		for _, space := range b.getOpenSpaces() {
			newBoard := b.move(space, b.curPiece())
			eval := minimax(newBoard, false)
			maxEval = max(maxEval, eval)
		}
		return maxEval
	} else {
		minEval := math.MaxInt
		for _, space := range b.getOpenSpaces() {
			newBoard := b.move(space, b.curPiece())
			eval := minimax(newBoard, true)
			minEval = min(minEval, eval)
		}
		return minEval
	}
}
