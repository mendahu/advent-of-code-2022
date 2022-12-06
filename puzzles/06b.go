package puzzles

import (
	"advent/utils"
	"strconv"
)

func Puzzle06b() string {
	data := utils.FileReader("data/06.txt")

	device := CommDevice{}

	stremStart := device.GetStreamStart(data[0], 14)

	return strconv.Itoa(stremStart)
}