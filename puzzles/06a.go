package puzzles

import (
	"advent/utils"
	"strconv"
)

type CommDevice struct {}

// Implement a Go Set
// type void struct{}
// var member void

func (d CommDevice) GetStreamStart(stream string, identifier int) int {
	chars := []rune(stream)

	// Using Map
	CHECKER:
		for i := 0; i < len(chars) - identifier; i++ {
			check := map[rune]bool {}

			for j := 0; j < identifier; j++ {
				if check[chars[i+j]] {
					continue CHECKER
				} else {
					check[chars[i+j]] = true
				}
			}

			return i + identifier
		}
	
	// Using Go Set - ended up twice as slow on part 2
	// for i := 0; i < len(chars) - identifier + 1; i++ {
	// 	set := make(map[rune]void)

	// 	for j := 0; j < identifier; j++ {
	// 		set[chars[i+j]] = member
	// 	}

	// 	if len(set) == identifier {
	// 		return i + identifier
	// 	}
	// }

	return len(chars)
}


func Puzzle06a() string {
	data := utils.FileReader("data/06.txt")

	device := CommDevice{}

	stremStart := device.GetStreamStart(data[0], 4)

	return strconv.Itoa(stremStart)
}