package puzzles

import (
	"advent/utils"
	"bytes"
	"fmt"
)

func BreakShipCrateData(data []string) (shipData []string, moveData []string) {
	step := 0

	for _, datum := range data {
		if datum == "" {
			step++
			continue
		}

		if step == 0 {
			shipData = append(shipData, datum)
		} else {
			moveData = append(moveData, datum)
		}
	}

	return
}

type Ship struct {
	Stacks		[][]string
}

func BuildShip(shipData []string) Ship {	
	// amount of stacks left to right
	stackCount := ((len(shipData[0]) - 3) / 4) + 1
	// initiliaze empty slices
	stacks := make([][]string, stackCount)
	maxHeight := len(shipData) - 1
	
	// loop left to right through stacks
	for i := 0; i < stackCount; i++ {

		// loop bottom to top through the stack
		for j := maxHeight - 1; j >= 0; j-- {

			// Get stack Index and stack
			stacki := (i * 4) + 1
			crate := string(shipData[j][stacki])
			
			// Empty string indicates we've reached the top of the stack
			if crate == " " {
				break
			}
			
			// Place stacks into nested slice
			stacks[i] = append(stacks[i], crate)
		}
	}

	return Ship{stacks}
}

func (s Ship) GetTopCrate(stack int) (crate string) {
	return s.Stacks[stack][len(s.Stacks[stack]) - 1]
}

func (s Ship) MoveCrate(startStacki int, endStacki int, moves int) {
	for i := 0; i < moves; i++ {
		s.Stacks[endStacki] = append(s.Stacks[endStacki], s.GetTopCrate(startStacki))
		s.Stacks[startStacki] = s.Stacks[startStacki][:len(s.Stacks[startStacki]) - 1]
	}
}

func ParseMove (move string) (start int, end int, moves int) {
	_, err := fmt.Sscanf(move, "move %d from %d to %d", &moves, &start, &end,)

	if err != nil {
		panic(err)
	}

	return
}

func (s Ship) RevealTopCrates() string {
	var crates bytes.Buffer

	for i := 0; i < len(s.Stacks); i++ {
		crate := s.GetTopCrate(i)
		crates.WriteString(crate)
	}

	return crates.String()
}

func Puzzle05a() string {
	data := utils.FileReader("data/05.txt")

	shipData, moveData := BreakShipCrateData(data)

	ship := BuildShip(shipData)

	
	for _, move := range moveData {
		start, end, moves := ParseMove(move)
		ship.MoveCrate(start - 1, end - 1, moves)
	}

	return ship.RevealTopCrates()
}