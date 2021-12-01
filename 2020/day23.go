package main

import (
	"container/ring"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func cpy(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}

func printRing(r *ring.Ring, num int) {
	fmt.Printf("ring: ")
	r2 := r
	for i := 0; i < num; i++ {
		fmt.Printf("%v-", r2.Value)
		r2 = r2.Next()
	}
	fmt.Println()
}

func main() {
	input := "871369452"
	numCups := 1000000
	numRounds := 10000000
	//numCups := len(input) + 20
	//	numRounds := 100
	r := ring.New(numCups)
	cache := map[int]*ring.Ring{}
	for _, s := range strings.Split(input, "") {
		n, err := strconv.Atoi(s)
		assertOk(err)
		r.Value = n
		cache[n] = r
		r = r.Next()
	}
	if numCups > len(input) {
		for i := len(input) + 1; i < numCups+1; i++ {
			r.Value = i
			cache[i] = r
			r = r.Next()
		}
	}
	fmt.Println("starting")
	for round := 1; round <= numRounds; round++ {
		if round%10000 == 0 {
			fmt.Println("round", round, "percent", (round*100.0)/(numRounds*1.0), "time", time.Now())
		}
		subring := r.Unlink(3)
		destination := r.Value.(int)
		for true {
			destination--
			if destination == 0 {
				destination = numCups
			}
			if subring.Value.(int) != destination &&
				subring.Next().Value.(int) != destination &&
				subring.Next().Next().Value.(int) != destination {
				break
			}
		}
		toLinkBack := cache[destination]
		toLinkBack.Link(subring)
		r = r.Next()
	}

	if numCups < 100 {
		fmt.Println("final is...")
		printRing(r, numCups)
	}
	for i := 0; i < numCups; i++ {
		if r.Value.(int) == 1 {
			fmt.Println("next is", r.Next().Value.(int), "next next is", r.Next().Next().Value.(int))
			break
		}
		r = r.Next()
	}
}

// This way was too slow
func main1() {
	input := "871369452"
	//input := "389125467"
	numCups := 1000000
	numRounds := 10000000
	//numCups := len(input)
	//numRounds := 100
	cups := make([]int, numCups)
	for i, s := range strings.Split(input, "") {
		n, err := strconv.Atoi(s)
		assertOk(err)
		cups[i] = n
	}
	for i := len(input); i < numCups; i++ {
		cups[i] = i + 1
	}
	currentCupIndex := 0
	for round := 1; round <= numRounds; round++ {
		if round%10000 == 0 {
			fmt.Println("round", round, "percent", (round*100.0)/(numRounds*1.0), "time", time.Now())
		}
		//fmt.Println("===== cups", cups, "current cup", cups[currentCupIndex], "at index", currentCupIndex)
		currentCupVal := cups[currentCupIndex]
		// Find cups to pull
		nextThree := []int{}
		indexesPulled := []int{}
		for i := currentCupIndex + 1; i <= currentCupIndex+3; i++ {
			idx := i % len(cups)
			nextThree = append(nextThree, cups[idx])
			indexesPulled = append(indexesPulled, idx)
		}
		// Find destination
		destination := cups[currentCupIndex]
		for true {
			destination--
			if destination == 0 {
				destination = numCups
			}
			if nextThree[0] != destination && nextThree[1] != destination && nextThree[2] != destination {
				break
			}
		}
		// Pull them
		rest := make([]int, len(cups)-3)
		skips := 0
		targetIdx := -1
		for i := 0; i < len(cups); i++ {
			if i == indexesPulled[0] ||
				i == indexesPulled[1] ||
				i == indexesPulled[2] {
				skips++
				continue
			}
			rest[i-skips] = cups[i]
			if rest[i-skips] == destination {
				targetIdx = i - skips
			}
		}
		// Put them back in
		newCups := make([]int, numCups)
		copy(newCups[0:targetIdx+1], rest[0:targetIdx+1])
		copy(newCups[targetIdx+1:targetIdx+4], nextThree)
		copy(newCups[targetIdx+4:], rest[targetIdx+1:])

		// Find the next index
		for i := 0; i < len(newCups); i++ {
			if newCups[i] == currentCupVal {
				currentCupIndex = (i + 1) % len(newCups)
				break
			}
		}
		cups = newCups
	}
	// find 1
	//fmt.Println("cups", cups)
	for i := 0; i < len(cups); i++ {
		if cups[i] == 1 {
			fmt.Println("next two are", cups[i+1], "and", cups[i+2], "result", cups[i+1]*cups[i+2])
			break
		}
	}

}
