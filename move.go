package ttgo

func (b board) move(i int, p piece) board {
	b[i] = p
	return b
}

func (b *board) isWon() bool {
	for _, row := range b.getRows() {
		if row.isUniform() && !row.isEmpty() {
			return true
		}
	}
	return false
}
