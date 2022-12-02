package puzzles

import (
	"advent/utils"
	"fmt"
	"sort"
	"strconv"
)

func Puzzle01a() int64 {
	data := utils.FileReader("data/01.txt")

	totals := []int{0}
	elf := 0

	// Go sort library solution

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

	max := totals[len(totals) - 1]

	// 0N solution with single loop and branching

	// var max int

	// for _, cal := range data {

	// 	if cal == "" {
	// 		if totals[elf] > max {
	// 			max = totals[elf]
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

	return int64(max)
}