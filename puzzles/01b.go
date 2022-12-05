package puzzles

import (
	"advent/utils"
	"fmt"
	"sort"
	"strconv"
)

func Puzzle01b() string {
	data := utils.FileReader("data/01.txt")

	totals := []int{0}
	elf := 0

	// Go sort solution

	for _, cal := range data {
		if cal == "" {
			elf++
			totals = append(totals, 0)
			continue
		}

		calInt, err := strconv.Atoi(cal)

		if err != nil {
			fmt.Println(err)
		}

		totals[elf] = totals[elf] + calInt
	}

	sort.Ints(totals)

	first := totals[len(totals) - 1]
	second := totals[len(totals) - 2]
	third := totals[len(totals) - 3]
	
	// 0N solution, with branching

	// var first int
	// var second int
	// var third int

	// for _, cal := range data {

	// 	if cal == "" {

	// 		if totals[elf] > third && totals[elf] < second {
	// 			third = totals[elf]
	// 		}

	// 		if totals[elf] > second && totals[elf] < first {
	// 			third = second
	// 			second = totals[elf]
	// 		}

	// 		if totals[elf] > first {
	// 			second = first
	// 			first = totals[elf]
	// 		}

	// 		elf++
	// 		totals = append(totals, 0)
	// 		continue
	// 	}

	// 	calInt, err := strconv.Atoi(cal)

	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}

	// 	totals[elf] = totals[elf] + calInt
	// }

	answer := strconv.Itoa(first + second + third)

	return answer
}