package ttgo

import (
	"testing"
	. "ttgo/assert"
)

func (board *Board) isEmpty() bool {
	for _, c := range board {
		if c != 0 {
			return false
		}
	}
	return true
}

func TestInitializesBoardOfSizeNine(t *testing.T) {
	var b Board
	AssertEquals(t, 9, len(b))
}

func TestInitializesBoardBlank(t *testing.T) {
	board := new(Board)
	Assert(t, board.isEmpty())
}

func TestTokensAssigned(t *testing.T) {
	AssertEquals(t, Piece(0), Blank)
	AssertEquals(t, Piece(1), X)
	AssertEquals(t, Piece(2), O)
}

func TestWinnerReturnsX(t *testing.T) {
	AssertEquals(t, "X", getWinBoard().getWinner())
}

func TestWinnerReturnsO(t *testing.T) {
	oWinBoard := Board{O, O, O, X, X, Blank, X, Blank, Blank}
	AssertEquals(t, "O", oWinBoard.getWinner())
}

func TestWinnerReturnsTie(t *testing.T) {
	AssertEquals(t, "Tie", getTieBoard().getWinner())
}

func TestWinnerPanicsForUnfinishedGame(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			Assert(t, true)
		}
	}()
	getBlankBoard().getWinner()
	Assert(t, false)
}

func TestNextStatePlaysMove(t *testing.T) {
	var board Board
	nextBoard, err := board.NextState("0")
	AssertEquals(t, 'X', rune(nextBoard.String()[0]))
	AssertEquals(t, nil, err)
}

func TestNextStateResponds(t *testing.T) {
	var board Board
	firstMove, _ := board.Move(0, X)
	expected := NextBoard(&firstMove).String()
	nextBoard, _ := board.NextState("0")

	AssertEquals(t, expected, nextBoard.String())
}

func TestNextStateReturnsErrorForEmptyInput(t *testing.T) {
	var board Board

	nextBoardEmpty, emptyErr := board.NextState("")

	AssertEquals(t, newInputErr().Error(), emptyErr.Error())
	AssertEquals(t, board, nextBoardEmpty)
}

func TestNextStateReturnsErrorFor_NonIntegerInput(t *testing.T) {
	var board Board
	next, err := board.NextState("not an index")

	AssertEquals(t, newInputErr().Error(), err.Error())
	AssertEquals(t, board, next)
}

func TestNextStateReturnsErrorFor_OutOfBoundsInput(t *testing.T) {
	var board Board
	nextBoardTooBig, tooBigErr := board.NextState("10")
	nextBoardTooSmall, tooSmallErr := board.NextState("-1")

	AssertEquals(t, newInputErr().Error(), tooBigErr.Error())
	AssertEquals(t, board, nextBoardTooBig)
	AssertEquals(t, newInputErr().Error(), tooSmallErr.Error())
	AssertEquals(t, board, nextBoardTooSmall)
}

func TestNextStateReturns_ReplayMenu_ForFinishedGame(t *testing.T) {
	board := Board{X, X, Blank, O, Blank, O, Blank, Blank, Blank}
	nextState, _ := board.NextState("2")
	expected, _ := board.Move(2, X)

	AssertDeepEquals(t, newGameOverMenu(&expected), nextState)
}

func TestReturnsErrorFor_UnknownOption(t *testing.T) {
	next, err := MainMenu.NextState("3")
	expectedErr := newInvalidOptionErr("3")

	AssertDeepEquals(t, MainMenu, next)
	AssertEquals(t, expectedErr.Error(), err.Error())
}
