package main

import (
	//"regexp"
	"fmt"
	"unicode"
	//	"strconv"
	//	"strings"
)

func score(c rune) int {
	var score int
	if unicode.IsUpper(c) {
		score = int(c) - int('A') + 27
	} else {
		score = int(c) - int('a') + 1
	}
	return score
}

func part1(lines []string) {
	prioritySum := 0
	for _, line := range lines {
		c1 := line[0 : len(line)/2]
		c2 := line[len(line)/2 : len(line)]
		set1 := map[rune]struct{}{}
		set2 := map[rune]struct{}{}
		for _, c := range c1 {
			set1[c] = struct{}{}
		}
		for _, c := range c2 {
			set2[c] = struct{}{}
		}
		for c, _ := range set1 {
			if _, ok := set2[c]; ok {
				prioritySum += score(c)
			}
		}
	}
	fmt.Println(prioritySum)
}
func part2(lines []string) {
	prioritySum := 0
	for i := 0; i < len(lines); i += 3 {
		r1 := lines[i]
		r2 := lines[i+1]
		r3 := lines[i+2]
		set1 := map[rune]struct{}{}
		set2 := map[rune]struct{}{}
		set3 := map[rune]struct{}{}
		for _, c := range r1 {
			set1[c] = struct{}{}
		}
		for _, c := range r2 {
			set2[c] = struct{}{}
		}
		for _, c := range r3 {
			set3[c] = struct{}{}
		}
		for c, _ := range set1 {
			if _, ok := set2[c]; ok {
				if _, ok := set3[c]; ok {
					prioritySum += score(c)
				}
			}
		}
	}
	fmt.Println(prioritySum)
}

func main() {
	lines, err := getLines("day3_input.txt")
	assertOk(err)
	part1(lines)
	part2(lines)
}
