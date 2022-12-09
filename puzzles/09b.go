package puzzles

import (
	"advent/utils"
	"fmt"
	"strconv"
)

func Puzzle09b() string {
	data := utils.FileReader("data/09.txt")

	positionTracker := make(map[string]bool)
	positionTracker["[0,0]"] = true

	knots := [10]Position{
		{0,0},
		{0,0},
		{0,0},
		{0,0},
		{0,0},
		{0,0},
		{0,0},
		{0,0},
		{0,0},
		{0,0},
	}

	for _, move := range data {
		dir, count := parseCommand(move)
		
		// Loop each move individually
		for {
			if count < 1 {
				break
			}

			knots[0].move(dir) // initial knot move
			currDir := dir

			// Move each subsequent knot in response
			for k := 1; k < len(knots) && !areTouching(knots[k - 1], knots[k]); k++ {

				knots[k].move(currDir)

				// Correct diagonals edge case
				switch currDir {
				case "L":
					fallthrough
				case "R":
					if knots[k].y == knots[k - 1].y {
						break
					}

					if knots[k].y < knots[k - 1].y {
						knots[k].y++
						currDir = "U"
					} else {
						knots[k].y--
						currDir = "D"
					}
				case "U":
					fallthrough
				case "D":
					if knots[k].x == knots[k - 1].x {
						break
					}

					if knots[k].x < knots[k - 1].x {
						knots[k].x++
						currDir = "R"
					} else {
						knots[k].x--
						currDir = "L"
					}
				}
				
			}

			count--
			visitedCoord := getCoorLabel(knots[9])
			positionTracker[visitedCoord] = true
		}
	}
	fmt.Println(positionTracker)


	return strconv.Itoa(len(positionTracker))
}