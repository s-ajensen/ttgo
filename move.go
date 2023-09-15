package ttgo

func (b board) move(i int, p piece) board {
	b[i] = p
	return b
}
