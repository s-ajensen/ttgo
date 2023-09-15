package ttgo

import "testing"

func (b board) isEmpty() bool {
	res := true
	for _, c := range b {
		if c != 0 {
			res = false
		}
	}
	return res
}

func TestInitializesBoardOfSizeNine(t *testing.T) {
	var b board
	equals(t, 9, len(b))
}

func TestInitializesBoardBlank(t *testing.T) {
	var b board
	assert(t, b.isEmpty())
}
