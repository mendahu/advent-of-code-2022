package puzzles

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

type Directory struct {
	files		map[string]int
	dir			map[string]*Directory
	path		string
}

func (d Directory) CalcSize() int {
	files := 0
	dirs := 0

	for _, size := range d.files {
		files = files + size
	}

	for _, dir := range d.dir {
		dirs = dirs + dir.CalcSize()
	}

	return files + dirs
}

func MakeNewDirectory(path string) Directory {
	return Directory{
		files: make(map[string]int),
		dir: make(map[string]*Directory),
		path: path,
	}
}

func BuildDir(commands []string) Directory {
	root := MakeNewDirectory("root")
	var currentDir *Directory
	currentDir = &root

	for i := 1; i < len(commands); i++ {
		args := strings.Fields(commands[i])

		// is Command
		if args[0] == "$" {
			if args[1] == "ls" {
				continue
			}

			if (args[1]) == "cd" {

				if args[2] == ".." {
					paths := strings.Split(currentDir.path, "/")
					paths = paths[1:len(paths) - 1]
					currentDir = &root
					for _, val := range paths {
						currentDir = currentDir.dir[val]
					}
					continue
				}
	
				currentDir = currentDir.dir[args[2]]
				continue
			}
		}

		// is a directory listing
		if args[0] == "dir" {
			newPath := currentDir.path + "/" + args[1]
			newDir := MakeNewDirectory(newPath)
			currentDir.dir[args[1]] = &newDir
			continue
		}

		//is a file
		fileSize, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error converting string in file size")
		}
		currentDir.files[args[1]] = fileSize
	}

	return root
}

func (d Directory) GetTotalSizeBelow(limit int) int {
	size := 0

	currSize := d.CalcSize()
	if (currSize < limit) {
		size = size + currSize
	}

	for _, dir := range d.dir {
		size = size + dir.GetTotalSizeBelow(limit)
	}

	return size
}

func Puzzle07a() string {
	data := utils.FileReader("data/07.txt")

	root := BuildDir(data)

	size := root.GetTotalSizeBelow(100000)


	return strconv.Itoa(size);
}