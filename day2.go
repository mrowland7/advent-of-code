package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getLines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

type Rule struct {
	low    int
	high   int
	letter string
}

type PassRule struct {
	password string
	rule     Rule
}

func evaluatePart1(entries []PassRule) {
	validCt := 0
	for _, entry := range entries {
		ct := strings.Count(entry.password, entry.rule.letter)
		if ct >= entry.rule.low && ct <= entry.rule.high {
			fmt.Printf("%v is:\t valid\n, entry", entry)
			validCt++
		} else {
			fmt.Printf("%v is:\t invalid\n, entry", entry)
		}
	}
	fmt.Printf("validCt %v\n", validCt)
}

func evaluatePart2(entries []PassRule) {
	validCt := 0
	for _, entry := range entries {
		inPos1 := false
		inPos2 := false

		if len(entry.password) >= entry.rule.low {
			inPos1 = string(entry.password[entry.rule.low]) == entry.rule.letter
		}
		if len(entry.password) >= entry.rule.high {
			inPos2 = string(entry.password[entry.rule.high]) == entry.rule.letter
		}

		if (inPos1 && !inPos2) || (!inPos1 && inPos2) {
			fmt.Printf("%v is:\t valid\n, entry", entry)
			validCt++
		} else {
			fmt.Printf("%v is:\t invalid\n, entry", entry)
		}
	}
	fmt.Printf("validCt %v\n", validCt)
}

func main() {
	lines, err := getLines("day2_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("lines: %v\n", lines)

	// Parse
	entries := []PassRule{}
	for _, line := range lines {
		first := strings.Split(line, ":")
		if len(first) != 2 {
			log.Fatal("bad split: %v", line)
		}
		rawRule := first[0]
		pw := first[1]

		second := strings.Split(rawRule, " ")
		if len(second) != 2 {
			log.Fatal("bad split: %v", rawRule)
		}
		okRange := second[0]
		letter := second[1]

		third := strings.Split(okRange, "-")
		if len(third) != 2 {
			log.Fatal("bad split: %v", okRange)
		}
		min, err := strconv.Atoi(third[0])
		if err != nil {
			log.Fatal("couldn't parse int 1 in line: %v", line)
		}
		max, err := strconv.Atoi(third[1])
		if err != nil {
			log.Fatal("couldn't parse int 2 in line: %v", line)
		}
		entries = append(entries, PassRule{password: pw, rule: Rule{low: min, high: max, letter: letter}})
	}
	fmt.Printf("entries: %v\n", entries)

	// Evaluate
	evaluatePart2(entries)
}
