package puzzles

import (
	"advent/utils"
	"strconv"
)

func Puzzle09b() string {
	data := utils.FileReader("data/09.txt")

	positionTracker := make(map[Position]bool)
	
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
	
	positionTracker[knots[9]] = true

	for _, move := range data {
		dir, count := parseCommand(move)
		
		// Loop each move individually
		for c := 0; c < count; c++ {
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

			positionTracker[knots[9]] = true
		}
	}

	return strconv.Itoa(len(positionTracker))
}