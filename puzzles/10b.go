package puzzles

import (
	"advent/utils"
	"fmt"
	"strings"
)

func getPixelCoords(cycle int) (x, y int) {
	x = (cycle - 1) % 40
	y = int(cycle / 40)
	return
}

func (p *Processor) DrawPixel() {
	x, y := getPixelCoords(p.CycleCount)
	fmt.Println(x, y)
	if (p.InSprite()) {
		fmt.Println("Drawing pixel in position: ", x)
		p.Display.Pixels[y][x] = "#"
		fmt.Println("Current CRT row: ", p.Display.Pixels[y])
	}
}

func (p *Processor) InSprite() bool {
	c := p.CycleCount
	r := p.Register
	return c % 40 >= r && c % 40 <= r + 2
}

func (p *Processor) OutPutDisplay() string {
	var output string

	for _, x := range p.Display.Pixels {
		line := strings.Join(x, "")
		output = output + line + "\n"
	}

	return output
}

func createEmptyDisplay(width int, count int) Display {
	var pixels [][]string

	for y := 0; y < count; y++ {
		line := []string{}
		for x := 0; x < width; x++ {
			line = append(line, ".")
		}
		pixels = append(pixels, line)
	}

	return Display{pixels}
}

func Puzzle10b() string {
	data := utils.FileReader("data/10.txt")

	cpu := Processor{
		SignalCycleBase: 20, 
		SignalCycleMult: 40, 
		Register: 1,
		CycleCount: 1,
		Display: createEmptyDisplay(40, 6),
	}
	i := 0

	for {
		fmt.Println("")
		if (i < len(data)) {
			command, cycles, val := parseInstruction(data[i])
			cpu.addInstruction(command, cycles, val)
			fmt.Println("Start cycle: ", cpu.CycleCount)
		}

		cpu.cycle()

		i++
		if len(cpu.Instructions) < 1 && i >= len(data) {
			break
		}
	}

	return cpu.OutPutDisplay()
}