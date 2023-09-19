package menu

import (
	. "fmt"
	"ttgo"
)

type Menu struct {
	label   string
	options map[string]Stringer
}

func (menu Menu) String() string {
	return menu.label
}

func (menu Menu) NextState(selection string) (Stringer, error) {
	next := menu.options[selection]
	if next == nil {
		return menu, nil
	}
	return next, nil
}

var mainMenu = Menu{
	"Unbeatable Tic-Tac-Toe\nPlay as:\n1) X\n2) O\n",
	map[string]Stringer{
		"1": ttgo.Board{},
		"2": ttgo.NextBoard(new(ttgo.Board)),
	}}
