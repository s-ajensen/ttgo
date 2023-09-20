package main

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
	"ttgo"
	. "ttgo/assert"
)

var out = bytes.Buffer{}
var in = bytes.Buffer{}

func clearIO() {
	out.Reset()
	in.Reset()
}

type mockState struct {
	idx int
}

func (state mockState) String() string {
	return fmt.Sprintf("mock state: %d", state.idx)
}

func (state mockState) NextState(selection string) (ttgo.GameState, error) {
	idx, _ := strconv.Atoi(selection)
	return mockStates[idx+1], nil
}

type errorState struct {
	didError bool
}

func (state errorState) String() string {
	return "error state "
}

func (state errorState) NextState(selection string) (ttgo.GameState, error) {
	if state.didError {
		return mockStates[0], nil
	}
	return errorState{true}, errors.New("error state msg")
}

var mockInitialState = mockState{0}
var mockMiddleState = mockState{1}
var mockStates = []ttgo.GameState{mockInitialState, mockMiddleState, ttgo.Exit{}}

func TestPrintsInitialState(t *testing.T) {
	clearIO()
	in.WriteString("0\n1\n")
	gameLoop(mockInitialState, &in, &out)
	result := out.String()

	Assert(t, strings.Contains(result, mockInitialState.String()))
}

func TestPrintsNextState(t *testing.T) {
	clearIO()
	in.WriteString("0\n1\n")
	gameLoop(mockInitialState, &in, &out)
	result := out.String()

	Assert(t, strings.Contains(result, mockMiddleState.String()))
}

func TestStopsForExitState(t *testing.T) {
	clearIO()
	in.WriteString("0\n1\n")
	gameLoop(mockInitialState, &in, &out)
	result := out.String()

	Assert(t, strings.Contains(result, "Goodbye!\n"))
}

func TestRedrawsState_WhenErrorThrown(t *testing.T) {
	clearIO()
	in.WriteString("0\n0\n1\n")
	gameLoop(errorState{false}, &in, &out)
	result := out.String()

	Assert(t, strings.Contains(result, "error state error state"))
}

func TestPrintsErrorMessage_WhenErrorThrown(t *testing.T) {
	clearIO()
	in.WriteString("0\n0\n1\n")
	gameLoop(errorState{false}, &in, &out)
	result := out.String()

	Assert(t, strings.Contains(result, "error state msg"))
}
