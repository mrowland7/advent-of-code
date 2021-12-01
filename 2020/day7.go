package main

import (
	"log"
	//"regexp"
	"fmt"
	"strconv"
	"strings"
)

type BagHold struct {
	name  string
	count int
}

type Bag struct {
	name              string
	contains          []BagHold
	visited           bool
	containsShinyGold bool
}

func canContainShinyGold(b *Bag, bagMap map[string]*Bag) bool {
	if b.visited {
		return b.containsShinyGold
	}
	b.visited = true
	for _, b2 := range b.contains {
		if b2.name == "shiny gold" {
			b.containsShinyGold = true
			return true
		}
		if canContainShinyGold(bagMap[b2.name], bagMap) {
			b.containsShinyGold = true
			return true
		}
	}
	b.containsShinyGold = false
	return false
}

func nestedBagCt(b *Bag, bagMap map[string]*Bag) int {
	sum := 0
	for _, b2 := range b.contains {
		sum += b2.count * (1 + nestedBagCt(bagMap[b2.name], bagMap))
	}
	return sum
}

func main() {
	lines, err := getLines("day7_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Get bag rules
	bags := map[string]*Bag{}
	for _, line := range lines {
		line := line[:len(line)-1]
		sp1 := strings.Split(line, " bags contain ")
		// Parse
		if len(sp1) != 2 {
			log.Fatal("bad line sp1", line, sp1)
		}
		name := sp1[0]
		rawRule := sp1[1]
		otherBags := []BagHold{}
		if rawRule != "no other bags" {
			sp2 := strings.Split(rawRule, ", ")
			for _, r := range sp2 {
				sp3 := strings.Split(r, " ")
				if len(sp3) != 4 {
					log.Fatal("bad line sp3:", line, ";", sp3, r)
				}
				count, err := strconv.Atoi(sp3[0])
				if err != nil {
					log.Fatal("bad count", line, count)
				}
				otherName := sp3[1] + " " + sp3[2]
				otherBags = append(otherBags, BagHold{name: otherName, count: count})
			}
		}
		b := &Bag{name: name, contains: otherBags, visited: false}
		bags[name] = b
		//fmt.Println(b)
	}

	// Walk the tree. Naively (no memoization)
	totalCt := 0
	for _, b := range bags {
		if canContainShinyGold(b, bags) {
			totalCt++
		}
	}
	fmt.Println("total:", totalCt)
	fmt.Println("nested shiny gold ct", nestedBagCt(bags["shiny gold"], bags))
}
