package main

import (
	"log"
	//"regexp"
	"fmt"
	"strconv"
	//	"strings"
)

// Returns p1's deck, p2's deck, and whether p1 won.
func combat(p1Deck []int, p2Deck []int) ([]int, []int, bool) {
	fmt.Println("=== Game ===")
	prevHands := map[string]bool{}
	numRounds := 0
	for len(p1Deck) > 0 && len(p2Deck) > 0 {
		numRounds++
		fmt.Println("-- Round", numRounds, "--")
		fmt.Printf("Player 1's deck: %v\n", p1Deck)
		fmt.Printf("Player 2's deck: %v\n", p2Deck)
		fmt.Printf("Player 1 plays: %v\n", p1Deck[0])
		fmt.Printf("Player 2 plays: %v\n", p2Deck[0])
		_, loop := prevHands[fmt.Sprintf("%v", p1Deck)+";"+fmt.Sprintf("%v", p2Deck)]
		if loop {
			fmt.Println("Player 1 wins due to looping")
			return p1Deck, p2Deck, true
		}
		prevHands[fmt.Sprintf("%v", p1Deck)+";"+fmt.Sprintf("%v", p2Deck)] = true
		p1Wins := true
		if len(p1Deck)-1 >= p1Deck[0] && len(p2Deck)-1 >= p2Deck[0] {
			p1SubDeck := make([]int, p1Deck[0]+1)
			p2SubDeck := make([]int, p2Deck[0]+1)
			copy(p1SubDeck, p1Deck)
			copy(p2SubDeck, p2Deck)
			_, _, p1Wins = combat(p1SubDeck[1:], p2SubDeck[1:])
			fmt.Println("Player 1 wins?", p1Wins, "due to subgame")
		} else {
			p1Wins = p1Deck[0] > p2Deck[0]
		}

		if p1Wins {
			p1Deck = append(append(p1Deck[1:], p1Deck[0]), p2Deck[0])
			p2Deck = p2Deck[1:]
		} else {
			p2Deck = append(append(p2Deck[1:], p2Deck[0]), p1Deck[0])
			p1Deck = p1Deck[1:]
		}
	}
	p1Wins := len(p1Deck) > 0
	return p1Deck, p2Deck, p1Wins
}

func main() {
	lines, err := getLines("day22_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Get decks
	p1Deck := []int{}
	p2Deck := []int{}
	p1 := true
	for _, line := range lines {
		if line == "" {
			p1 = false
			continue
		}
		if line[0] == 'P' {
			continue
		}
		n, err := strconv.Atoi(line)
		assertOk(err)
		if p1 {
			p1Deck = append(p1Deck, n)
		} else {
			p2Deck = append(p2Deck, n)
		}
	}
	fmt.Println("p1", p1Deck, len(p1Deck))
	fmt.Println("p2", p2Deck, len(p2Deck))

	// Play game
	p1Deck, p2Deck, _ = combat(p1Deck, p2Deck)

	// Get scores
	wDeck := p1Deck
	if len(p2Deck) > 0 {
		wDeck = p2Deck
	}
	score := 0
	for i := 0; i < len(wDeck); i++ {
		score += (len(wDeck) - i) * wDeck[i]
	}
	fmt.Println("winning score is", score)
}
