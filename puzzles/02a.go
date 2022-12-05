package puzzles

import (
	"advent/utils"
	"strconv"
	"strings"
)

func Puzzle02a() string {
	data := utils.FileReader("data/02.txt")

	score := 0

	winVals := map[string]int{
		"A": 0,
		"B": 1,
		"C": 2,
		"X": 0,
		"Y": 1,
		"Z": 2,
	}

	for _, result := range data {

		// A Rock - B Paper - C Scissors
		// X Rock - Y Paper - Z Scissors
		// Win 6 - Draw 3 - Lose 0
		// Rock 1 - Paper 2 - Scissors 3

		// Modular Subtraction Method

		fields := strings.Fields(result)

		oppChoice := winVals[fields[0]]
		myChoice := winVals[fields[1]]

		switch ((myChoice - oppChoice) % 3) {
		case 0:
			score = score + 3
		case 1:
			score = score + 6
		case 2:
			score = score + 0
		}

		score = score + myChoice + 1


		// Older Brute Force method

		// switch result {
		// case "A X":
		// 	score = score + 3 + 1
		// case "A Y":
		// 	score = score + 6 + 2
		// case "A Z":
		// 	score = score + 0 + 3
		// case "B X":
		// 	score = score + 0 + 1
		// case "B Y":
		// 	score = score + 3 + 2
		// case "B Z":
		// 	score = score + 6 + 3
		// case "C X":
		// 	score = score + 6 + 1
		// case "C Y":
		// 	score = score + 0 + 2
		// case "C Z":
		// 	score = score + 3 + 3
		// }
	}

	answer := strconv.Itoa(score)

	return answer
}