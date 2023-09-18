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

var menus = map[Menu]map[string]GameState{
	mainMenu: map[string]GameState{
		"1": GameState{menu: game},
		"2": GameState{menu: game, board: NextBoard(new(Board))},
	},
}

func render(writer io.Writer, state GameState) {
	if state.menu == mainMenu {
		fmt.Fprintf(writer, "Unbeatable Tic-Tac-Toe\nPlay as:\n1) X 2) O\n")
	} else {
		fmt.Fprintf(writer, state.board.String())
	}
}

func nextState(writer io.Writer, state GameState, input string) GameState {
	selection := menus[state.menu][input]
	if selection.menu == none {
		fmt.Fprintf(writer, "Bad selection, try again:\n")
		return state
	}
	return selection
}
