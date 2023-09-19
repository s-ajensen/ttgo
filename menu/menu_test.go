package menu

import (
	"errors"
	"testing"
	"ttgo"
	. "ttgo/assert"
)

func TestRendersPlayerOptions(t *testing.T) {
	expected := "Unbeatable Tic-Tac-Toe\nPlay as:\n1) X\n2) O\n"

	AssertEquals(t, expected, mainMenu.String())
}

func TestSelectsPlayAsX(t *testing.T) {
	expected := ttgo.Board{}

	next, err := mainMenu.NextState("1")
	AssertEquals(t, expected, next)
	AssertEquals(t, nil, err)
}

func TestSelectsPlayAsO(t *testing.T) {
	var board ttgo.Board
	expected := ttgo.NextBoard(&board)

	next, err := mainMenu.NextState("2")
	AssertEquals(t, expected, next)
	AssertEquals(t, nil, err)
}

func TestReturnsErrorFor_UnknownOption(t *testing.T) {
	next, err := mainMenu.NextState("3")
	expectedErr := errors.New("Invalid option '3'\nTry again:\n")

	AssertDeepEquals(t, mainMenu, next)
	AssertEquals(t, expectedErr.Error(), err.Error())
}
