package puzzles

import (
	"advent/utils"
	"fmt"
	"strconv"
)

type Position struct {
	x int
	y int
}

func (p *Position) move(dir string) {
	switch dir {
		case "L": 
			p.x--
		case "R": 
			p.x++
		case "U": 
			p.y++
		case "D": 
			p.y--
		}
}

func parseCommand(move string) (dir string, count int) {
	_, err := fmt.Sscanf(move, "%s %d", &dir, &count)

	if err != nil {
		panic(err)
	}

	return
}

func getCoorLabel(coords Position) string {
	x := strconv.Itoa(coords.x)
	y := strconv.Itoa(coords.y)
	return "[" + x + "," + y + "]"
}

func areTouching(first Position, second Position) bool {
	isXAdjacent := second.x >= first.x - 1 && second.x <= first.x + 1
	isYAdjacent := second.y >= first.y - 1 && second.y <= first.y + 1
	return isXAdjacent && isYAdjacent
}

func Puzzle09a() string {
	data := utils.FileReader("data/09.txt")

	positionTracker := make(map[string]bool)
	positionTracker["[0,0]"] = true

	hPos := Position{0,0}
	tPos := Position{0,0}

	for _, move := range data {
		dir, count := parseCommand(move)
		
		for {
			if count < 1 {
				break
			}
			
			hPos.move(dir)
			count--

			if !areTouching(hPos, tPos) {
				tPos.move(dir)

				// Correct diagonals edge case
				switch dir {
				case "L":
					fallthrough
				case "R":
					tPos.y = hPos.y
				case "U":
					fallthrough
				case "D":
					tPos.x = hPos.x
				}
			}

			visitedCoord := getCoorLabel(tPos)
			positionTracker[visitedCoord] = true
		}
	}


	return strconv.Itoa(len(positionTracker))
}