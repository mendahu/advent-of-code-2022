package main

import (
	"advent/puzzles"
	"fmt"
	"os"
)

func main() {
	code := os.Args[1]

	puzzles := map[string]func()int64{
		"01a": puzzles.Puzzle01a,
		"01b": puzzles.Puzzle01b,
		"02a": puzzles.Puzzle02a,
		"02b": puzzles.Puzzle02b,
	}

	answer := puzzles[code]()
	fmt.Println(answer)
}