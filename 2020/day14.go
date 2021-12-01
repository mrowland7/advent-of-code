package main

import (
	"log"
	//"regexp"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func applyBitmask(value int, mask map[int]int) int {
	for i, val := range mask {
		set := (value & (1 << i)) > 0
		//fmt.Println("value:", value, "bit", i, "set?", set, ", set to", val)
		if set && val == 0 {
			value ^= 1 << i
		} else if !set && val == 1 {
			value ^= 1 << i
		}
	}
	return value
}

func getAddresses(address int, mask map[int]int) []int {
	result := []int{}
	// First make base number and collect wildcards
	wildcardIndexes := []int{}
	for i, val := range mask {
		if val == -1 {
			wildcardIndexes = append(wildcardIndexes, i)
		} else {
			address |= val << i
		}
	}
	numWildcards := len(wildcardIndexes)
	maxWildcardNum := 1 << numWildcards
	if numWildcards == 0 {
		return []int{address}
	}
	sort.Ints(wildcardIndexes)
	//fmt.Println("wildcards are", wildcardIndexes)
	for w := 0; w < maxWildcardNum; w++ {
		mask := map[int]int{}
		for i := 0; i < len(wildcardIndexes); i++ {
			// 6th wildcard number w = 6
			// 0110
			// applied to some number like
			// __X___X__X_X
			// then we want mask[original index] = ith bit of w
			set := w&(1<<i) > 0
			if set {
				mask[wildcardIndexes[i]] = 1
			} else {
				mask[wildcardIndexes[i]] = 0
			}
		}
		newAddr := applyBitmask(address, mask)
		//fmt.Println("addr", address, "mask: ", mask, "newaddr", newAddr)
		result = append(result, newAddr)
	}

	return result
}

func main() {
	lines, err := getLines("day14_input.txt")
	//lines, err := getLines("day14_dbg.txt")
	if err != nil {
		log.Fatal(err)
	}

	memory := map[int]int{}
	currentMask := map[int]int{}
	for _, line := range lines {
		if line[0:4] == "mask" {
			parts := strings.Split(line, " = ")
			currentMask = map[int]int{}
			for i, bit := range parts[1] {
				if bit == rune('0') {
					currentMask[35-i] = 0
				} else if bit == rune('1') {
					currentMask[35-i] = 1
				} else {
					currentMask[35-i] = -1
				}
			}
			//fmt.Println("mask now", currentMask)
			continue
		}
		parts := strings.Split(line, " = ")
		value, err := strconv.Atoi(parts[1])
		assertOk(err)
		parts2 := strings.Split(parts[0], "[")
		addr, err := strconv.Atoi(strings.TrimRight(parts2[1], "]"))
		assertOk(err)
		for _, a := range getAddresses(addr, currentMask) {
			fmt.Println("writing address", a)
			memory[a] = value
		}
		fmt.Println("=====")
		// part 1
		//		memory[addr] = applyBitmask(value, currentMask)
	}

	sum := int64(0)
	for addr, m := range memory {
		sum += int64(m)
		fmt.Println("addr:", addr, "value:", m)
	}
	fmt.Println("sum", sum)

}
