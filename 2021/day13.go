package main

import (
	//"regexp"
	"fmt"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

func main() {
	lines, err := getLines("day13_input.txt")
	//lines, err := getLines("day13_dbg.txt")
	assertOk(err)
	folds := []string{
		"x=655",
		"y=447",
		"x=327",
		"y=223",
		"x=163",
		"y=111",
		"x=81",
		"y=55",
		"x=40",
		"y=27",
		"y=13",
		"y=6",
		//	"y=7",
		//	"x=5",
	}
	dots := map[coord]struct{}{}
	for _, line := range lines {
		parts := strings.Split(line, ",")
		x, err := strconv.Atoi(parts[0])
		assertOk(err)
		y, err := strconv.Atoi(parts[1])
		assertOk(err)
		dots[coord{x: x, y: y}] = struct{}{}
	}
	fmt.Println("number of dots to start:", len(dots))
	//	printDots(dots)
	for i, fold := range folds {
		//for dot, _ := range dots {
		//	fmt.Println("dot", dot.x, dot.y)
		//}
		newDots := map[coord]struct{}{}
		parts := strings.Split(fold, "=")
		axis := parts[0]
		foldLine, err := strconv.Atoi(parts[1])
		assertOk(err)
		for dot, _ := range dots {
			if (axis == "x" && dot.x < foldLine) ||
				(axis == "y" && dot.y < foldLine) {
				newDots[dot] = struct{}{}
			} else if axis == "x" {
				delta := dot.x - foldLine
				newDots[coord{x: foldLine - delta, y: dot.y}] = struct{}{}
			} else {
				delta := dot.y - foldLine
				newDots[coord{x: dot.x, y: foldLine - delta}] = struct{}{}
			}
		}
		dots = newDots
		fmt.Println("number of dots after fold", i+1, len(dots))
		printDots(dots)
	}
}

func printDots(dots map[coord]struct{}) {
	for i := 0; i < 15; i++ {
		for j := 0; j < 50; j++ {
			if _, ok := dots[coord{x: j, y: i}]; ok {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}
