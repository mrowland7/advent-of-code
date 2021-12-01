package main

import (
	"log"
	//"regexp"
	"fmt"
	"sort"
	"strconv"
	//	"strings"
)

func total(adapters []int) {
	jump1s := 0
	jump3s := 0
	for i, n := range adapters {
		fmt.Println("adapter: ", n)
		if i == 0 {
			continue
		}
		diff := n - adapters[i-1]
		if diff == 1 {
			jump1s++
		}
		if diff == 3 {
			jump3s++
		}
		if diff > 3 {
			break
		}
	}
	jump3s++ // to final device
	fmt.Println(jump1s, jump3s, jump1s*jump3s)
}

func combinations(adapters []int) {
	ways := make([]int64, len(adapters))
	for i, n := range adapters {
		if i == 0 {
			ways[i] = 1
			continue
		}
		sub1 := int64(0)
		sub2 := int64(0)
		sub3 := int64(0)
		if i-1 >= 0 && ways[i-1] >= 0 && n-adapters[i-1] <= 3 {
			sub1 = ways[i-1]
		}
		if i-2 >= 0 && ways[i-2] >= 0 && n-adapters[i-2] <= 3 {
			sub2 = ways[i-2]
		}
		if i-3 >= 0 && ways[i-3] >= 0 && n-adapters[i-3] <= 3 {
			sub3 = ways[i-3]
		}
		ways[i] = sub1 + sub2 + sub3

		fmt.Println("adapter:", n, "total ways", ways[i], "(", sub1, sub2, sub3, ")")
	}
}

func main() {
	lines, err := getLines("day10_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	adapters := []int{0}
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		adapters = append(adapters, n)
	}
	sort.Ints(adapters)
	//total(adapters)
	combinations(adapters)
}
