package puzzles

import (
	"advent/utils"
	"fmt"
)

type SectionAssignmentPair struct {
	r1sta   int
	r1end		int
	r2sta  	int
	r2end		int
}

func parseAssignmentPair(raw string) SectionAssignmentPair {
	var r1sta, r1end, r2sta, r2end int

	_, err := fmt.Sscanf(raw, "%d-%d,%d-%d", &r1sta, &r1end, &r2sta, &r2end)

	if err != nil {
		panic(err)
	}

	return SectionAssignmentPair{r1sta, r1end, r2sta, r2end}
}

func (ap SectionAssignmentPair) oneOverlapsTwo() bool {
	return ap.r1sta <= ap.r2sta && ap.r1end >= ap.r2end
}

func (ap SectionAssignmentPair) twoOverlapsOne() bool {
	return ap.r2sta <= ap.r1sta && ap.r2end >= ap.r1end
}

func Puzzle04a() int64 {
	data := utils.FileReader("data/04.txt")

	sumOverlaps := 0

	for _, raw := range data {
		ap := parseAssignmentPair(raw)
		if ap.oneOverlapsTwo() || ap.twoOverlapsOne() {
			sumOverlaps++
		}
	}

	return int64(sumOverlaps)
}