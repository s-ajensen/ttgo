package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"ttgo"
)

func gameLoop(initState ttgo.GameState, in io.Reader, out io.Writer) {
	reader := bufio.NewReader(in)
	var state = initState
	for state.String() != ttgo.ExitFlag {
		fmt.Fprint(out, state.String())
		line, _ := reader.ReadString('\n')
		newState, err := state.NextState(strings.Trim(line, "\n"))
		if err != nil {
			fmt.Fprint(out, err.Error())
		}
		state = newState
	}
	fmt.Fprintf(out, "Goodbye!\n")
}

func main() {
	gameLoop(ttgo.MainMenu, os.Stdin, os.Stdout)
}
