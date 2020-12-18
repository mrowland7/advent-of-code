package main

import (
	"log"
	//"regexp"
	"fmt"
	//	"strconv"
	"strings"
)

const fieldSize = 30

func initHyperfield() [][][][]string {
	hyperfield := make([][][][]string, fieldSize)
	for w := 0; w < fieldSize; w++ {
		hyperfield[w] = initCubefield()
	}
	return hyperfield
}

func initCubefield() [][][]string {
	cubefield := make([][][]string, fieldSize)
	for i := 0; i < fieldSize; i++ {
		cubefield[i] = make([][]string, fieldSize)
		for j := 0; j < fieldSize; j++ {
			cubefield[i][j] = make([]string, fieldSize)
			for k := 0; k < fieldSize; k++ {
				cubefield[i][j][k] = "."
			}
		}
	}
	return cubefield
}

func printCubefield(cf [][][]string) {
	for i := 0; i < fieldSize; i++ {
		fmt.Println("z=", i)
		for j := 0; j < fieldSize; j++ {
			if i >= fieldSize/2-2 && i <= fieldSize/2+2 {
				fmt.Println(strings.Join(cf[i][j], ""))
			}
		}
	}
}

func inBounds(x int) bool {
	return x >= 0 && x < fieldSize
}

func main() {
	lines, err := getLines("day17_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	hf := initHyperfield()
	initialX := fieldSize / 2
	initialY := fieldSize / 2
	initialZ := fieldSize / 2
	initialW := fieldSize / 2
	for i, line := range lines {
		for j, str := range strings.Split(line, "") {
			fmt.Println(i, j, str)
			hf[initialW][initialZ][i+initialX][j+initialY] = str
		}
	}
	//printCubefield(cubefield)

	type neighbor struct {
		dx int
		dy int
		dz int
		dw int
	}
	dirs := []neighbor{
		{1, 1, 1, -1}, {1, 1, 0, -1}, {1, 1, -1, -1},
		{1, 0, 1, -1}, {1, 0, 0, -1}, {1, 0, -1, -1},
		{1, -1, 1, -1}, {1, -1, 0, -1}, {1, -1, -1, -1},
		{0, 1, 1, -1}, {0, 1, 0, -1}, {0, 1, -1, -1},
		{0, 0, 1, -1}, {0, 0, 0, -1}, {0, 0, -1, -1},
		{0, -1, 1, -1}, {0, -1, 0, -1}, {0, -1, -1, -1},
		{-1, 1, 1, -1}, {-1, 1, 0, -1}, {-1, 1, -1, -1},
		{-1, 0, 1, -1}, {-1, 0, 0, -1}, {-1, 0, -1, -1},
		{-1, -1, 1, -1}, {-1, -1, 0, -1}, {-1, -1, -1, -1},

		{1, 1, 1, 0}, {1, 1, 0, 0}, {1, 1, -1, 0},
		{1, 0, 1, 0}, {1, 0, 0, 0}, {1, 0, -1, 0},
		{1, -1, 1, 0}, {1, -1, 0, 0}, {1, -1, -1, 0},
		{0, 1, 1, 0}, {0, 1, 0, 0}, {0, 1, -1, 0},
		{0, 0, 1, 0} /*{0, 0, 0, 0},*/, {0, 0, -1, 0},
		{0, -1, 1, 0}, {0, -1, 0, 0}, {0, -1, -1, 0},
		{-1, 1, 1, 0}, {-1, 1, 0, 0}, {-1, 1, -1, 0},
		{-1, 0, 1, 0}, {-1, 0, 0, 0}, {-1, 0, -1, 0},
		{-1, -1, 1, 0}, {-1, -1, 0, 0}, {-1, -1, -1, 0},

		{1, 1, 1, 1}, {1, 1, 0, 1}, {1, 1, -1, 1},
		{1, 0, 1, 1}, {1, 0, 0, 1}, {1, 0, -1, 1},
		{1, -1, 1, 1}, {1, -1, 0, 1}, {1, -1, -1, 1},
		{0, 1, 1, 1}, {0, 1, 0, 1}, {0, 1, -1, 1},
		{0, 0, 1, 1}, {0, 0, 0, 1}, {0, 0, -1, 1},
		{0, -1, 1, 1}, {0, -1, 0, 1}, {0, -1, -1, 1},
		{-1, 1, 1, 1}, {-1, 1, 0, 1}, {-1, 1, -1, 1},
		{-1, 0, 1, 1}, {-1, 0, 0, 1}, {-1, 0, -1, 1},
		{-1, -1, 1, 1}, {-1, -1, 0, 1}, {-1, -1, -1, 1},
	}
	for round := 1; round <= 6; round++ {
		fmt.Println("=========================")
		nexthf := initHyperfield()
		newActiveCount := 0
		for w := 0; w < fieldSize; w++ {
			for i := 0; i < fieldSize; i++ {
				for j := 0; j < fieldSize; j++ {
					for k := 0; k < fieldSize; k++ {
						numActiveNeighbors := 0
						oldValue := hf[w][i][j][k]
						for _, d := range dirs {
							if inBounds(i+d.dx) && inBounds(j+d.dy) && inBounds(k+d.dz) && inBounds(w+d.dw) &&
								hf[w+d.dw][i+d.dx][j+d.dy][k+d.dz] == "#" {
								numActiveNeighbors++
							}
						}
						if oldValue == "#" && (numActiveNeighbors == 2 || numActiveNeighbors == 3) {
							nexthf[w][i][j][k] = "#"
							newActiveCount++
						} else if oldValue == "." && numActiveNeighbors == 3 {
							nexthf[w][i][j][k] = "#"
							newActiveCount++
						} else {
							nexthf[w][i][j][k] = "."
						}
					}
				}
			}
		}
		hf = nexthf
		//printCubefield(cubefield)
		fmt.Println("After round", round, "active count is", newActiveCount)
	}
}
