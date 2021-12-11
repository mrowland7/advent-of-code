package main

import (
	//"regexp"
	"fmt"
	//	"strconv"
	//	"strings"
)

func closeMatches(open, clos rune) bool {
	return (open == '(' && clos == ')') ||
		(open == '[' && clos == ']') ||
		(open == '{' && clos == '}') ||
		(open == '<' && clos == '>')
}

func main() {
	lines, err := getLines("day10_input.txt")
	assertOk(err)
	//	lines := []string{
	//		//	"{([(<{}[<>[]}>{[]{[(<()>",
	//		//	"[[<[([]))<([[{}[[()]]]",
	//		//	"[{[{({}]{}}([{[{{{}}([]",
	//		//	"[<(<(<(<{}))><([]([]()",
	//		"<{([([[(<>()){}]>(<<{{",
	//		"[({(<(())[]>[[{[]{<()<>>",
	//		"[(()[<>])]({[<{<<[]>>(",
	//		"(((({<>}<{<{<>}{[]{[]{}",
	//		"{<[[]]>}<{[{[{[]{()[[[]",
	//		"<{([{{}}[<[[[<>{}]]]>[]]",
	//	}
	errorScoreMap := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	autoScoreMap := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
	errorScore := 0
	//autoScore := 0
	for _, line := range lines {
		stack := []rune{}
		corrupt := false
		for _, r := range line {
			if r == '(' || r == '[' || r == '{' || r == '<' {
				stack = append(stack, r)
			} else if closeMatches(stack[len(stack)-1], r) {
				stack = stack[:len(stack)-1]
			} else {
				errorScore += errorScoreMap[r]
				corrupt = true
				break
			}
		}
		if corrupt {
			continue
		}
		thisScore := 0
		for i := len(stack) - 1; i >= 0; i-- {
			thisScore = thisScore*5 + autoScoreMap[stack[i]]
		}
		fmt.Println(thisScore)
		//autoScore += thisScore
	}
	//fmt.Println("final scores", errorScore, autoScore)
	// Then:
	// $ go run helpers.go day10.go | sort -n | head -n 24
}
