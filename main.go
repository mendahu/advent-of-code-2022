package main

import (
	"advent/lib"
	"fmt"
	"os"
)

func main() {
	code := os.Args[1]

	puzzle := lib.GetPuzzleFunc(code)
	answer := puzzle()
	fmt.Println(answer)
}