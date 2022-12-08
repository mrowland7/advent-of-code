package main

import (
	//"regexp"
	"fmt"
	"strconv"
	"strings"
)

type File struct {
	name string
	size int
}

type Dir struct {
	name    string
	subdirs []*Dir
	files   []*File
	parent  *Dir
	subsize int
}

func printDir(d *Dir) string {
	subdirs := []string{}
	for _, d := range d.subdirs {
		subdirs = append(subdirs, d.name)
	}
	files := []string{}
	for _, f := range d.files {
		files = append(files, f.name)
	}
	return fmt.Sprintf("{Name: %v, subdirs: %v, files: %v, parent: %v", d.name, subdirs, files, d.parent.name)
}

func main() {
	lines, err := getLines("day7_input.txt")
	assertOk(err)
	// Starts in root
	// $ cd /
	root := &Dir{
		name:    "root",
		parent:  &Dir{name: "n/a"},
		subsize: -1,
	}
	currentDir := root
	lineIdx := 0
	// Build the tree
	for lineIdx < len(lines) {
		nextCommand := lines[lineIdx]
		fmt.Println(nextCommand, "## from dir", printDir(currentDir))
		if nextCommand[2:] == "ls" {
			lineIdx++
			currentDir.subdirs = []*Dir{}
			currentDir.files = []*File{}
			for ; lineIdx < len(lines) && lines[lineIdx][0] != '$'; lineIdx++ {
				lsLine := lines[lineIdx]
				fmt.Println("=== ls line", lineIdx, lsLine)
				if lsLine[0:3] == "dir" {
					currentDir.subdirs = append(currentDir.subdirs, &Dir{
						name:    lsLine[4:],
						subdirs: []*Dir{},
						files:   []*File{},
						parent:  currentDir,
						subsize: -1,
					})
				} else {
					pieces := strings.Split(lsLine, " ")
					size, err := strconv.Atoi(pieces[0])
					assertOk(err)
					currentDir.files = append(currentDir.files, &File{
						name: pieces[1],
						size: size,
					})
				}
			}
		} else if nextCommand[2:] == "cd .." {
			currentDir = currentDir.parent
			lineIdx++
		} else if nextCommand[2:] == "cd /" {
			currentDir = root
			lineIdx++
		} else if nextCommand[2:4] == "cd" {
			for _, dir := range currentDir.subdirs {
				if dir.name == nextCommand[5:] {
					currentDir = dir
					break
				}
			}
			lineIdx++
		} else {
			panic("unknown command" + nextCommand)
		}
	}
	// Traverse the tree
	need := 30000000 - (70000000 - subSize(root))
	fmt.Println(subSize(root))
	fmt.Println(totalUnder10k(root))
	fmt.Println(need)
}

func subSize(d *Dir) int {
	s := 0
	for _, f := range d.files {
		s += f.size
	}
	for _, subdir := range d.subdirs {
		s += subSize(subdir)
	}
	d.subsize = s
	fmt.Println(s, "is the subsize of", d.name)
	return s
}

func totalUnder10k(d *Dir) int {
	s := 0
	if d.subsize <= 100000 {
		s += d.subsize
	}
	for _, subdir := range d.subdirs {
		s += totalUnder10k(subdir)
	}
	return s
}
