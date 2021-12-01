package main

import (
	"log"
	//"regexp"
	"fmt"
	"strconv"
	"strings"
)

type TicketRule struct {
	name      string
	range1_lo int
	range1_hi int
	range2_lo int
	range2_hi int
}

type Ticket struct {
	values []int
}

func parseTicket(line string) Ticket {
	valuesStr := strings.Split(line, ",")
	values := []int{}
	for _, s := range valuesStr {
		values = append(values, mustParseInt(s))
	}
	return Ticket{values: values}
}

func mustParseInt(value string) int {
	n, err := strconv.Atoi(value)
	assertOk(err)
	return n
}

func ticketErrorRates(nearbyTickets []Ticket, rules []TicketRule) {
}

func main() {
	lines, err := getLines("day16_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	block := 1
	rules := []TicketRule{}
	var myTicket Ticket
	nearbyTickets := []Ticket{}
	for _, line := range lines {
		if line == "" {
			block++
			continue
		}
		if block == 1 {
			parts1 := strings.Split(line, ": ")
			parts2 := strings.Split(parts1[1], " or ")
			rule1 := strings.Split(parts2[0], "-")
			rule2 := strings.Split(parts2[1], "-")
			rules = append(rules, TicketRule{
				name:      parts1[0],
				range1_lo: mustParseInt(rule1[0]),
				range1_hi: mustParseInt(rule1[1]),
				range2_lo: mustParseInt(rule2[0]),
				range2_hi: mustParseInt(rule2[1]),
			})
		} else if block == 2 {
			if line[0:3] == "you" {
				continue
			}
			myTicket = parseTicket(line)
		} else if block == 3 {
			if line[0:3] == "nea" {
				continue
			}
			nearbyTickets = append(nearbyTickets, parseTicket(line))
		}
	}
	fmt.Println("rules are:", rules)
	fmt.Println("my ticket is:", myTicket)

	invalidSum := 0
	validTickets := []Ticket{}
	for _, ticket := range nearbyTickets {
		validTicket := true
		for _, value := range ticket.values {
			validSomewhere := false
			for _, rule := range rules {
				if (value >= rule.range1_lo && value <= rule.range1_hi) ||
					(value >= rule.range2_lo && value <= rule.range2_hi) {
					validSomewhere = true
					break
				}
			}
			if !validSomewhere {
				validTicket = false
				invalidSum += value
			}
		}
		if validTicket {
			validTickets = append(validTickets, ticket)
		}
	}
	fmt.Println("invalid sum:", invalidSum)
	validIndexes := map[string][]int{}
	for _, rule := range rules {
		validIndexes[rule.name] = []int{}
		for i := 0; i < len(myTicket.values); i++ {
			validIndex := true
			for _, ticket := range validTickets {
				value := ticket.values[i]
				allowed := (value >= rule.range1_lo && value <= rule.range1_hi) ||
					(value >= rule.range2_lo && value <= rule.range2_hi)
				if !allowed {
					validIndex = false
				}
			}
			if validIndex {
				validIndexes[rule.name] = append(validIndexes[rule.name], i)
			}
		}
	}
	mapping := map[string]int{}
	taken := map[int]bool{}
	for n := 0; n < 25; n++ {
		fmt.Println("===")
		fmt.Println("iter", n, "mapping", mapping, "validIndexes", validIndexes)
		remaining := map[string][]int{}
		for field, validIndexes := range validIndexes {
			// Assign
			if len(validIndexes) == 1 {
				mapping[field] = validIndexes[0]
				taken[validIndexes[0]] = true
				continue
			}
			remaining[field] = []int{}
			for _, idx := range validIndexes {
				if _, gone := taken[idx]; !gone {
					remaining[field] = append(remaining[field], idx)
				}
			}
		}
		validIndexes = remaining
	}
	fmt.Println("mapping is:", mapping)

	product := 1
	for mapping, idx := range mapping {
		if len(mapping) >= 9 && mapping[0:5] == "depar" {
			product *= myTicket.values[idx]
		}
	}
	fmt.Println("product:", product)
}
