package main

import (
	//"regexp"
	"fmt"
	"sort"
	"strconv"
	//	"strings"
)

func main() {
	lines, err := getLines("day1_input.txt")
	assertOk(err)
	elves := []int{}
	elf := 0
	for _, line := range lines {
		if line == "" {
			elves = append(elves, elf)
			elf = 0
			continue
		}
		x, err := strconv.Atoi(line)
		assertOk(err)
		elf += x
	}
	sort.Ints(elves)
	fmt.Println(elves[len(elves)-1] + elves[len(elves)-2] + elves[len(elves)-3])
}
