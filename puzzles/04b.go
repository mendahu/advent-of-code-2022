package puzzles

import "advent/utils"

func (ap SectionAssignmentPair) hasAnyOverlap() bool {
	return (ap.r1end >= ap.r2sta && ap.r1sta <= ap.r2end) || (ap.r2end >= ap.r1sta && ap.r2sta <= ap.r1end)
}

func Puzzle04b() int64 {
	data := utils.FileReader("data/04.txt")

	sumOverlaps := 0

	for _, raw := range data {
		ap := parseAssignmentPair(raw)
		if ap.hasAnyOverlap() {
			sumOverlaps++
		}
	}

	return int64(sumOverlaps)
}