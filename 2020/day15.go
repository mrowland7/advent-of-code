package main

import (
	"log"
	//"regexp"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines, err := getLines("day15_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	//puzzleLen := 2020
	puzzleLen := 30000000
	input := strings.Split(lines[0], ",")
	spoken := map[int][]int{}
	turns := make([]int, puzzleLen+1)
	for turn := 1; turn <= puzzleLen; turn++ {
		if turn <= len(input) {
			n, err := strconv.Atoi(input[turn-1])
			assertOk(err)
			turns[turn] = n
			spoken[n] = []int{turn}
			fmt.Println("turn", turn, "elf says:", n)
			continue
		}
		lastVal := turns[turn-1]
		turnsSpoken, ok := spoken[lastVal]
		if !ok {
			log.Fatal("mismatch!", lastVal, turnsSpoken)
		}
		if len(turnsSpoken) == 1 {
			// New; say 0
			turns[turn] = 0
			spoken[0] = append(spoken[0], turn)
		} else {
			// Previously said; say the diff between the last two
			last := turnsSpoken[len(turnsSpoken)-1]
			penultimate := turnsSpoken[len(turnsSpoken)-2]
			newVal := last - penultimate
			turns[turn] = last - penultimate
			spoken[newVal] = append(spoken[newVal], turn)
		}
		if turn%100000 == 0 {
			fmt.Println("turn", turn, "elf says:", turns[turn])
		}
	}
}
