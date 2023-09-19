package ttgo

import (
	"testing"
	. "ttgo/assert"
)

func TestRendersPlayerOptions(t *testing.T) {
	expected := "Unbeatable Tic-Tac-Toe\nPlay as:\n1) X\n2) O\n"

	AssertEquals(t, expected, mainMenu.String())
}

func TestSelectsPlayAsX(t *testing.T) {
	expected := Board{}

	next, err := mainMenu.NextState("1")
	AssertEquals(t, expected, next)
	AssertEquals(t, nil, err)
}

func TestSelectsPlayAsO(t *testing.T) {
	var board Board
	expected := NextBoard(&board)

	next, err := mainMenu.NextState("2")
	AssertEquals(t, expected, next)
	AssertEquals(t, nil, err)
}
