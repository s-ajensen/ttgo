package menu

import (
	"fmt"
	"io"
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
	label   string
	options map[string]GameState
}

var menus = map[Menu]MenuNode{
	mainMenu: {
		label: "Unbeatable Tic-Tac-Toe\nPlay as:\n1) X\n2) O\n",
		options: map[string]GameState{
			"1": {menu: game},
			"2": {menu: game, board: NextBoard(new(Board))},
		}},
}

func render(writer io.Writer, state GameState) {
	if state.menu == game {
		fmt.Fprintf(writer, state.board.String())
	} else {
		label := menus[state.menu].label
		fmt.Fprintf(writer, label)
	}
}

func nextState(writer io.Writer, state GameState, input string) GameState {
	selection := menus[state.menu].options[input]
	if selection.menu == none {
		fmt.Fprintf(writer, "Bad selection, try again:\n")
		return state
	}
	return selection
}
