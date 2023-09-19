package ttgo

import (
	. "fmt"
	"strconv"
)

type Piece byte

const (
	Blank = Piece(0)
	X     = Piece(1)
	O     = Piece(2)
)

func (board *Board) getWinner() string {
	if !board.isGameOver() {
		panic("Game is not over and has no winner!")
	}
	if board.isTied() {
		return "Tie"
	}
	xCount, oCount := board.pieceCount(X), board.pieceCount(O)
	if xCount == oCount {
		return "O"
	}
	return "X"
}

func (board *Board) validateSelection(selection string) (Board, error) {
	space, parseErr := strconv.Atoi(selection)
	if parseErr != nil {
		return *board, newInputErr()
	}
	newBoard, moveErr := board.Move(space, board.CurPiece())
	if moveErr != nil {
		return *board, moveErr
	}
	return newBoard, nil
}

func (board Board) NextState(selection string) (Stringer, error) {
	newBoard, err := board.validateSelection(selection)
	if err != nil {
		return board, err
	}
	if newBoard.isGameOver() {
		return replayMenu, nil
	}
	return newBoard, nil
}

func (menu Menu) NextState(selection string) (Stringer, error) {
	next := menu.options[selection]
	if next == nil {
		return menu, newInvalidOptionErr(selection)
	}
	return next, nil
}
