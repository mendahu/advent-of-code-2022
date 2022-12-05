package puzzles

import (
	"advent/utils"
	"strconv"
)

func getPriority(letter rune) int32 {
	if letter < 91 && letter > 64 {
		return letter - 65 + 27
	} else if letter < 123 && letter > 96 {
		return letter - 96
	}

	return 0
}

func Puzzle03a() string {
	data := utils.FileReader("data/03.txt")

	var sumPriorities int32
	sumPriorities = 0

	for _, rucksack := range data {
		items := []rune(rucksack)
		itemCount := len(items)
		compartmentSize := itemCount / 2

		itemsCountedLeft := map[rune]bool {}
		itemsCountedRight := map[rune]bool {}


		for i := 0; i < compartmentSize; i++ {

			// Check left side
			itemLeft := items[i]

			if itemsCountedRight[itemLeft] {
				sumPriorities = sumPriorities + getPriority(itemLeft)
				break;
			} else {
				itemsCountedLeft[itemLeft] = true
			}

			// Check right side
			itemRight := items[i + compartmentSize]

			if itemsCountedLeft[itemRight] {
				sumPriorities = sumPriorities + getPriority(itemRight)
				break;
			} else {
				itemsCountedRight[itemRight] = true
			}
		}
	}

	answer := strconv.Itoa(int(sumPriorities))

	return answer
}