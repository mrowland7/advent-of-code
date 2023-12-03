package main

import (
	"log"
	//"regexp"
	"fmt"
	"strconv"
	"strings"
)

type game struct {
	red   int
	green int
	blue  int
}

func main() {
	lines, err := getLines("day2_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	maxGame := &game{12, 13, 14}
	sumIDs := 0
	sumPowers := 0
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		gameNumPieces := strings.Split(parts[0], " ")
		gameNum, err := strconv.Atoi(gameNumPieces[1])
		assertOk(err)
		rounds := strings.Split(parts[1], "; ")
		validGame := true
		minSet := &game{0, 0, 0}
		for _, round := range rounds {
			roundTotals := &game{}
			cubes := strings.Split(round, ", ")
			for _, cube := range cubes {
				vals := strings.Split(cube, " ")
				amt, err := strconv.Atoi(vals[0])
				assertOk(err)
				color := vals[1]
				if color == "red" {
					roundTotals.red = roundTotals.red + amt
				} else if color == "blue" {
					roundTotals.blue = roundTotals.blue + amt
				} else if color == "green" {
					roundTotals.green = roundTotals.green + amt
				}
			}
			if roundTotals.red > maxGame.red ||
				roundTotals.blue > maxGame.blue ||
				roundTotals.green > maxGame.green {
				validGame = false
			}
			if roundTotals.red > minSet.red {
				minSet.red = roundTotals.red
			}
			if roundTotals.blue > minSet.blue {
				minSet.blue = roundTotals.blue
			}
			if roundTotals.green > minSet.green {
				minSet.green = roundTotals.green
			}
		}
		minSetPower := minSet.red * minSet.green * minSet.blue
		fmt.Println(line, "--> Game", gameNum, ",", validGame, ",", minSet, minSetPower)
		if validGame {
			sumIDs += gameNum
		}
		sumPowers += minSetPower
	}
	fmt.Println(sumIDs)
	fmt.Println(sumPowers)
}
