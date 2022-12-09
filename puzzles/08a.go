package puzzles

import (
	"advent/utils"
	"strconv"
)

type TreeViews struct {
	Left		int
	Right		int
	Top			int
	Bottom	int
}

type Tree struct {
	Height	rune
	Visible	bool
	Views		TreeViews
}

func (t Tree) GetScenicScore() int {
	return t.Views.Left * t.Views.Right * t.Views.Top * t.Views.Bottom
}

type TreeGrid struct {
	grid		[][]Tree
}

func (t TreeGrid) CountVisible() int {
	count := 0;

	for _, row := range t.grid {
		for _, tree := range row {
			if tree.Visible {
				count++
			}
		}
	}

	return count
}

func MakeGrid(data []string) TreeGrid {
	var forest TreeGrid
	forest.grid = make([][]Tree, len(data))

	// Create grid via rows, set visibility from left and right
	for i := 0; i < len(data); i++ {
		row := make([]Tree, len(data[i]))
		dataRow := []rune(data[i])

		var leftBound int32
		leftBound = -1

		// Loop through row from left, create Tree and set visibility bounds
		for l := 0; l < len(dataRow); l++ {
			height := dataRow[l] - 48

			// initialize new tree
			row[l] = Tree{
				Height: height, 
				Visible: height > leftBound, 
				Views: TreeViews{0,0,0,0},
			}

			// set new Visibility
			if row[l].Height > leftBound {
				leftBound = row[l].Height
			}

			viewCount := 0
			blocked := false
			checkIndex := l - 1;
			for {
				if blocked || checkIndex < 0 {
					break
				}

				if row[checkIndex].Height >= row[l].Height {
					blocked = true
				} else {
					checkIndex--
				}
				viewCount++
			}

			row[l].Views.Left = viewCount
		}

		var rightBound int32
		rightBound = -1

		// Loop from right, set visibility
		for r := len(dataRow) - 1; r >= 0; r-- {

			// Once the tree is max height, no need to keep checking
			if rightBound != 9 && row[r].Height > rightBound {
				row[r].Visible = true
				rightBound = row[r].Height
			}

			// Check right views
			viewCount := 0
			blocked := false
			checkIndex := r + 1
			for {
				if blocked || checkIndex > len(dataRow) - 1 {
					break
				}

				if row[checkIndex].Height >= row[r].Height {
					blocked = true
				} else {
					checkIndex++
				}
				viewCount++
			}

			row[r].Views.Right = viewCount
		}

		forest.grid[i] = row

	}

	// Set visibility from top and bottom
	for col := 0; col < len(data); col++ {
		var upperBound int32
		upperBound = -1

		// Loop from top to bottom, set visibility
		for t := 0; t < len(data); t++ {

			// Once the tree is max height, no need to keep checking
			if upperBound != 9 && forest.grid[t][col].Height > upperBound {
				forest.grid[t][col].Visible = true
				upperBound = forest.grid[t][col].Height
			}

			// Check top views
			viewCount := 0
			blocked := false
			checkIndex := t - 1
			for {
				if blocked || checkIndex < 0 {
					break
				}

				if forest.grid[checkIndex][col].Height >= forest.grid[t][col].Height {
					blocked = true
				} else {
					checkIndex--
				}
				viewCount++
			}

			forest.grid[t][col].Views.Top = viewCount
		}

		var lowerBound int32
		lowerBound = -1

		// Loop from bottom to top, set visibility
		for b := len(data) - 1; b >= 0; b-- {

			// Once the tree is max height, no need to keep checking
			if lowerBound != 9 && forest.grid[b][col].Height > lowerBound {
				forest.grid[b][col].Visible = true
				lowerBound = forest.grid[b][col].Height
			}

			// Check bottom views
			viewCount := 0
			blocked := false
			checkIndex := b + 1
			for {
				if blocked || checkIndex > len(data) - 1 {
					break
				}

				if forest.grid[checkIndex][col].Height >= forest.grid[b][col].Height {
					blocked = true
				} else {
					checkIndex++
				}
				viewCount++
			}

			forest.grid[b][col].Views.Bottom = viewCount
		}
	}

	return forest
}

func Puzzle08a() string {
	data := utils.FileReader("data/08.txt")

	forest := MakeGrid(data)
	visibleTrees := forest.CountVisible()



	return strconv.Itoa(visibleTrees)
}