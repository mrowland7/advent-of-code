package main

import (
	"log"
	//"regexp"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines, err := getLines("day2_input.txt")
	assertOk(err)
	depth := 0
	horiz := 0
	aim := 0
	for _, line := range lines {
		vals := strings.Split(line, " ")
		if len(vals) != 2 {
			log.Fatal("bad input: ", vals)
		}
		dir, numStr := vals[0], vals[1]
		num, err := strconv.Atoi(numStr)
		assertOk(err)
		if dir == "forward" {
			horiz += num
			depth += aim * num
		} else if dir == "up" {
			//depth -= num
			aim -= num
		} else if dir == "down" {
			//depth += num
			aim += num
		}
	}
	fmt.Printf("horiz %v depth %v product %v\n", horiz, depth, horiz*depth)
}
