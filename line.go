package ttgo

type line []piece

func (pieces line) isUniform() bool {
	return (pieces[0] == pieces[1]) && (pieces[1] == pieces[2])
}

func (pieces line) isEmpty() bool {
	for i := range pieces {
		if pieces[i] != blank {
			return false
		}
	}
	return true
}
