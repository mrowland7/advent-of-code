package main

import (
	"log"
	//"regexp"
	"fmt"
	"strconv"
	"strings"
)

type BoardCell struct {
	num    string
	marked bool
}

func (b BoardCell) String() string {
	return fmt.Sprintf("[%v %v]", b.num, b.marked)
}

func winner(board []*BoardCell) bool {
	offsets := [][]int{
		{0, 1, 2, 3, 4},
		{5, 6, 7, 8, 9},
		{10, 11, 12, 13, 14},
		{15, 16, 17, 18, 19},
		{20, 21, 22, 23, 24},
		{0, 5, 10, 15, 20},
		{1, 6, 11, 16, 21},
		{2, 7, 12, 17, 22},
		{3, 8, 13, 18, 23},
		{4, 9, 14, 19, 24},
	}

	for _, o := range offsets {
		result := true
		for _, o2 := range o {
			result = result && board[o2].marked
		}
		if result {
			return true
		}
	}

	return false
}

func mark(num string, board []*BoardCell) {
	for _, cell := range board {
		if cell.num == num {
			cell.marked = true
			break
		}
	}
}

func main() {
	lines, err := getLines("day4_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	nums := strings.Split(lines[0], ",")
	boards := [][]*BoardCell{}
	nextBoard := []*BoardCell{}
	for _, line := range lines[2:] {
		if line == "" {
			boards = append(boards, nextBoard)
			nextBoard = []*BoardCell{}
			continue
		}
		lineNums := strings.Split(line, " ")
		for _, lineNum := range lineNums {
			if lineNum == "" {
				continue
			}
			nextBoard = append(nextBoard, &BoardCell{
				num:    lineNum,
				marked: false,
			})
		}
	}

	// Get winner
	winningBoard := -1
	lastNum := -1
	alreadyWon := map[int]struct{}{}
	winners := 0
	for _, num := range nums {
		fmt.Printf("calling %v, %v %v\n", num, winners, len(boards))
		if winners == len(boards) {
			break
		}
		for i := 0; i < len(boards); i++ {
			if _, ok := alreadyWon[i]; ok {
				continue
			}
			mark(num, boards[i])
			if winner(boards[i]) {
				alreadyWon[i] = struct{}{}
				winners++
				fmt.Printf("board %v wins, num won %v, %v\n", i, len(alreadyWon), len(boards))
				if winners == len(boards) {
					lastNum, err = strconv.Atoi(num)
					assertOk(err)
					winningBoard = i
				}
				//fmt.Printf("winner! lastnum %v, index %v, %v\n", lastNum, i, boards[i])
				//winningBoard = i
			}
		}
		//if winningBoard >= 0 {
		//	break
		//}
	}

	// Calculate score for winner
	s := 0
	for _, cell := range boards[winningBoard] {
		if !cell.marked {
			n, err := strconv.Atoi(cell.num)
			assertOk(err)
			s += n
		}
	}

	fmt.Printf("result %v*%v=%v\n", s, lastNum, s*lastNum)
}
