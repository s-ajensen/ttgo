package ttgo

type board [9]piece
type line []piece

const boardSize int = 3

func (b *board) getRows() []line {
	return []line{
		b[:3], b[3:6], b[6:9],
	}
}

func (b *board) getCols() []line {
	cols := [3]line{}
	for i, p := range b {
		cols[i%boardSize] = append(cols[i%boardSize], p)
	}
	return cols[:]
}

func (pieces line) isUniform() bool {
	return (pieces[0] == pieces[1]) && (pieces[1] == pieces[2])
}

func (pieces line) isEmpty() bool {
	for _, p := range pieces {
		if p != b {
			return false
		}
	}
	return true
}
