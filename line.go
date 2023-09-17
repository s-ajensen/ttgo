package ttgo

type Line []Piece

func (pieces Line) isUniform() bool {
	return (pieces[0] == pieces[1]) && (pieces[1] == pieces[2])
}

func (pieces Line) isEmpty() bool {
	for i := range pieces {
		if pieces[i] != blank {
			return false
		}
	}
	return true
}
