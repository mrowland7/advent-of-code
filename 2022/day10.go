package main

import (
	//"regexp"
	"fmt"
	"strconv"
	//	"strings"
)

type Instruction struct {
	op     string
	num    int
	cycles int
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func main() {
	lines, err := getLines("day10_input.txt")
	//	lines, err := getLines("day10_dbg1.txt")
	assertOk(err)
	ixns := []*Instruction{}
	for _, line := range lines {
		var cycles int
		switch line[0:4] {
		case "addx":
			cycles = 2
		case "noop":
			cycles = 1
		default:
			cycles = 0
		}
		ixn := &Instruction{
			op:     line[0:4],
			cycles: cycles,
		}
		if len(line) > 4 {
			num, err := strconv.Atoi(line[5:])
			assertOk(err)
			ixn.num = num
		}
		ixns = append(ixns, ixn)
	}
	cycle := 0
	registerX := 1
	score := 0
	for _, ixn := range ixns {
		for c := 0; c < ixn.cycles; c++ {
			if cycle%40 == 0 {
				fmt.Println()
			}
			if (cycle-20)%40 == 0 {
				score += cycle * registerX
				//fmt.Println("Adding", cycle, "*", registerX)
			}
			if abs(registerX-(cycle%40)) <= 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			cycle++
		}
		switch ixn.op {
		case "addx":
			registerX += ixn.num
		}
	}
	fmt.Println(score)
}
