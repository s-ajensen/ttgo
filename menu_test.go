package ttgo

import (
	"strings"
	"testing"
	. "ttgo/assert"
)

func TestRendersPlayerOptions(t *testing.T) {
	expected := "Unbeatable Tic-Tac-Toe\nPlay as:\n1) X\n2) O\n"

	AssertEquals(t, expected, MainMenu.String())
}

func TestSelectsPlayAsX(t *testing.T) {
	expected := Board{}

	next, err := MainMenu.NextState("1")
	AssertEquals(t, expected, next)
	AssertEquals(t, nil, err)
}

func TestSelectsPlayAsO(t *testing.T) {
	var board Board
	expected := NextBoard(&board)

	next, err := MainMenu.NextState("2")
	AssertEquals(t, expected, next)
	AssertEquals(t, nil, err)
}

func TestWinLabel_FormatsTie(t *testing.T) {
	AssertEquals(t, "Tie!\n", getWinLabel(getTieBoard()))
}

func TestWinLabel_FormatsWin(t *testing.T) {
	AssertEquals(t, "X wins!\n", getWinLabel(getWinBoard()))
}

func TestGameOverMenu_RendersFinalBoard(t *testing.T) {
	expected := getWinBoard().String()
	Assert(t, strings.Contains(newGameOverMenu(getWinBoard()).String(), expected))
}

func TestGameOverManu_RendersWinLabel(t *testing.T) {
	expected := getWinLabel(getTieBoard())
	Assert(t, strings.Contains(newGameOverMenu(getTieBoard()).String(), expected))
}

func TestGameOverMenu_RendersOptions(t *testing.T) {
	expected := "Play again?\n1) Yes\n2) Quit\n"
	Assert(t, strings.Contains(newGameOverMenu(getWinBoard()).String(), expected))
}

func TestGameOverMenu_PlaysAgain(t *testing.T) {
	next, _ := newGameOverMenu(getWinBoard()).NextState("1")
	AssertDeepEquals(t, MainMenu, next)
}

func TestGameOverMenu_Exits(t *testing.T) {
	next, _ := newGameOverMenu(getWinBoard()).NextState("2")
	AssertDeepEquals(t, Exit{}, next)
}

func TestExit_SaysExit(t *testing.T) {
	AssertEquals(t, ExitFlag, Exit{}.String())
}
