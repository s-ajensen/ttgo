package ttgo

import (
	"errors"
	"fmt"
)

func newInputErr() error {
	return errors.New(fmt.Sprintf("Invalid move!\nEnter an integer between 0 and %d.\n", len(Board{})-1))
}

func newSpaceTakenErr(move int) error {
	return errors.New(fmt.Sprintf("space at index '%d' already taken\n", move))
}

func newInvalidOptionErr(selection string) error {
	return errors.New(fmt.Sprintf("invalid menu option '%s'\nTry again:\n", selection))
}
