package puzzles

import (
	"advent/utils"
	"fmt"
	"strconv"
)

func Puzzle11b() string {
	data := utils.FileReader("data/11.txt")

	monkeys := buildMonkeys(data)
	mod := getCommonModule(monkeys)

	r := 0

	for {
		if r >= 10000 {
			break;
		}

		for i := 0; i < len(monkeys); i++ {
			monkeys[i].inspectItems(monkeys, 1, mod)
		}

		r++
	}

	fmt.Println(monkeys)
	monkeyBusiness := strconv.Itoa(getMonkeyBusiness(monkeys))
	return monkeyBusiness
}