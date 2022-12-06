package main

import (
	//"regexp"
	"fmt"
	//	"strconv"
	//	"strings"
)

func part1(lines []string) {
	score := 0
	// X, Y, Z = rock, paper, scissors
	for _, line := range lines {
		them := line[0]
		me := line[2]
		switch them {
		case 'A':
			switch me {
			case 'X':
				score += 1 + 3
			case 'Y':
				score += 2 + 6
			case 'Z':
				score += 3 + 0
			default:
				panic("bad line:" + line)
			}
		case 'B':
			switch me {
			case 'X':
				score += 1 + 0
			case 'Y':
				score += 2 + 3
			case 'Z':
				score += 3 + 6
			default:
				panic("bad line:" + line)
			}
		case 'C':
			switch me {
			case 'X':
				score += 1 + 6
			case 'Y':
				score += 2 + 0
			case 'Z':
				score += 3 + 3
			default:
				panic("bad line:" + line)
			}
		default:
			panic("bad line:" + line)
		}

	}
	fmt.Println(score)
}

func part2(lines []string) {
	score := 0
	// X, Y, Z = lose, draw, win
	for _, line := range lines {
		them := line[0]
		me := line[2]
		switch them {
		case 'A':
			switch me {
			case 'X':
				score += 3 + 0
			case 'Y':
				score += 1 + 3
			case 'Z':
				score += 2 + 6
			default:
				panic("bad line:" + line)
			}
		case 'B':
			switch me {
			case 'X':
				score += 1 + 0
			case 'Y':
				score += 2 + 3
			case 'Z':
				score += 3 + 6
			default:
				panic("bad line:" + line)
			}
		case 'C':
			switch me {
			case 'X':
				score += 2 + 0
			case 'Y':
				score += 3 + 3
			case 'Z':
				score += 1 + 6
			default:
				panic("bad line:" + line)
			}
		default:
			panic("bad line:" + line)
		}

	}
	fmt.Println(score)
}

func main() {
	lines, err := getLines("day2_input.txt")
	assertOk(err)
	part1(lines)
	part2(lines)
}
