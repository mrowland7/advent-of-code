package main

import (
	"log"
	//"regexp"
	"fmt"
	//	"strconv"
	"strings"
)

func main() {
	lines, err := getLines("day21_input.txt")
	//lines, err := getLines("day21_dbg.txt")

	if err != nil {
		log.Fatal(err)
	}
	// Map from allergen -> set of possible candidates
	candidates := map[string]map[string]bool{}
	ingredientCount := map[string]int{}
	for _, line := range lines {
		parts := strings.Split(line, " (")
		allergens := strings.Split(strings.Trim(parts[1], ")")[9:], ", ")
		ingredients := strings.Split(parts[0], " ")
		for _, a := range allergens {
			prev, ok := candidates[a]
			// If it's the first time seeing the allergen, add all the ingredients as candidates
			if !ok {
				candidates[a] = map[string]bool{}
				for _, ing := range ingredients {
					candidates[a][ing] = true
				}
			} else {
				// Otherwise, eliminate any previous contenders that are NOT in here.
				for ing1, v := range prev {
					if !v {
						continue
					}
					foundAgain := false
					for _, ing2 := range ingredients {
						if ing1 == ing2 {
							foundAgain = true
						}
					}
					if !foundAgain {
						prev[ing1] = false
					}
				}
			}
		}
		// Stat keeping
		for _, ing := range ingredients {
			if _, ok := ingredientCount[ing]; !ok {
				ingredientCount[ing] = 0
			}
			ingredientCount[ing]++
		}
	}
	totalCt := 0
	for ing, count := range ingredientCount {
		stillCandidate := false
		for _, candidates := range candidates {
			if candidates[ing] {
				stillCandidate = true
				break
			}
		}
		if !stillCandidate {
			totalCt += count
		}
	}
	fmt.Println("sum is", totalCt)

	for allergen, potentials := range candidates {
		for p, v := range potentials {
			if !v {
				delete(potentials, p)
			}
		}
		fmt.Println(allergen, "->", potentials)
	}

}
