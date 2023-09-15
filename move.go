package ttgo

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
