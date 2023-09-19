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

func TestRendersGameOverOptions(t *testing.T) {
	expected := "Play again?\n1) Yes\n2)Quit\n"
	AssertEquals(t, expected, gameOverMenu.String())
}

func TestGameOverMenu_PlaysAgain(t *testing.T) {
	next, _ := gameOverMenu.NextState("1")
	AssertDeepEquals(t, mainMenu, next)
}

func TestGameOverMenu_Exits(t *testing.T) {
	next, _ := gameOverMenu.NextState("2")
	AssertDeepEquals(t, Exit{}, next)
}

func TestExit_SaysGoodbye(t *testing.T) {
	AssertEquals(t, "Goodbye!", Exit{}.String())
}
