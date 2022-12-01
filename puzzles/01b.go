package puzzles

import (
	"advent/utils"
	"fmt"
	"strconv"
)

func Puzzle01b() int64 {
	data := utils.FileReader("data/01.txt")

	totals := []int{0}
	elf := 0
	var first int
	var second int
	var third int

	for _, cal := range data {

		if cal == "" {

			if totals[elf] > third && totals[elf] < second {
				third = totals[elf]
			}

			if totals[elf] > second && totals[elf] < first {
				third = second
				second = totals[elf]
			}

			if totals[elf] > first {
				second = first
				first = totals[elf]
			}

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

	return int64(first + second + third)

}