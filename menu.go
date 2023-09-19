package ttgo

import (
	. "fmt"
)

type Menu struct {
	label   string
	options map[string]Stringer
}

func (menu Menu) String() string {
	return menu.label
}

var mainMenu = Menu{
	"Unbeatable Tic-Tac-Toe\nPlay as:\n1) X\n2) O\n",
	map[string]Stringer{
		"1": Board{},
		"2": NextBoard(new(Board)),
	}}

var replayMenu = Menu{}
