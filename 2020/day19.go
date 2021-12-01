package main

import (
	"log"
	//"regexp"
	"fmt"
	"strconv"
	"strings"
)

type Pattern struct {
	id int

	// IDs of sub-patterns. a b | c d e
	a int
	b int
	c int
	d int
	e int

	// value to directly match
	value string
}

type MemoPair struct {
	id int
	s  string
}

var memo = map[MemoPair]bool{}

func matches(patterns map[int]*Pattern, id int, s string) bool {
	prev, exists := memo[MemoPair{id, s}]
	if exists {
		return prev
	}
	p := patterns[id]
	if p.value != "" {
		match := p.value == s
		memo[MemoPair{id, s}] = match
		return match
	}
	if len(s) == 0 {
		memo[MemoPair{id, s}] = false
		return false
	}
	// If either side is just another pattern, recur that way.
	if p.a >= 0 && p.b == -1 {
		leftMatches := matches(patterns, p.a, s)
		if leftMatches {
			memo[MemoPair{id, s}] = true
			return true
		}
	}
	if p.c >= 0 && p.d == -1 {
		rightMatches := matches(patterns, p.c, s)
		if rightMatches {
			memo[MemoPair{id, s}] = true
			return true
		}
	}
	// Otherwise, check all the splits of the string.
	if len(s) >= 2 && (p.b >= 0 || p.d >= 0) {
		for i := 1; i < len(s); i++ {
			leftS := s[:i]
			rightS := s[i:]
			if p.b >= 0 {
				leftMatch := matches(patterns, p.a, leftS)
				rightMatch := matches(patterns, p.b, rightS)
				if leftMatch && rightMatch {
					memo[MemoPair{id, s}] = true
					return true
				}
			}
			if p.d >= 0 && p.e == -1 {
				leftMatch := matches(patterns, p.c, leftS)
				rightMatch := matches(patterns, p.d, rightS)
				if leftMatch && rightMatch {
					memo[MemoPair{id, s}] = true
					return true
				}
			}
		}
	}
	// Special case: have to check two splits
	if len(s) >= 3 && p.e >= 0 {
		for i := 1; i < len(s)-1; i++ {
			for j := i + 1; j < len(s); j++ {
				leftS := s[0:i]
				middleS := s[i:j]
				rightS := s[j:]
				leftMatch := matches(patterns, p.c, leftS)
				midMatch := matches(patterns, p.d, middleS)
				rightMatch := matches(patterns, p.e, rightS)
				if leftMatch && midMatch && rightMatch {
					memo[MemoPair{id, s}] = true
					return true
				}
			}
		}
	}
	memo[MemoPair{id, s}] = false
	return false
}

func main() {
	lines, err := getLines("day19_input.txt")
	//lines, err := getLines("day19_dbg.txt")
	if err != nil {
		log.Fatal(err)
	}
	block := 1
	patterns := map[int]*Pattern{}
	strs := []string{}
	for _, line := range lines {
		if line == "" {
			block++
			continue
		}
		if block == 1 {
			a, b, c, d, e := -1, -1, -1, -1, -1
			value := ""
			parts1 := strings.Split(line, ": ")
			id, err := strconv.Atoi(parts1[0])
			assertOk(err)
			if parts1[1] == "\"a\"" {
				value = "a"
			} else if parts1[1] == "\"b\"" {
				value = "b"
			} else {
				parts2 := strings.Split(parts1[1], " | ")
				if len(parts2) == 2 {
					parts4 := strings.Split(parts2[1], " ")
					c, err = strconv.Atoi(parts4[0])
					assertOk(err)
					if len(parts4) == 2 {
						d, err = strconv.Atoi(parts4[1])
						assertOk(err)
					}
				}
				parts3 := strings.Split(parts2[0], " ")
				a, err = strconv.Atoi(parts3[0])
				assertOk(err)
				if len(parts3) == 2 {
					b, err = strconv.Atoi(parts3[1])
					assertOk(err)
				}
			}
			p := &Pattern{
				id:    id,
				a:     a,
				b:     b,
				c:     c,
				d:     d,
				e:     e,
				value: value,
			}
			patterns[id] = p
		} else {
			strs = append(strs, line)
		}
	}
	// Special ones
	patterns[8] = &Pattern{
		id:    8,
		a:     42,
		b:     -1,
		c:     42,
		d:     8,
		e:     -1,
		value: "",
	}
	patterns[11] = &Pattern{
		id:    11,
		a:     42,
		b:     31,
		c:     42,
		d:     11,
		e:     31,
		value: "",
	}

	numMatches := 0
	for _, s := range strs {
		fmt.Println("Checking if", s, "matches: ...")
		res := matches(patterns, 0, s)
		fmt.Println("==", res)
		if res {
			numMatches++
		}
	}

	fmt.Println("matches", numMatches)
	fmt.Println("memo size", len(memo))
}
