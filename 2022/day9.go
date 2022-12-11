package main

import (
	//"regexp"
	"fmt"
	"strconv"
	//	"strings"
)

type Point struct {
	x int
	y int
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func updateT(h, t *Point) {
	if abs(h.x-t.x) == 2 {
		t.x = (t.x + h.x) / 2
		if abs(h.y-t.y) == 1 {
			t.y = h.y
		}
		if abs(h.y-t.y) == 2 {
			t.y = (t.y + h.y) / 2
		}
	} else if abs(h.y-t.y) == 2 {
		t.y = (t.y + h.y) / 2
		if abs(h.x-t.x) == 1 {
			t.x = h.x
		}
		if abs(h.x-t.x) == 2 {
			t.x = (t.x + h.x) / 2
		}
	}
}

func printState(head *Point, tails *[]Point) {
	for y := 10; y > -10; y-- {
		for x := -10; x < 10; x++ {
			if head.x == x && head.y == y {
				fmt.Printf("H")
			} else {
				printed := false
				for i := 0; i < len(*tails); i++ {
					if (*tails)[i].x == x && (*tails)[i].y == y {
						fmt.Printf(strconv.Itoa(i + 1))
						printed = true
						break
					}
				}
				if !printed {
					fmt.Printf(".")
				}
			}
		}
		fmt.Println()
	}
}

func main() {
	lines, err := getLines("day9_input.txt")
	//lines, err := getLines("day9_dbg.txt")
	//lines, err := getLines("day9_dbg2.txt")
	//lines, err := getLines("day9_dbg3.txt")
	assertOk(err)
	tVisited := map[Point]struct{}{
		Point{0, 0}: struct{}{},
	}
	currH := Point{0, 0}
	tails := []Point{
		Point{0, 0},
		Point{0, 0},
		Point{0, 0},
		Point{0, 0},
		Point{0, 0},
		Point{0, 0},
		Point{0, 0},
		Point{0, 0},
		Point{0, 0},
	}
	for _, line := range lines {
		dir := string(line[0])
		amt, err := strconv.Atoi(line[2:])
		assertOk(err)
		//fmt.Println("==============", line)
		for x := 0; x < amt; x++ {
			switch dir {
			case "U":
				currH.y = currH.y + 1
			case "D":
				currH.y = currH.y - 1
			case "R":
				currH.x = currH.x + 1
			case "L":
				currH.x = currH.x - 1
			default:
				panic("bad dir:" + dir)
			}
			updateT(&currH, &tails[0])
			for t := 1; t < len(tails); t++ {
				updateT(&tails[t-1], &tails[t])
			}
			tVisited[tails[8]] = struct{}{}
			//printState(&currH, &tails)
			//fmt.Println()
		}
		//fmt.Println("===============")
	}
	/*
		for y := 20; y > -20; y-- {
			for x := -20; x < 20; x++ {
				if x == 0 && y == 0 {
					fmt.Printf("s")
				} else if _, ok := tVisited[Point{x, y}]; ok {
					fmt.Printf("#")
				} else {
					fmt.Printf(".")
				}
			}
			fmt.Println()
		}
	*/
	fmt.Println(len(tVisited))
}
