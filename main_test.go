package main

import (
	"advent/lib"
	"os"
	"testing"
)

// go test -bench BenchmarkPuzzle -args 01a

func BenchmarkPuzzle(b *testing.B) {
	code := os.Args[4]

	puzzle := lib.GetPuzzleFunc(code)

	for i := 0; i < b.N; i++ {
		puzzle()
	}
}