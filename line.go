package ttgo

type Line []Piece

var pieceFmt = map[Piece]string{Blank: "-", X: "X", O: "O"}

func (pieces *Line) String() string {
	var lineStr string
	for i, piece := range *pieces {
		if i != 0 {
			lineStr = lineStr + " "
		}
		lineStr = lineStr + pieceFmt[piece]
	}
	return lineStr
}

func (pieces Line) isUniform() bool {
	return (pieces[0] == pieces[1]) && (pieces[1] == pieces[2])
}

func (pieces Line) isEmpty() bool {
	for i := range pieces {
		if pieces[i] != Blank {
			return false
		}
	}
	return true
}
