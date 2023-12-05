package main

import (
	"log"
	//"regexp"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	lines, err := getLines("day4_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	points := 0.0
	copies := map[int]int{}
	for i, line := range lines {
		cardPieces := strings.Split(line, ": ")
		copies[i+1] += 1
		nums := strings.Split(cardPieces[1], " | ")
		winningNums := strings.Split(nums[0], " ")
		cardNums := strings.Split(nums[1], " ")
		winners := []int{}
		for _, cardNum := range cardNums {
			for _, winningNum := range winningNums {
				if cardNum == winningNum {
					num, err := strconv.Atoi(cardNum)
					assertOk(err)
					winners = append(winners, num)
				}
			}
		}
		for step := 1; step <= len(winners); step++ {
			copies[i+1+step] += copies[i+1]
		}
		if len(winners) > 0 {
			points += math.Pow(2, float64(len(winners)-1))
		}

		fmt.Println(line, "--->", winners)
	}
	fmt.Println("part 1 points:", points)
	totalCopies := 0
	for _, cpy := range copies {
		totalCopies += cpy
	}
	fmt.Println("part 2 copies:", totalCopies)
}
