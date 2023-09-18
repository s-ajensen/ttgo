package menu

import (
	"fmt"
	"io"
	. "ttgo"
)

type Menu int

const (
	mainMenu = Menu(0)
	game     = Menu(1)
)

type GameState struct {
	menu  Menu
	board Board
}

func render(writer io.Writer, state GameState) {
	if state.menu == mainMenu {
		fmt.Fprintf(writer, "Unbeatable Tic-Tac-Toe\nPlay as:\n1) X 2) O\n")
	} else {
		fmt.Fprintf(writer, state.board.String())
	}
}

func nextState(writer io.Writer, state GameState, input string) GameState {
	board := Board{}
	if input == "1" {
		return GameState{menu: game, board: board}
	}
	if input == "2" {
		return GameState{menu: game, board: NextBoard(&board)}
	}
	fmt.Fprintf(writer, "Bad selection, try again:\n")
	return state
}
