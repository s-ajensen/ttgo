package ttgo

type piece byte

const (
	blank = piece(0)
	x     = piece(1)
	o     = piece(2)
)

type board [9]piece
