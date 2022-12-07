package puzzles

import (
	"advent/utils"
	"strconv"
)

func (d Directory) FindDirToDelete(spaceNeeded int, maxSize int) int {
	size := d.CalcSize()
	candidateSize := maxSize

	if (size > spaceNeeded && size < candidateSize) {
		candidateSize = size
	}

	for _, dir := range d.dir {
		dirSize := dir.FindDirToDelete(spaceNeeded, candidateSize)
		if (dirSize > spaceNeeded && dirSize < candidateSize) {
			candidateSize = dirSize
		}
	}

	return candidateSize
}

func Puzzle07b() string {
	data := utils.FileReader("data/07.txt")

	root := BuildDir(data)
	size := root.CalcSize()
	freeSpace := 70000000 - size
	spaceNeeded := 30000000 - freeSpace

	candidateSize := root.FindDirToDelete(spaceNeeded, size)

	return strconv.Itoa(candidateSize)
}