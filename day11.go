package main

import (
	"log"
	//"regexp"
	"fmt"
	//	"strconv"
	"strings"
)

func getSurroundingSeats1(h, w int, seatMap [][]string) []string {
	res := []string{}
	for hidx := h - 1; hidx <= h+1; hidx++ {
		if hidx < 0 || hidx >= len(seatMap) {
			continue
		}
		for widx := w - 1; widx <= w+1; widx++ {
			if widx < 0 || widx >= len(seatMap[0]) {
				continue
			}
			if widx == w && hidx == h {
				continue
			}
			res = append(res, seatMap[hidx][widx])
		}
	}
	return res
}

func getSurroundingSeats2(h, w int, seatMap [][]string) []string {
	res := []string{}
	dirs := []struct {
		hdiff int
		wdiff int
	}{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for _, dir := range dirs {
		hidx := h
		widx := w
		for true {
			hidx += dir.hdiff
			widx += dir.wdiff
			if hidx < 0 || hidx >= len(seatMap) || widx < 0 || widx >= len(seatMap[0]) {
				break
			}
			s := seatMap[hidx][widx]
			if s != "." {
				res = append(res, s)
				break
			}
			// XXX for part 1
			//break
		}
	}
	return res
}

func main() {
	lines, err := getLines("day11_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	width := len(lines[0])
	height := len(lines)
	seatMap := make([][]string, height)
	for i, line := range lines {
		seatMap[i] = strings.Split(line, "")
		if len(seatMap[i]) != width {
			log.Fatal("bad split", i, line, seatMap[i])
		}
		fmt.Println(line)
	}
	changed := true
	for changed {
		changed = false
		nextSeatMap := make([][]string, height)
		for h, row := range seatMap {
			nextSeatMap[h] = make([]string, len(row))
			for w, seat := range row {
				otherSeats := getSurroundingSeats2(h, w, seatMap)
				occCt := 0
				for _, os := range otherSeats {
					if os == "#" {
						occCt++
					}
				}
				switch seat {
				case "L":
					{
						if occCt == 0 {
							nextSeatMap[h][w] = "#"
							changed = true
						} else {
							nextSeatMap[h][w] = "L"
						}
					}
				case "#":
					{
						if occCt >= 5 {
							nextSeatMap[h][w] = "L"
							changed = true
						} else {
							nextSeatMap[h][w] = "#"
						}
					}
				case ".":
					{
						nextSeatMap[h][w] = "."
					}
				default:
					log.Fatal("unknown letter", seat)
				}
			}
		}
		seatMap = nextSeatMap
	}
	// Count
	finalOccCt := 0
	for _, row := range seatMap {
		for _, seat := range row {
			if seat == "#" {
				finalOccCt++
			}
		}
	}
	fmt.Println(finalOccCt)
}
