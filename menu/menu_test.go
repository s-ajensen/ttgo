package menu

import (
	"bytes"
	"testing"
	"ttgo"
	. "ttgo/assert"
)

var output bytes.Buffer

func TestWritesPlayerOptions(t *testing.T) {
	output.Reset()
	render(&output, GameState{menu: mainMenu})
	expected := "Unbeatable Tic-Tac-Toe\nPlay as:\n1) X 2) O\n"

	AssertEquals(t, expected, output.String())
}

func TestSelectsPlayAsX(t *testing.T) {
	output.Reset()
	expected := GameState{menu: game, board: ttgo.Board{}}
	AssertEquals(t, expected, nextState(&output, GameState{}, "1"))
}

func TestSelectsPlayAsO(t *testing.T) {
	output.Reset()
	board := ttgo.Board{}
	expected := GameState{menu: game, board: board.Move(0, ttgo.X)}

	AssertEquals(t, expected, nextState(&output, GameState{}, "2"))
}

func TestRetriesSelectPlayAs_OnBadInput(t *testing.T) {
	output.Reset()
	expected := GameState{menu: mainMenu}

	AssertEquals(t, expected, nextState(&output, GameState{}, "3"))
	AssertEquals(t, "Bad selection, try again:\n", output.String())
}

func TestWritesEmptyBoard(t *testing.T) {
	output.Reset()
	board := ttgo.Board{}
	render(&output, GameState{menu: game, board: board})
	expected := board.String()

	AssertEquals(t, expected, output.String())
}
