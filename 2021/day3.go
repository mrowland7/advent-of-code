package main

import (
	"log"
	//"regexp"
	"fmt"
	"strconv"
	//	"strings"
)

type bitCount struct {
	zeroes int
	ones   int
}

func (b bitCount) String() string {
	return fmt.Sprintf("(zeroes: %v, ones %v)", b.zeroes, b.ones)
}

func getBitCounts(lines []string) []*bitCount {
	bitCounts := make([]*bitCount, len(lines[0]))
	for i := 0; i < len(lines[0]); i++ {
		bitCounts[i] = &bitCount{zeroes: 0, ones: 0}
	}
	for _, line := range lines {
		for i, r := range line {
			cts := bitCounts[i]
			if r == '0' {
				cts.zeroes++
			} else {
				cts.ones++
			}
		}
	}
	return bitCounts
}

func main() {
	lines, err := getLines("day3_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	//part1(lines)
	ox := part2(lines, true)
	co2 := part2(lines, false)
	fmt.Printf("ox %v, co2 %v, %v\n", ox, co2, ox*co2)
}

func part2(lines []string, oxMode bool) int64 {
	remainingLines := lines
	iters := len(lines[0])
	for i := 0; i < iters; i++ {
		fmt.Printf("iter %v oxmode %v line left %v\n", i, oxMode, len(remainingLines))
		if len(remainingLines) == 1 {
			break
		}
		bitCounts := getBitCounts(remainingLines)
		var keepOnes bool
		if oxMode {
			keepOnes = bitCounts[i].ones >= bitCounts[i].zeroes
		} else {
			keepOnes = bitCounts[i].ones < bitCounts[i].zeroes
		}

		newRemainingLines := []string{}
		for _, line := range remainingLines {
			valAtIndex := line[i]
			if valAtIndex == '1' && keepOnes {
				newRemainingLines = append(newRemainingLines, line)
			} else if valAtIndex == '0' && !keepOnes {
				newRemainingLines = append(newRemainingLines, line)
			}
		}
		remainingLines = newRemainingLines
	}

	fmt.Printf("last line is %v (%v)\n", remainingLines[0], len(remainingLines))
	remaining, err := strconv.ParseInt(remainingLines[0], 2, 64)
	assertOk(err)
	return remaining
}

func part1(lines []string) {
	bitCounts := getBitCounts(lines)
	fmt.Printf("%+v\n", bitCounts)
	gammaStr := ""
	epsilonStr := ""
	for _, cts := range bitCounts {
		if cts.zeroes > cts.ones {
			gammaStr += "0"
			epsilonStr += "1"
		} else {
			gammaStr += "1"
			epsilonStr += "0"
		}
	}
	gamma, err := strconv.ParseInt(gammaStr, 2, 64)
	assertOk(err)
	eps, err := strconv.ParseInt(epsilonStr, 2, 64)
	assertOk(err)
	fmt.Printf("%v, %v, %v * %v = %v\n", gammaStr, epsilonStr, gamma, eps, gamma*eps)
}
