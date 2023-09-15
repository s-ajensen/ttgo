package ttgo

type board [9]piece
type line []piece

const boardSize int = 3

func (b *board) curPiece() piece {
	xCount := 0
	oCount := 0
	for _, p := range b {
		if p == x {
			xCount++
		}
		if p == o {
			oCount++
		}
	}
	if xCount == oCount {
		return x
	}
	return o
}

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

func (b *board) getDiags() []line {
	diags := [2]line{}
	for i := 0; i < len(b); i += boardSize + 1 {
		diags[0] = append(diags[0], b[i])
	}
	for i := 2; i < len(b)-1; i += boardSize - 1 {
		diags[1] = append(diags[1], b[i])
	}
	return diags[:]
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
