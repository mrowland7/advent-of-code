package main

import (
	//"log"
	//"regexp"
	"fmt"
	"strconv"
	//	"strings"
)

func part1(lines []string) {
	prevNum := int64(-1)
	increases := 0
	for i, line := range lines {
		num, err := strconv.ParseInt(line, 10, 64)
		assertOk(err)
		if i > 0 {
			if num > prevNum {
				increases++
			}
		}
		prevNum = num
	}
	fmt.Printf("increases: %v\n", increases)
}

func part2(lines []string) {
	windows := make([]int64, len(lines))
	nums := make([]int64, len(lines))
	for i, line := range lines {
		num, err := strconv.ParseInt(line, 10, 64)
		assertOk(err)
		nums[i] = num
	}

	for i, n := range nums {
		if i == 0 {
			windows[i] = n
		} else if i == 1 {
			windows[i] = n + nums[i-1]
		} else {
			windows[i] = n + nums[i-1] + nums[i-2]
		}
	}

	increases := 0
	for i, w := range windows {
		if i > 2 {
			if w > windows[i-1] {
				increases++
			}
		}
	}
	fmt.Printf("increases: %v\n", increases)
}

func main() {
	lines, err := getLines("day1_input.txt")
	assertOk(err)
	//part1(lines)
	part2(lines)
}
