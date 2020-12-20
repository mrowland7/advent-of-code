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
// edges are numbered like:
//   __0__
//  |     |
// 3|  t2 |1
//  |_____|
//     2
//
//   __0__
//  |     |
// 3|  t1 |1
//  |_____|
//     2
//

type orientation struct {
	degrees int
	flipped bool
}

type maptile struct {
	id int
	or orientation
	t  *Tile
}

func seamonsterSearch(tiles []*Tile, edgeMap map[string][]int) {
	maps := getTiledMaps(tiles, edgeMap)
	//maps = [][][]maptile{maps[0]}
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
	fmt.Println()
	if false {
		return
	}
	textTileLength := len(tiles[0].edges[0]) - 2
	textSideLength := len(maps[0][0]) * textTileLength

	// Ok, now we have the maps. Time to search the seas!!
	for _, tileMap := range maps {
		lines := make([][]string, textSideLength)
		for i := 0; i < len(tileMap); i++ {
			tileRow := tileMap[i]
			lineRowStart := i * textTileLength
			for j := 0; j < len(tileRow); j++ {
				//fmt.Println("=== Placing tile", i, j)
				tileIdOr := tileRow[j]
				tile := indexedTiles[tileIdOr.id]
				// Apply rotation to lines
				//Flip first
				textLines := make([][]string, len(tile.raw))
				for i := 0; i < len(tile.raw); i++ {
					if tileIdOr.or.flipped {
						textLines[i] = strings.Split(reverse(tile.raw[i]), "")
					} else {
						textLines[i] = strings.Split(tile.raw[i], "")
					}
				}
				// Then rotate
				rot := tileIdOr.or.degrees
				for rot > 0 {
					//fmt.Println("rot is", rot)
					newLines := make([][]string, len(textLines))
					for i := 0; i < len(textLines); i++ {
						newLines[i] = make([]string, len(textLines[i]))
					}
					n := len(textLines)
					for i := 0; i < len(textLines); i++ {
						for j := 0; j < len(textLines[i]); j++ {
							newLines[i][j] = textLines[n-j-1][i]
						}
					}
					textLines = newLines
					rot -= 90
				}

				for r := 1; r < len(textLines)-1; r++ {
					line := textLines[r]
					lines[lineRowStart+r-1] = append(lines[lineRowStart+r-1], line[1:len(line)-1]...)
				}
			}
		}
		for _, l := range lines {
			fmt.Println(strings.Join(l, ""))
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
		fmt.Println("num monsters", seaMonsterCount, "sea roughness count:", seaCount)
	}
}

func getTiledMaps(tiles []*Tile, edgeMap map[string][]int) [][][]maptile {
	indexedTiles := map[int]*Tile{}
	for _, t := range tiles {
		indexedTiles[t.id] = t
	}
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
	var maps [][][]maptile
	for cornerIdx := 0; cornerIdx < 1; cornerIdx++ {
		for flipIdx := 1; flipIdx <= 1; flipIdx++ {
			flipped := false
			if flipIdx == 1 {
				flipped = true
			}
			fmt.Println("\n===== corner", cornerIdx, "flip", flipped)
			placed := map[int]bool{}
			mapSize := int(math.Sqrt(float64(len(idsToNeighbors))))
			tileMap := make([][]maptile, mapSize)
			for i := 0; i < mapSize; i++ {
				tileMap[i] = make([]maptile, mapSize)
				for j := 0; j < mapSize; j++ {
					fmt.Println("- placing", i, j)
					if i == 0 && j == 0 {
						id := corners[cornerIdx]
						orient := getInitialOrientation(edgeMap, indexedTiles[id], flipped)
						tileMap[i][j] = maptile{
							id: id,
							or: orient,
							t:  indexedTiles[id],
						}
						placed[corners[cornerIdx]] = true
						continue
					}
					// Look at the neighbors of the tile to the left and up.
					neighbors := map[int]bool{}
					if i-1 < 0 {
						for n, _ := range idsToNeighbors[tileMap[i][j-1].id] {
							neighbors[n] = true
						}
					} else if j-1 < 0 {
						for n, _ := range idsToNeighbors[tileMap[i-1][j].id] {
							neighbors[n] = true
						}
					} else {
						leftNeighbors := map[int]bool{}
						upNeighbors := map[int]bool{}
						for n, _ := range idsToNeighbors[tileMap[i][j-1].id] {
							leftNeighbors[n] = true
						}
						for n, _ := range idsToNeighbors[tileMap[i-1][j].id] {
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
							// Ok, we're POTENTIALLY placing this file here.
							var upMaptile *maptile = nil
							var leftMaptile *maptile = nil
							if i-1 >= 0 {
								upMaptile = &tileMap[i-1][j]
							}
							if j-1 >= 0 {
								leftMaptile = &tileMap[i][j-1]
							}
							orientation, err := getOrientation(indexedTiles[n], upMaptile, leftMaptile)
							if err != nil {
								//Ok, no way to put it there.
								fmt.Println("- no match for tile", n)
								continue
							}
							fmt.Println("- match for tile", n, ":", orientation)
							tileMap[i][j] = maptile{
								id: n,
								or: *orientation,
								t:  indexedTiles[n],
							}
							placed[n] = true
							break
						}
					}
				}
			}
			maps = append(maps, tileMap)
		}
	}
	return maps
}

func getInitialOrientation(edgeMap map[string][]int, tile *Tile, flipped bool) orientation {
	// "Flips" will be considered as left-to-right
	degrees := -1
	// First find the two edges that face in.
	topCt := len(edgeMap[tile.edges[0]])
	rightCt := len(edgeMap[tile.edges[1]])
	bottomCt := len(edgeMap[tile.edges[2]])
	leftCt := len(edgeMap[tile.edges[3]])
	if !flipped {
		if topCt == 2 && rightCt == 2 {
			degrees = 90
		} else if rightCt == 2 && bottomCt == 2 {
			degrees = 0
		} else if bottomCt == 2 && leftCt == 2 {
			degrees = 270
		} else if leftCt == 2 && topCt == 2 {
			degrees = 180
		}
	} else {
		if topCt == 2 && rightCt == 2 {
			degrees = 180
		} else if rightCt == 2 && bottomCt == 2 {
			degrees = 270
		} else if bottomCt == 2 && leftCt == 2 {
			degrees = 0
		} else if leftCt == 2 && topCt == 2 {
			degrees = 90
		}
	}
	if degrees == -1 {
		log.Fatal("ugh")
	}

	return orientation{
		degrees: degrees,
		flipped: flipped,
	}
}

// Gets the facing edge
func getEdge(startIdx int, mt *maptile) string {
	idx := startIdx
	turns := mt.or.degrees / 90
	if !mt.or.flipped {
		idx -= turns
	} else {
		idx += turns
	}
	if idx < 0 {
		idx = idx + 4
	}
	idx = idx % 4
	if mt.or.flipped {
		return reverse(mt.t.edges[idx])
	} else {
		return mt.t.edges[idx]
	}
}

func getOrientation(tile *Tile, up *maptile, left *maptile) (*orientation, error) {
	//fmt.Println("--  Getting orientation for tile", tile.id)
	var exposedUp, exposedLeft string
	if up != nil {
		exposedUp = getEdge(2, up)
	}
	if left != nil {
		exposedLeft = getEdge(3, left)
	}

	//fmt.Println("--  up is", exposedUp, "left is", exposedLeft)
	for i, e := range tile.edges {
		//fmt.Println("--  edge", i, "is", e)
		// Normal match: you have to REVERSE the one edge, since they're store in clockwise order
		if reverse(e) == exposedUp || up == nil {
			// Possible... check if the others match
			leftEdgeIdx := (i - 1 + 4) % 4
			if reverse(tile.edges[leftEdgeIdx]) == exposedLeft || left == nil {
				//	fmt.Println("--  match 1, i=", i, ", leftEdgeIdx", leftEdgeIdx)
				return &orientation{
					degrees: i * 90,
					flipped: false,
				}, nil
			}
		}
		// Check reverse match
		if e == exposedUp || up == nil {
			leftEdgeIdx := (i + 1 + 4) % 4
			if tile.edges[leftEdgeIdx] == exposedLeft || left == nil {
				//fmt.Println("--  match 2, i=", i, ", leftEdgeIdx", leftEdgeIdx)
				return &orientation{
					degrees: i * 90,
					flipped: true,
				}, nil
			}
		}
	}

	return nil, fmt.Errorf("no match")
}

func main() {
	lines, err := getLines("day20_dbg.txt")
	//	lines, err := getLines("day20_input.txt")
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
					reverse(collected_lines[len(collected_lines)-1]),
					reverse(edge3),
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
