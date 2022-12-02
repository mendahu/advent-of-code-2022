package puzzles

import "advent/utils"

func Puzzle02b() int64 {
	data := utils.FileReader("data/02.txt")

	score := 0

	for _, result := range data {

		// A Rock - B Paper - C Scissors
		// X Lose - Y Draw - Z Win
		// Win 6 - Draw 3 - Lose 0
		// Rock 1 - Paper 2 - Scissors 3

		switch result {
		case "A X":
			score = score + 0 + 3
		case "A Y":
			score = score + 3 + 1
		case "A Z":
			score = score + 6 + 2
		case "B X":
			score = score + 0 + 1
		case "B Y":
			score = score + 3 + 2
		case "B Z":
			score = score + 6 + 3
		case "C X":
			score = score + 0 + 2
		case "C Y":
			score = score + 3 + 3
		case "C Z":
			score = score + 6 + 1
		}
	}
	
	return int64(score)
}