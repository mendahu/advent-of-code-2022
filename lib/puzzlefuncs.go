package lib

import "advent/puzzles"

func GetPuzzleFunc(code string) func()int64 {
	puzzles := map[string]func()int64{
		"01a": puzzles.Puzzle01a,
		"01b": puzzles.Puzzle01b,
		"02a": puzzles.Puzzle02a,
		"02b": puzzles.Puzzle02b,
		"03a": puzzles.Puzzle03a,
		"03b": puzzles.Puzzle03b,
		"04a": puzzles.Puzzle04a,
		"04b": puzzles.Puzzle04b,
	}

	return puzzles[code]
}