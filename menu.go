package ttgo

import "fmt"

type Menu struct {
	label   string
	options map[string]GameState
}

func (menu Menu) String() string {
	return menu.label
}

type Exit struct{}

const ExitFlag = "EXIT"

func (exit Exit) String() string {
	return ExitFlag
}

func (exit Exit) NextState(selection string) (GameState, error) {
	return nil, nil
}

var MainMenu = Menu{
	"Unbeatable Tic-Tac-Toe\nPlay as:\n1) X\n2) O\n",
	map[string]GameState{
		"1": Board{},
		"2": NextBoard(new(Board)),
	}}

func getWinLabel(board *Board) string {
	if board.isTied() {
		return "Tie!\n"
	}
	return fmt.Sprintf("%s wins!\n", board.getWinner())
}

func newGameOverMenu(board *Board) Menu {
	label := board.String() + getWinLabel(board) + "Play again?\n1) Yes\n2) Quit\n"
	return Menu{
		label,
		map[string]GameState{
			"1": MainMenu,
			"2": Exit{},
		},
	}
}
