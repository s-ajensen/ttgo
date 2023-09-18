package menu

import (
	"fmt"
	"io"
	"strconv"
	. "ttgo"
)

type Menu int

const (
	none     = Menu(0)
	mainMenu = Menu(1)
	game     = Menu(2)
)

type GameState struct {
	menu  Menu
	board Board
}

type MenuNode struct {
	renderFn func(GameState) string
	nextFn   func(GameState, string) GameState
}

func renderMenu(label string) func(state GameState) string {
	return func(_ GameState) string {
		return label
	}
}

func nextMenu(options map[string]GameState) func(GameState, string) GameState {
	return func(_ GameState, input string) GameState {
		return options[input]
	}
}

func playFn() func(GameState, string) GameState {
	return func(state GameState, s string) GameState {
		selection, _ := strconv.Atoi(s)
		board := state.board
		newBoard := board.Move(selection, board.CurPiece())
		return GameState{game, newBoard}
	}
}

var menus = map[Menu]MenuNode{
	mainMenu: {
		renderFn: renderMenu("Unbeatable Tic-Tac-Toe\nPlay as:\n1) X\n2) O\n"),
		nextFn: nextMenu(map[string]GameState{
			"1": {menu: game},
			"2": {menu: game, board: NextBoard(new(Board))},
		})},
	game: {
		renderFn: func(state GameState) string {
			return state.board.String()
		},
		nextFn: playFn()},
}

func render(writer io.Writer, state GameState) {
	label := menus[state.menu].renderFn(state)
	fmt.Fprintf(writer, label)
}

func nextState(writer io.Writer, state GameState, input string) GameState {
	selection := menus[state.menu].nextFn(state, input)
	if selection.menu == none {
		fmt.Fprintf(writer, "Bad selection, try again:\n")
		return state
	}
	return selection
}
