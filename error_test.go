package ttgo

import (
	"testing"
	. "ttgo/assert"
)

func TestInputErrorFormat(t *testing.T) {
	AssertEquals(t, "Invalid move!\nEnter an integer between 0 and 8.\n", newInputErr().Error())
}

func TestSpaceTakenErrorFormat(t *testing.T) {
	AssertEquals(t, "space at index '0' already taken\n", newSpaceTakenErr(0).Error())
}

func TestInvalidOptionErrorFormat(t *testing.T) {
	AssertEquals(t, "invalid menu option '0'\nTry again:\n", newInvalidOptionErr("0").Error())
}
