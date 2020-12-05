package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	//"strconv"
	//"strings"
)

func getLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

type Seat struct {
	row   int
	col   int
	raw   string
	score int
}

func main() {
	lines, err := getLines("day5_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	seats := []*Seat{}
	for _, line := range lines {
		if len(line) != 10 {
			log.Fatal("bad line", line)
		}
		row := line[:7]
		seat := line[7:]

		ri := 6.0
		rowVal := 0.0
		for _, c := range row {
			if c == 'B' {
				rowVal += math.Pow(2, ri)
			}
			ri--
		}
		ci := 2.0
		colVal := 0.0
		for _, c := range seat {
			if c == 'R' {
				colVal += math.Pow(2, ci)
			}
			ci--
		}
		score := rowVal*8 + colVal
		s := &Seat{row: int(rowVal), col: int(colVal), raw: line, score: int(score)}
		seats = append(seats, s)
	}
	sort.Slice(seats, func(i, j int) bool {
		return seats[i].score < seats[j].score
	})

	maxSeatScore := -1
	minSeatScore := 10000
	maxSeat := &Seat{}
	minSeat := &Seat{}
	prevSeatID := 90
	for _, s := range seats {
		if s.score-1 != prevSeatID {
			fmt.Println("=========== missing!")
		}
		prevSeatID = s.score
		fmt.Println(s)
		if s.score > maxSeatScore {
			maxSeatScore = s.score
			maxSeat = s
		}
		if s.score < minSeatScore {
			minSeatScore = s.score
			minSeat = s
		}
	}
	fmt.Println("max seat", maxSeat, "min seat", minSeat)
}
