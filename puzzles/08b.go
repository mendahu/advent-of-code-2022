package puzzles

import (
	"advent/utils"
	"strconv"
)

func (t TreeGrid) GetBestScenicScore() int {
	bestScore := 0;

	for _, row := range t.grid {
		for _, tree := range row {
			currentScore := tree.GetScenicScore()
			if currentScore > bestScore {
				bestScore = currentScore
			}
		}
	}

	return bestScore
}

func Puzzle08b() string {
	data := utils.FileReader("data/08.txt")

	forest := MakeGrid(data)
	bestScenicScore := forest.GetBestScenicScore()

	return strconv.Itoa(bestScenicScore)
}