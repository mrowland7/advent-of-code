package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

func anyone(lines []string) {
	collected_line := ""
	totalCt := 0
	for _, line := range lines {
		fmt.Println(line)
		if len(line) > 0 {
			collected_line = collected_line + " " + line
			continue
		}
		// Have everything in one line; parse
		answers := strings.Fields(collected_line)
		allAns := map[rune]bool{}
		for _, answer := range answers {
			for _, c := range answer {
				allAns[c] = true
			}
		}
		fmt.Println(len(allAns), collected_line)
		totalCt += len(allAns)
		fmt.Println("\n==========\n")

		// Reset
		collected_line = ""
	}
	fmt.Println("total", totalCt)
}

func everyone(lines []string) {
	collected_line := ""
	totalCt := 0
	for _, line := range lines {
		if len(line) > 0 {
			sp := strings.Split(line, "")
			sort.Strings(sp)
			fmt.Println(len(line), strings.Join(sp, ""), line)
			collected_line = collected_line + " " + line
			continue
		}
		// Have everything in one line; parse
		answers := strings.Fields(collected_line)
		allAns := map[string]bool{}
		for _, c := range "abcdefghijklmnopqrstuvwxyz" {
			allPresent := true
			for _, ans := range answers {
				if !strings.ContainsRune(ans, c) {
					allPresent = false
				}
			}
			if allPresent {
				allAns[string(c)] = true
			}
		}

		fmt.Println(len(allAns), allAns)
		totalCt += len(allAns)
		fmt.Println("running total", totalCt, "\n==========\n")

		// Reset
		collected_line = ""
	}
	fmt.Println("total", totalCt)
}

func main() {
	lines, err := getLines("day6_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// anyone(lines)
	everyone(lines)
}
