package puzzles

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	Command		string
	Cycles		int
	Value			int
}

func (i *Instruction) cycle() {
	i.Cycles--
}

type Display struct {
	Pixels		[][]string
}

type Processor struct {
	SignalCycleBase		int
	SignalCycleMult		int
	SignalStrength		int
	CycleCount				int
	Register 					int
	Instructions			[]Instruction
	Display						Display
}

func (p Processor) getSignalStrength() int {
	return p.CycleCount * p.Register
}

func (p *Processor) cycle() {
	
	p.DrawPixel()
	p.CycleCount++
	
	p.Instructions[0].cycle()
	if p.Instructions[0].Cycles == 0 {
		p.Register = p.Register + p.Instructions[0].Value
		fmt.Println("Finished Executing ", p.Instructions[0].Command, p.Instructions[0].Value)
		p.Instructions = p.Instructions[1:]
	}
	
	if p.CycleCount >= p.SignalCycleBase && (p.CycleCount - p.SignalCycleBase) % p.SignalCycleMult == 0 {
		p.SignalStrength = p.SignalStrength + p.getSignalStrength()
		fmt.Println(p.CycleCount, p.Register, p.getSignalStrength())
	}
}

func (p *Processor) addInstruction(command string, cycles int, value int) {
	p.Instructions = append(p.Instructions, Instruction{command, cycles, value})
}

func parseInstruction(ins string) (command string, cycles int, val int) {
	fields := strings.Fields(ins)
	if fields[0] == "noop" {
		return fields[0], 1, 0
	}
	val, err := strconv.Atoi(fields[1])
	if err != nil {
		panic(err)
	}
	return fields[0], 2, val
}

func Puzzle10a() string {
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
		fmt.Println("Start of Cycle: ")
		if (i < len(data)) {
			command, cycles, val := parseInstruction(data[i])
			fmt.Println("Initiating Command: ", command, cycles, val)
			cpu.addInstruction(command, cycles, val)
		}

		cpu.cycle()
		fmt.Println("End of Cycle :")

		i++
		if len(cpu.Instructions) < 1 && i >= len(data) {
			break
		}
	}

	return strconv.Itoa(cpu.SignalStrength)
}