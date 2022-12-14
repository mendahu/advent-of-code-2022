package puzzles

import (
	"advent/utils"
	"strconv"
)

func Puzzle03b() string {
	data := utils.FileReader("data/03.txt")

	var sumPriorities int32
	sumPriorities = 0

	group := 0;

	// loop through groups
	for {
		if (group  * 3 >= len(data)) {
			break;
		}
		
		// get overal index of elves in this group
		starti := group * 3
		e1i := starti
		e2i := starti + 1
		e3i := starti + 2

		// Three rucksacks for the group, arrays of runes
		elfRucks := [][]rune{
			[]rune(data[e1i]),
			[]rune(data[e2i]),
			[]rune(data[e3i]),
		}

		// Map for checking against, inventory
		rucksacks := map[int]map[rune]bool{}
		rucksacks[0] = map[rune]bool{}
		rucksacks[1] = map[rune]bool{}
		rucksacks[2] = map[rune]bool{}

		item := 0

		// loop through items
		ITEMS:
			for {
				// no more items, break out
				if len(elfRucks[0]) <= item && len(elfRucks[1]) <= item && len(elfRucks[2]) <= item {
					break;
				}

				elf := 0

				// loop through elves 1-3
				for {

					// Found badge or no more elves, break out
					if elf > 2 {
						break;
					}

					partner1i := ((elf + 1) % 3)
					partner2i := ((elf + 2) % 3)

					// current elf has items left
					if len(elfRucks[elf]) > item {

						currentItem := elfRucks[elf][item]

						p1val := rucksacks[partner1i][currentItem]
						p2val := rucksacks[partner2i][currentItem]

						if p1val && p2val {
							sumPriorities = sumPriorities + getPriority(currentItem)
							break ITEMS
						} else {
							rucksacks[elf][currentItem] = true
						}
					}

					elf++
				}
				item++
			}
		group++
	}

	answer := strconv.Itoa(int(sumPriorities))

	return answer
}