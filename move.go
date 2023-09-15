package ttgo

func (b board) move(i int, p piece) board {
	b[i] = p
	return b
}

func (b *board) isWon() bool {
	lines := append(b.getRows(),
		append(b.getCols(), b.getDiags()...)...)
	for _, line := range lines {
		if line.isUniform() && !line.isEmpty() {
			return true
		}
	}
	return false
}
