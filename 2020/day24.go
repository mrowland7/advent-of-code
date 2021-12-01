package main

import (
	"log"
	//"regexp"
	"fmt"
	//	"strconv"
	//	"strings"
)

// Grid is like this:
//
//     a  b  c
//    j nw ne k
//   l w  x  e  m
//    n sw se o
//     p  q  r
//
// Coordinates are:
// x -> (0, 0)
// w -> (0, -1) and e -> (0, 1)
// b -> (-2, 0) and q -> (2, 0)
// nw -> (-1, -0.5) and ne (-1, 0.5)

type coord struct {
	row float64
	col float64
}

func main() {
	lines, err := getLines("day24_input.txt")
	//lines, err := getLines("day24_dbg.txt")
	//lines, err := getLines("day24_dbg2.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Collect dirs
	dirs := [][]string{}
	for _, line := range lines {
		dir := []string{}
		for i := 0; i < len(line); i++ {
			if i+1 == len(line) {
				dir = append(dir, string(line[i]))
				continue
			}
			nextTwo := string(line[i : i+2])
			if len(nextTwo) != 2 {
				log.Fatal("ugh", nextTwo)
			}
			if nextTwo == "ne" || nextTwo == "nw" || nextTwo == "se" || nextTwo == "sw" {
				dir = append(dir, string(line[i:i+2]))
				i++
			} else {
				dir = append(dir, string(line[i]))
			}
		}
		dirs = append(dirs, dir)
	}
	// Move
	tiles := map[coord]bool{} // true == black side
	for _, dir := range dirs {
		r := 0.0
		c := 0.0
		for _, d := range dir {
			switch d {
			case "e":
				c += 1.0
			case "w":
				c -= 1.0
			case "ne":
				r -= 1.0
				c += 0.5
			case "nw":
				r -= 1.0
				c -= 0.5
			case "se":
				r += 1.0
				c += 0.5
			case "sw":
				r += 1.0
				c -= 0.5
			default:
				log.Fatal("unknown dir", d)
			}
		}
		prev, exists := tiles[coord{row: r, col: c}]
		if !exists {
			tiles[coord{row: r, col: c}] = true
		} else {
			tiles[coord{row: r, col: c}] = !prev
		}
	}
	for c, v := range tiles {
		if !v {
			delete(tiles, c)
		}
	}
	fmt.Println("starter count", len(tiles), "from total directions", len(dirs))
	// Now let's do some game of life
	deltas := []coord{
		{row: 0.0, col: 1.0},
		{row: 0.0, col: -1.0},
		{row: 1.0, col: 0.5},
		{row: -1.0, col: 0.5},
		{row: 1.0, col: -0.5},
		{row: -1.0, col: -0.5},
	}
	for i := 1; i <= 100; i++ {
		newTiles := map[coord]bool{}
		neighborCounts := map[coord]int{}
		for c, _ := range tiles {
			for _, d := range deltas {
				neighborCoord := coord{row: c.row + d.row, col: c.col + d.col}
				neighborCounts[neighborCoord] += 1
			}
		}
		// Flip any black tiles
		for c, _ := range tiles {
			neighborCt, _ := neighborCounts[c]
			if neighborCt == 1 || neighborCt == 2 {
				newTiles[c] = true
			}
		}
		// Flip any neighboring white tiles
		for c, ct := range neighborCounts {
			if ct == 2 {
				newTiles[c] = true
			}
		}
		tiles = newTiles
		fmt.Println("After day", i, "there are", len(tiles))
	}
}
