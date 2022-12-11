package puzzles

import (
	"advent/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return int(x / y)
}

type Operator struct {
	Left   string
	Right  string
	OpFunc func(x, y int) int
}

type Monkey struct {
	Inspections 		int
	Items						[]int
	Operator				Operator
	Divider					int
	Destinations	 	map[bool]int
}

func parseItems(data string) (items []int) {
	stringItems := string([]rune(data)[18:])
	arrStrItems := strings.Split(stringItems, ", ")

	for _, item := range arrStrItems {
		i, err := strconv.Atoi(item)

		if err != nil {
			panic(err)
		}
		items = append(items, i)
	}

	return
}

func parseOperation(data string) (l, r string, opFunc func(x,y int) int) {
	opFuncs := map[string]func(x,y int)int {
		"+": add,
		"-": subtract,
		"*": multiply,
		"/": divide,
	}

	stringOp := string([]rune(data)[19:])
	
	var op string
	_, err := fmt.Sscanf(stringOp, "%s %s %s", &l, &op, &r)

	if err != nil {
		panic(err)
	}

	opFunc = opFuncs[op]

	return
}

func parseDivider(data string) int {
	stringDiv := string([]rune(data)[21:])
	div, err := strconv.Atoi(stringDiv)

	if err != nil {
		panic(err)
	}

	return div
}

func parseDestination(data []string) map[bool]int {
	trueDes := string([]rune(data[0])[29:])
	falseDes := string([]rune(data[1])[30:])

	t, err := strconv.Atoi(trueDes)

	if err != nil {
		panic(err)
	}

	f, err := strconv.Atoi(falseDes)

	if err != nil {
		panic(err)
	}

	dests := map[bool]int {
		true: t,
		false: f,
	}

	return dests
}

func parseMonkey(data []string) (items []int, l, r string, opFunc func(x, y int) int, div int, dests map[bool]int) {
	items = parseItems(data[0])
	l, r, opFunc = parseOperation(data[1])
	div = parseDivider(data[2])
	dests = parseDestination(data[3:5])
	return
}

func buildMonkeys(data []string) []Monkey {
	i := 1

	monkeys := []Monkey{}

	for {
		if i > len(data) {
			break;
		}

		items, l, r, opFunc, div, dests := parseMonkey(data[i:i+4])
		operator := Operator{l, r, opFunc}

		monkeys = append(monkeys, Monkey{0, items, operator, div, dests})

		i = i + 7
	}

	return monkeys
}

func (m *Monkey) addItem(worry int) {
	m.Items = append(m.Items, worry)
}

func inspectItem(item int, operator Operator) int {
	var l, r int

	if operator.Left == "old" {
		l = item
	} else {
		lInt, err := strconv.Atoi(operator.Left)

		if err != nil {
			panic(err)
		}

		l = lInt
	}

	if operator.Right == "old" {
		r = item
	} else {
		rInt, err := strconv.Atoi(operator.Right)

		if err != nil {
			panic(err)
		}

		r = rInt
	}

	return operator.OpFunc(l, r)
}

func (m *Monkey) inspectItems(monkeys []Monkey, worryDiv int, mod int) {
	for _, item := range m.Items {
		postInspection := inspectItem(item, m.Operator)
		postDivider := int(postInspection / worryDiv)
		
		postModulator := postDivider % mod
		passesTest := postDivider % m.Divider == 0
		destination := m.Destinations[passesTest]
		monkeys[destination].addItem(postModulator)
		m.Inspections++
	}
	m.Items = nil
}

func getMonkeyBusiness(monkeys []Monkey) int {
	var scores []int

	for _, m := range monkeys {
		scores = append(scores, m.Inspections)
	}

	sort.Ints(scores)
	best := scores[len(scores) - 1]
	secondBest := scores[len(scores) - 2]
	return best * secondBest
}

func getCommonModule(monkeys []Monkey) int {
	mod := 1
	for _, m := range monkeys {
		mod = mod * m.Divider
	}

	return mod
}

func Puzzle11a() string {
	data := utils.FileReader("data/11.txt")

	monkeys := buildMonkeys(data)
	mod := getCommonModule(monkeys)

	r := 0

	for {
		if r >= 20 {
			break;
		}

		for i := 0; i < len(monkeys); i++ {
			monkeys[i].inspectItems(monkeys, 3, mod)
		}

		r++
	}

	fmt.Println(monkeys)
	monkeyBusiness := strconv.Itoa(getMonkeyBusiness(monkeys))
	return monkeyBusiness
}