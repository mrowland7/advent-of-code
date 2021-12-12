package main

import (
	//"regexp"
	"fmt"
	"strconv"
	//	"strings"
)

func main() {
	lines := []string{
		"6744638455",
		"3135745418",
		"4754123271",
		"4224257161",
		"8167186546",
		"2268577674",
		"7177768175",
		"2662255275",
		"4655343376",
		"7852526168",
	}
	tiles := [][]int{}
	for i, line := range lines {
		tiles = append(tiles, make([]int, len(line)))
		for j, r := range line {
			n, err := strconv.Atoi(string(r))
			assertOk(err)
			tiles[i][j] = n
		}
	}
	totalFlashes := 0
	for iter := 0; iter < 10000; iter++ {
		fmt.Println("iter", iter, ", flashes so far", totalFlashes)
		flashesThisTurn := 0
		// Increase
		for i := 0; i < len(tiles); i++ {
			for j := 0; j < len(tiles[0]); j++ {
				tiles[i][j]++
			}
		}
		// Begin flashing process
		checkFlashes := true
		for checkFlashes {
			checkFlashes = false
			for i := 0; i < len(tiles); i++ {
				for j := 0; j < len(tiles[0]); j++ {
					if tiles[i][j] <= 9 {
						continue
					}
					// flash!
					checkFlashes = true
					tiles[i][j] = -1
					totalFlashes++
					flashesThisTurn++
					for di := -1; di <= 1; di++ {
						for dj := -1; dj <= 1; dj++ {
							if di == 0 && dj == 0 {
								continue
							}
							if i+di >= 0 && i+di < len(tiles) && j+dj >= 0 && j+dj < len(tiles[0]) &&
								tiles[i+di][j+dj] >= 0 {
								tiles[i+di][j+dj]++
							}
						}
					}
				}
			}
		}
		// Reset all the -1s to 0.
		for i := 0; i < len(tiles); i++ {
			for j := 0; j < len(tiles[0]); j++ {
				if tiles[i][j] == -1 {
					tiles[i][j] = 0
				}
			}
		}
		if flashesThisTurn == len(tiles)*len(tiles[0]) {
			fmt.Println("Done at step", iter+1)
			break
		}
	}
	fmt.Println("total", totalFlashes)
}
