package main

import (
	"fmt"
	"regexp"
	"strconv"
	//	"strings"
)

//             [J] [Z] [G]
//             [Z] [T] [S] [P] [R]
// [R]         [Q] [V] [B] [G] [J]
// [W] [W]     [N] [L] [V] [W] [C]
// [F] [Q]     [T] [G] [C] [T] [T] [W]
// [H] [D] [W] [W] [H] [T] [R] [M] [B]
// [T] [G] [T] [R] [B] [P] [B] [G] [G]
// [S] [S] [B] [D] [F] [L] [Z] [N] [L]
//  1   2   3   4   5   6   7   8   9

func main() {
	lines, err := getLines("day5_input.txt")
	assertOk(err)
	stacks := [][]string{
		{"R", "W", "F", "H", "T", "S"},
		{"W", "Q", "D", "G", "S"},
		{"W", "T", "B"},
		{"J", "Z", "Q", "N", "T", "W", "R", "D"},
		{"Z", "T", "V", "L", "G", "H", "B", "F"},
		{"G", "S", "B", "V", "C", "T", "P", "L"},
		{"P", "G", "W", "T", "R", "B", "Z"},
		{"R", "J", "C", "T", "M", "G", "N"},
		{"W", "B", "G", "L"},
	}
	for _, stack := range stacks { // laziness
		s := stack
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}
	fmt.Println(stacks)
	r := regexp.MustCompile("[0-9]+")
	for _, line := range lines {
		vals := r.FindAll([]byte(line), 3)
		num, err := strconv.Atoi(string(vals[0]))
		assertOk(err)
		from, err := strconv.Atoi(string(vals[1]))
		assertOk(err)
		to, err := strconv.Atoi(string(vals[2]))
		assertOk(err)

		fmt.Println("=======", line)
		// Do the move
		fromStack := stacks[from-1]
		toStack := stacks[to-1]
		moving := fromStack[len(fromStack)-num : len(fromStack)]
		//for i, j := 0, len(moving)-1; i < j; i, j = i+1, j-1 {
		//	moving[i], moving[j] = moving[j], moving[i]
		//}
		fmt.Println("moving", moving, "from", fromStack, "to", toStack)
		stacks[from-1] = fromStack[0 : len(fromStack)-num]
		stacks[to-1] = append(toStack, moving...)
		fmt.Println("result", stacks[from-1], stacks[to-1])
	}
	fmt.Println(stacks)
	for _, stack := range stacks {
		fmt.Printf(stack[len(stack)-1])
	}
	fmt.Println()
}
