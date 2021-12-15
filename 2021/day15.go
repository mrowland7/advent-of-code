package main

import (
	//"regexp"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type loc struct {
	i int
	j int
}

type node struct {
	enterCost         int
	tentativePathCost int
	visited           bool
}

func main() {
	lines, err := getLines("day15.txt")
	assertOk(err)
	tileLen := 100
	//	lines := []string{
	//		"1163751742",
	//		"1381373672",
	//		"2136511328",
	//		"3694931569",
	//		"7463417111",
	//		"1319128137",
	//		"1359912421",
	//		"3125421639",
	//		"1293138521",
	//		"2311944581",
	//	}
	//	sideLen := 10
	sideLen := tileLen * 5
	caves := make([][]node, sideLen)
	for i, line := range lines {
		caves[i] = make([]node, sideLen)
		caves[i+tileLen] = make([]node, sideLen)
		caves[i+2*tileLen] = make([]node, sideLen)
		caves[i+3*tileLen] = make([]node, sideLen)
		caves[i+4*tileLen] = make([]node, sideLen)
		vals := strings.Split(line, "")
		for j, v := range vals {
			n, err := strconv.Atoi(v)
			assertOk(err)
			for id := 0; id < 5; id++ {
				for jd := 0; jd < 5; jd++ {
					delta := id + jd
					enterCost := n + delta
					if enterCost > 9 {
						enterCost = enterCost - 9
					}
					caves[i+id*tileLen][j+jd*tileLen] = node{enterCost: enterCost, tentativePathCost: math.MaxInt32, visited: false}
				}
			}
		}
	}
	fmt.Println("read in caves")
	// Build out the lowest cost. Dijkstra's
	curr := loc{i: 0, j: 0}
	caves[curr.i][curr.j].tentativePathCost = 0
	for true {
		neighbors := []loc{
			//loc{curr.i - 1, curr.j - 1},
			loc{curr.i - 1, curr.j},
			//loc{curr.i - 1, curr.j + 1},
			loc{curr.i, curr.j - 1},
			loc{curr.i, curr.j + 1},
			//loc{curr.i + 1, curr.j - 1},
			loc{curr.i + 1, curr.j},
			//loc{curr.i + 1, curr.j + 1},
		}
		for _, neighbor := range neighbors {
			if !(neighbor.i >= 0 && neighbor.i < sideLen && neighbor.j >= 0 && neighbor.j < sideLen) {
				continue
			}
			costFromHere := caves[curr.i][curr.j].tentativePathCost + caves[neighbor.i][neighbor.j].enterCost
			if costFromHere < caves[neighbor.i][neighbor.j].tentativePathCost {
				caves[neighbor.i][neighbor.j].tentativePathCost = costFromHere
			}
		}
		caves[curr.i][curr.j].visited = true
		fmt.Println("visited", curr, "lowest cost", caves[curr.i][curr.j].tentativePathCost)
		if curr.i == sideLen-1 && curr.j == sideLen-1 {
			fmt.Println("done, cost", caves[curr.i][curr.j].tentativePathCost)
			break
		}
		// Pick next curr
		lowest := math.MaxInt32
		argmin := loc{i: -1, j: -1}
		for i := 0; i < sideLen; i++ {
			for j := 0; j < sideLen; j++ {
				if caves[i][j].tentativePathCost < lowest && !caves[i][j].visited {
					lowest = caves[i][j].tentativePathCost
					argmin = loc{i: i, j: j}
				}
			}
		}
		curr = argmin
	}
}
