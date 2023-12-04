package main

import (
	"log"
	//"regexp"
	"fmt"
	"strconv"
	//	"strings"
)

func isDigit(r byte) bool {
	_, err := strconv.Atoi(string(r))
	return err == nil
}

func isSymbol(r byte) bool {
	return !isDigit(r) && r != '.'
}

type loc struct {
	y int
	x int
}

type pair struct {
	a int
	b int
}

func main() {
	lines, err := getLines("day3_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sumPartNumbers := 0
	gearMap := map[loc]pair{}
	for i, line := range lines {
		fmt.Println(line)
		for j := 0; j < len(line); j++ {
			r := line[j]
			if !isDigit(r) {
				continue
			}
			// Get the number
			start := j
			for isDigit(r) {
				j++
				if j >= len(line) {
					break
				}
				r = line[j]
			}
			num, err := strconv.Atoi(line[start:j])
			assertOk(err)
			fmt.Println("found number", num)
			// Search around for a symbol
			foundSymbol := false
			for y := i - 1; y <= i+1; y++ {
				for x := start - 1; x <= j; x++ {
					if y < 0 || y >= len(lines) || x < 0 || x >= len(line) {
						continue
					}
					if isSymbol(lines[y][x]) {
						foundSymbol = true
						if lines[y][x] == '*' {
							vals := gearMap[loc{y: y, x: x}]
							if vals.a > 0 {
								vals.b = num
							} else {
								vals.a = num
							}
							gearMap[loc{y: y, x: x}] = vals
						}
					}
				}
			}
			if foundSymbol {
				sumPartNumbers += num
			}
		}
	}
	gearSum := 0
	for _, vals := range gearMap {
		gearSum += vals.a * vals.b
	}
	fmt.Println(sumPartNumbers)
	fmt.Println(gearSum)
}
