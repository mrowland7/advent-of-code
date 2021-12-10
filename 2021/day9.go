package main

import (
	//"regexp"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Returns the number of tiles flowing into the basin centered at i, j.
func basinSearch(seafloor *[][]int, i, j int) int {
	total := 1
	(*seafloor)[i][j] = 9
	if (i-1) >= 0 && (*seafloor)[i-1][j] < 9 {
		total += basinSearch(seafloor, i-1, j)
	}
	if (i+1) < len(*seafloor) && (*seafloor)[i+1][j] < 9 {
		total += basinSearch(seafloor, i+1, j)
	}
	if (j-1) >= 0 && (*seafloor)[i][j-1] < 9 {
		total += basinSearch(seafloor, i, j-1)
	}
	if (j+1) < len((*seafloor)[0]) && (*seafloor)[i][j+1] < 9 {
		total += basinSearch(seafloor, i, j+1)
	}
	return total
}

func main() {
	lines, err := getLines("day9_input.txt")
	assertOk(err)
	seafloor := [][]int{}
	// get map
	for i, line := range lines {
		nums := strings.Split(line, "")
		seafloor = append(seafloor, make([]int, len(nums)))
		for j, num := range nums {
			n, err := strconv.Atoi(num)
			assertOk(err)
			seafloor[i][j] = n
		}
	}

	// get score
	riskSum := 0
	a := len(seafloor)
	b := len(seafloor[0])
	counts := []int{}
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			left := j == 0 || seafloor[i][j] < seafloor[i][j-1]
			right := j == b-1 || seafloor[i][j] < seafloor[i][j+1]
			up := i == 0 || seafloor[i][j] < seafloor[i-1][j]
			down := i == a-1 || seafloor[i][j] < seafloor[i+1][j]
			if left && right && up && down {
				riskSum += seafloor[i][j] + 1
				// Start basin count.
				counts = append(counts, basinSearch(&seafloor, i, j))
			}
		}
	}
	fmt.Println("score sum", riskSum)
	sort.Ints(counts)
	fmt.Println("counts", counts)
}
