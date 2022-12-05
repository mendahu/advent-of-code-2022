package puzzles

import (
	"advent/utils"
	"strconv"
	"strings"
)

func Puzzle02b() string {
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
		// X Lose - Y Draw - Z Win
		// Win 6 - Draw 3 - Lose 0
		// Rock 1 - Paper 2 - Scissors 3

		// Modular Addition Method

		fields := strings.Fields(result)

		oppChoice := winVals[fields[0]]
		result := winVals[fields[1]]

		score = score + ((oppChoice + result + 2) % 3) + 1 + (result * 3)

		// Older Brute Force method

		// switch result {
		// case "A X":
		// 	score = score + 0 + 3
		// case "A Y":
		// 	score = score + 3 + 1
		// case "A Z":
		// 	score = score + 6 + 2
		// case "B X":
		// 	score = score + 0 + 1
		// case "B Y":
		// 	score = score + 3 + 2
		// case "B Z":
		// 	score = score + 6 + 3
		// case "C X":
		// 	score = score + 0 + 2
		// case "C Y":
		// 	score = score + 3 + 3
		// case "C Z":
		// 	score = score + 6 + 1
		// }
	}
	
	answer := strconv.Itoa(score)

	return answer
}