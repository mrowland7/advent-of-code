package main

import (
	//"regexp"
	"fmt"
	"strconv"
	//	"strings"
)

func main() {
	lines, err := getLines("day8_input.txt")
	//	lines, err := getLines("day8_dbg.txt")
	assertOk(err)
	forest := make([][]int, len(lines))
	for i, line := range lines {
		forest[i] = make([]int, len(line))
		for j, c := range line {
			x, err := strconv.Atoi(string(c))
			assertOk(err)
			forest[i][j] = x
		}
	}
	vizCount := (len(forest) - 1) * 4
	maxScore := -1
	for i := 1; i < len(forest)-1; i++ {
		for j := 1; j < len(forest[0])-1; j++ {
			visUp, visDown, visLeft, visRight := true, true, true, true
			distUp, distDown, distLeft, distRight := i, len(forest)-i-1, j, len(forest[0])-j-1
			for i2 := i - 1; i2 >= 0; i2-- {
				if forest[i2][j] >= forest[i][j] {
					visUp = false
					distUp = i - i2
					break
				}
			}
			for i2 := i + 1; i2 < len(forest); i2++ {
				if forest[i2][j] >= forest[i][j] {
					visDown = false
					distDown = i2 - i
					break
				}
			}
			for j2 := j - 1; j2 >= 0; j2-- {
				if forest[i][j2] >= forest[i][j] {
					visLeft = false
					distLeft = j - j2
					break
				}
			}
			for j2 := j + 1; j2 < len(forest[0]); j2++ {
				if forest[i][j2] >= forest[i][j] {
					visRight = false
					distRight = j2 - j
					break
				}
			}

			if visUp || visDown || visLeft || visRight {
				vizCount++
			}
			score := distUp * distDown * distLeft * distRight
			//	fmt.Println("Checking", i, j, "=", score, "with breakdown of", distUp, distDown, distLeft, distRight)
			//	fmt.Println("=== Vis", visUp, visDown, visLeft, visRight)
			if score > maxScore {
				maxScore = score
			}
		}
	}
	fmt.Println(vizCount)
	fmt.Println(maxScore)
}
