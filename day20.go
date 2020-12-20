package main

import (
	"log"
	//"regexp"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Tile struct {
	id    int
	edges []string
	raw   []string
}

// Finds the corners
func cornerSearch(tiles []*Tile, edgeMap map[string][]int) []int {
	cornerProd := 1
	corners := []int{}
	for _, t := range tiles {
		soloCount := 0
		for _, edge := range t.edges {
			if len(edgeMap[edge]) == 1 {
				soloCount++
			}
		}
		if soloCount == 2 {
			fmt.Println("Found a corner: ", t.id)
			corners = append(corners, t.id)
			cornerProd *= t.id
		}
	}
	fmt.Print("product is ", cornerProd)
	return corners
}

// returns the rotations for t1 where there's a match for t2
func findTileMatches(t1, t2 *Tile, baseRot int) map[int]bool {
	if t2 == nil {
		return map[int]bool{0: true, 90: true, 180: true, 270: true}
	}
	res := map[int]bool{}
	for i := 0; i < 4; i++ {
		t1e := t1.edges[i]
		for _, t2e := range t2.edges {
			if t1e == t2e || t1e == reverse(t2e) {
				res[(90*i+baseRot)%360] = true
				break
			}
		}
	}
	return res
}

func seamonsterSearch(tiles []*Tile, edgeMap map[string][]int) {
	maps := getTiledMaps(tiles, edgeMap)
	for cornerIdx, tileMap := range maps {
		fmt.Printf("\nmap %v is...", cornerIdx)
		for _, row := range tileMap {
			fmt.Println()
			for _, id := range row {
				fmt.Printf(" %v ", id)
			}
		}
	}
	indexedTiles := map[int]*Tile{}
	for _, t := range tiles {
		indexedTiles[t.id] = t
	}
	// Ok, now we have the maps. Time to search the seas!!
	for _, tileMap := range maps {
		sideLength := len(tileMap)*(len(tiles[0].edges[0])-1) + 2
		lines := make([][]string, sideLength)
		for i, _ := range lines {
			lines[i] = make([]string, sideLength)
		}
		for i := 0; i < len(tileMap); i++ {
			tileRow := tileMap[0]
			for j := 0; j < len(tileRow); j++ {
				tileId := tileRow[j]
				tile := indexedTiles[tileId]
				var top, bottom, left, right *Tile
				if i-1 >= 0 {
					top = indexedTiles[tileMap[i-1][j]]
				}
				if i+1 < len(tileMap) {
					bottom = indexedTiles[tileMap[i+1][j]]
				}
				if j-1 >= 0 {
					left = indexedTiles[tileMap[i][j-1]]
				}
				if j+1 < len(tileMap) {
					right = indexedTiles[tileMap[i][j+1]]
				}
				topMatches := findTileMatches(tile, top, 0)
				bottomMatches := findTileMatches(tile, bottom, 180)
				leftMatches := findTileMatches(tile, left, 90)
				rightMatches := findTileMatches(tile, right, 270)
				rot := 0
				matchRot := -1
				for rot < 360 {
					_, ok1 := topMatches[rot]
					_, ok2 := bottomMatches[rot]
					_, ok3 := leftMatches[rot]
					_, ok4 := rightMatches[rot]
					if ok1 && ok2 && ok3 && ok4 {
						if matchRot > -1 {
							log.Fatal("two matches!!")
						}
						matchRot = rot
					}
					rot += 90
				}
				rot = matchRot

				// Add to lines
				textLines := tile.raw
				if rot == 0 {
					// Top matches top tile, Right edge matches right tile, etc
				} else if rot == 90 {
					// Top edge matches right tile
				} else if rot == 180 {
					// Left edge (reversed) matches right tile
				} else if rot == 270 {
					// Bottom edge (reversed) matches right tile
				}
				for _, row := range textLines {
					linesI := i * len(row)
				}
			}
		}
		// Ok, we have the lines!! Now mark the sea monsters with O, then count the #
		seaMonsterCount := 0
		for i := 0; i < len(lines)-2; i++ {
			for j := 18; j < len(lines[i])-1; j++ {
				type spot struct {
					iv int
					jv int
				}
				spots := []spot{
					{i, j},
					{i + 1, j - 1},
					{i + 1, j},
					{i + 1, j + 1},
					{i + 2, j - 2},
					{i + 2, j - 5},
					{i + 2, j - 8},
					{i + 2, j - 11},
					{i + 2, j - 14},
					{i + 2, j - 17},
					{i + 1, j - 6},
					{i + 1, j - 7},
					{i + 1, j - 12},
					{i + 1, j - 13},
					{i + 1, j - 18},
				}
				allMatch := true
				for _, spot := range spots {
					if lines[spot.iv][spot.jv] != "#" {
						allMatch = false
						break
					}
				}
				if allMatch {
					seaMonsterCount++
					for _, spot := range spots {
						lines[spot.iv][spot.jv] = "O"
					}
				}

			}
		}
		seaCount := 0
		for _, line := range lines {
			for _, sq := range line {
				if sq == "#" {
					seaCount++
				}
			}
		}
		fmt.Println()
		fmt.Println("num monsters", seaMonsterCount, "sea roughness count:", seaCount)
	}
}

func getTiledMaps(tiles []*Tile, edgeMap map[string][]int) [][][]int {
	type edgePair struct {
		id1 int
		id2 int
	}
	// Set of all pairs
	pairs := map[edgePair]bool{}
	// Map id -> set of neighbors
	idsToNeighbors := map[int]map[int]bool{}
	for _, ids := range edgeMap {
		if len(ids) != 2 {
			continue
		}
		var ep edgePair
		if ids[0] < ids[1] {
			ep = edgePair{id1: ids[0], id2: ids[1]}
		} else {
			ep = edgePair{id1: ids[1], id2: ids[0]}
		}
		pairs[ep] = true
		if _, ok := idsToNeighbors[ep.id1]; !ok {
			idsToNeighbors[ep.id1] = map[int]bool{}
		}
		if _, ok := idsToNeighbors[ep.id2]; !ok {
			idsToNeighbors[ep.id2] = map[int]bool{}
		}
		idsToNeighbors[ep.id1][ep.id2] = true
		idsToNeighbors[ep.id2][ep.id1] = true
	}

	corners := cornerSearch(tiles, edgeMap)

	// 4 tile maps, with each corner at the top-left
	maps := make([][][]int, 4)
	for cornerIdx := 0; cornerIdx < 4; cornerIdx++ {

		placed := map[int]bool{}
		mapSize := int(math.Sqrt(float64(len(idsToNeighbors))))
		tileMap := make([][]int, mapSize)
		for i := 0; i < mapSize; i++ {
			tileMap[i] = make([]int, mapSize)
			for j := 0; j < mapSize; j++ {
				if i == 0 && j == 0 {
					tileMap[i][j] = corners[cornerIdx]
					placed[corners[cornerIdx]] = true
					continue
				}
				// Look at the neighbors of the tile to the left.
				neighbors := map[int]bool{}
				if i-1 < 0 {
					for n, _ := range idsToNeighbors[tileMap[i][j-1]] {
						neighbors[n] = true
					}
				} else if j-1 < 0 {
					for n, _ := range idsToNeighbors[tileMap[i-1][j]] {
						neighbors[n] = true
					}
				} else {
					leftNeighbors := map[int]bool{}
					upNeighbors := map[int]bool{}
					for n, _ := range idsToNeighbors[tileMap[i][j-1]] {
						leftNeighbors[n] = true
					}
					for n, _ := range idsToNeighbors[tileMap[i-1][j]] {
						upNeighbors[n] = true
					}
					neighbors = leftNeighbors
					// Get the set intersection (somewhat unwieldy)
					for n, _ := range upNeighbors {
						if _, ok := neighbors[n]; !ok {
							delete(neighbors, n)
						}
					}
					for n, _ := range leftNeighbors {
						if _, ok := upNeighbors[n]; !ok {
							delete(neighbors, n)
						}
					}
				}
				// Remove anything that's already been placed
				for p, _ := range placed {
					delete(neighbors, p)
				}
				// Figure out how many neighbors this tile should have!
				desiredNeighbors := 4
				// Horizontal edge
				if i == 0 || i == (mapSize-1) {
					desiredNeighbors--
				}
				if j == 0 || j == (mapSize-1) {
					desiredNeighbors--
				}
				for n, _ := range neighbors {
					numNeighbors := len(idsToNeighbors[n])
					if numNeighbors == desiredNeighbors {
						tileMap[i][j] = n
						placed[n] = true
						break
					}
				}
			}
		}
		maps[cornerIdx] = tileMap
	}
	return maps
}

func main() {
	lines, err := getLines("day20_dbg.txt")
	//lines, err := getLines("day20_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	currentTileId := 0
	tiles := []*Tile{}
	collected_lines := []string{}
	for _, line := range lines {
		if line == "" && currentTileId > 0 {
			edge3, edge4 := "", ""
			for _, l := range collected_lines {
				edge3 = edge3 + string(l[0])
				edge4 = edge4 + string(l[len(l)-1])
			}
			t := &Tile{
				id: currentTileId,
				edges: []string{
					collected_lines[0],
					edge4,
					collected_lines[len(collected_lines)-1],
					edge3,
				},
				raw: collected_lines,
			}
			tiles = append(tiles, t)

			collected_lines = []string{}
			continue
		}
		if line[0] == 'T' {
			currentTileId, err = strconv.Atoi(strings.Trim(line, "Tile :"))
			assertOk(err)
			continue
		}
		collected_lines = append(collected_lines, line)
	}

	edgeMap := map[string][]int{}
	for _, t := range tiles {
		for _, edge := range t.edges {
			e1 := edge
			e2 := reverse(e1)
			m1, ok := edgeMap[e1]
			if !ok {
				edgeMap[e1] = []int{t.id}
			} else {
				edgeMap[e1] = append(m1, t.id)
			}
			m2, ok := edgeMap[e2]
			if !ok {
				edgeMap[e2] = []int{t.id}
			} else {
				edgeMap[e2] = append(m2, t.id)
			}
		}
	}
	//cornerSearch(tiles, edgeMap)
	seamonsterSearch(tiles, edgeMap)
}
