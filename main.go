package main

import (
	"advent/puzzles"
	"fmt"
	"os"
)

func main() {
	puzzle := os.Args[1]

	p := map[string]func()int64{
		"01a": puzzles.Puzzle01a,
		"01b": puzzles.Puzzle01b,
		"02a": puzzles.Puzzle02a,
		"02b": puzzles.Puzzle02b}

	answer := p[puzzle]()
	fmt.Println(answer)
}