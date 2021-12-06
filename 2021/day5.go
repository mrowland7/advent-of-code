package main

import (
	"log"
	//"regexp"
	"fmt"
	"strconv"
	"strings"
)

type Coord struct {
	x int
	y int
}

type Line struct {
	start *Coord
	end   *Coord
}

func main() {
	lines, err := getLines("day5_input.txt")
	assertOk(err)
	lineSegments := []*Line{}
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		if len(parts) != 2 {
			log.Fatal("bad parts")
		}
		start := parts[0]
		end := parts[1]
		sp := strings.Split(start, ",")
		ep := strings.Split(end, ",")
		x1, err := strconv.Atoi(sp[0])
		assertOk(err)
		y1, err := strconv.Atoi(sp[1])
		assertOk(err)
		x2, err := strconv.Atoi(ep[0])
		assertOk(err)
		y2, err := strconv.Atoi(ep[1])
		assertOk(err)
		lineSegments = append(lineSegments, &Line{
			start: &Coord{x: x1, y: y1},
			end:   &Coord{x: x2, y: y2},
		})
	}
	fmt.Println("parsed")
	overlapCt := map[Coord]int{}
	for _, line := range lineSegments {
		if line.start.x == line.end.x {
			if line.start.y > line.end.y {
				tmp := line.end
				line.end = line.start
				line.start = tmp
			}
			for i := line.start.y; i <= line.end.y; i++ {
				c := Coord{x: line.start.x, y: i}
				val, ok := overlapCt[c]
				if !ok {
					val = 0
				}
				overlapCt[c] = val + 1
			}
		} else if line.start.y == line.end.y {
			if line.start.x > line.end.x {
				tmp := line.end
				line.end = line.start
				line.start = tmp
			}
			for i := line.start.x; i <= line.end.x; i++ {
				c := Coord{x: i, y: line.start.y}
				val, ok := overlapCt[c]
				if !ok {
					val = 0
				}
				overlapCt[c] = val + 1
			}
		} else {
			// 45 degrees. still put lower x first
			if line.start.x > line.end.x {
				tmp := line.end
				line.end = line.start
				line.start = tmp
			}
			yStep := 1
			if line.end.y < line.start.y {
				yStep = -1
			}
			for iter := 0; iter <= (line.end.x - line.start.x); iter++ {
				c := Coord{x: line.start.x + iter, y: line.start.y + iter*yStep}
				val, ok := overlapCt[c]
				if !ok {
					val = 0
				}
				overlapCt[c] = val + 1
			}
		}
	}

	numGt1 := 0
	for _, val := range overlapCt {
		if val > 1 {
			numGt1++
		}
	}
	fmt.Printf("%v\n", numGt1)
}
