package ttgo

import (
	"strconv"
)

type Piece byte
type GameState interface {
	String() string
	NextState(string) (GameState, error)
}

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

func (board Board) NextState(selection string) (GameState, error) {
	if board.isGameOver() {
		return newGameOverMenu(&board), nil
	}
	newBoard, err := board.validateSelection(selection)
	if err != nil {
		return board, err
	}
	if newBoard.isGameOver() {
		return newGameOverMenu(&newBoard), nil
	}
	newBoard = NextBoard(&newBoard)
	if newBoard.isGameOver() {
		return newGameOverMenu(&newBoard), nil
	}
	return newBoard, nil
}

func (menu Menu) NextState(selection string) (GameState, error) {
	next := menu.options[selection]
	if next == nil {
		return menu, newInvalidOptionErr(selection)
	}
	return next, nil
}
