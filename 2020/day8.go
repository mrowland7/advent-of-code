package main

import (
	"log"
	//"regexp"
	"fmt"
	"strconv"
	"strings"
)

type Instruction struct {
	cmd string
	num int
}

// Two values:
// - does it loop
// - what's the acc value when it does (or terminates)
func checkAcc(instructions []*Instruction, toChange int) (bool, int) {
	visited := map[int]bool{}
	acc := 0
	index := 0
	for index < len(instructions) {
		//fmt.Println("instruction", index, ":\t", instructions[index])
		_, ok := visited[index]
		if ok {
			//	fmt.Println("visited ", index, "already. acc is ", acc)
			return true, acc
		}
		visited[index] = true
		i := instructions[index]
		c := i.cmd
		if toChange == index {
			if c == "nop" {
				c = "jmp"
			}
			if c == "jmp" {
				c = "nop"
			}
		}
		switch c {
		case "acc":
			{
				acc += i.num
				index++
			}
		case "nop":
			{
				index++
			}
		case "jmp":
			{
				index += i.num
			}
		}
	}
	return false, acc
}

func main() {
	lines, err := getLines("day8_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	instructions := []*Instruction{}
	for ind, line := range lines {
		parts := strings.Split(line, " ")
		if len(parts) != 2 {
			log.Fatal("bad line ", line, parts)
		}
		num, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal("bad num", line, num)
		}
		i := &Instruction{cmd: parts[0], num: num}
		fmt.Println("Instruction", ind, i)
		instructions = append(instructions, i)
	}
	fmt.Println("===============")
	checkAcc(instructions, -1)
	for i := 0; i < len(instructions); i++ {
		loops, acc := checkAcc(instructions, i)
		if !loops {
			fmt.Println("Change index", i, "acc is", acc)
		}
	}
}
