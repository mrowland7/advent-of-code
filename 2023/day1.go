package main

import (
	"log"
	//"regexp"
	"fmt"
	//	"strconv"
	//	"strings"
)

func main() {
	lines, err := getLines("day1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	total := 0
	for _, line := range lines {
		first := -1
		last := -1
		for i, ch := range line {
			val := -1
			valFound := false
			// Check digits
			if ch == '1' ||
				ch == '2' ||
				ch == '3' ||
				ch == '4' ||
				ch == '5' ||
				ch == '6' ||
				ch == '7' ||
				ch == '8' ||
				ch == '9' ||
				ch == '0' {
				if first == -1 {
					val = int(ch - '0')
				}
				val = int(ch - '0')
				valFound = true
			}
			if i+3 <= len(line) {
				// Check word style
				w3 := line[i : i+3]
				if w3 == "one" {
					val = 1
					valFound = true
				} else if w3 == "two" {
					val = 2
					valFound = true
				} else if w3 == "six" {
					val = 6
					valFound = true
				}
			}
			if i+4 <= len(line) {
				// Check word style
				w4 := line[i : i+4]
				if w4 == "four" {
					val = 4
					valFound = true
				} else if w4 == "five" {
					val = 5
					valFound = true
				} else if w4 == "nine" {
					val = 9
					valFound = true
				}
			}
			if i+5 <= len(line) {
				// Check word style
				w5 := line[i : i+5]
				if w5 == "three" {
					val = 3
					valFound = true
				} else if w5 == "seven" {
					val = 7
					valFound = true
				} else if w5 == "eight" {
					val = 8
					valFound = true
				}
			}

			if valFound {
				if first == -1 {
					first = val
				}
				last = val
			}
		}
		if first == -1 {
			continue
		}
		lineVal := first*10 + last
		total += lineVal
		fmt.Println(line, first, last, lineVal, total)
	}
}
