package puzzles

import (
	"advent/utils"
)

func (s Ship) GetCrateSlice(stacki int, count int) []string {
	starti := len(s.Stacks[stacki]) - count
	return s.Stacks[stacki][starti:len(s.Stacks[stacki])]
}

func (s Ship) MoveCrates(startStacki int, endStacki int, moves int) {
	topCrates := s.GetCrateSlice(startStacki, moves)
	
	s.Stacks[endStacki] = append(s.Stacks[endStacki], topCrates...)
	s.Stacks[startStacki] = s.Stacks[startStacki][:len(s.Stacks[startStacki]) - moves]
}

func Puzzle05b() string {
	data := utils.FileReader("data/05.txt")

	shipData, moveData := BreakShipCrateData(data)

	ship := BuildShip(shipData)

	
	for _, move := range moveData {
		start, end, moves := ParseMove(move)
		ship.MoveCrates(start - 1, end - 1, moves)
	}

	return ship.RevealTopCrates()
}