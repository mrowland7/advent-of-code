package main

import (
	//"regexp"
	"fmt"
	"log"
	"math"
	"strconv"
	//	"strings"
)

type Instruction struct {
	dir    string
	amount int
}

func getNextDir(current string, direction string, degrees int) string {
	if degrees == 0 {
		return current
	}
	if degrees%90 != 0 {
		log.Fatal("bad degrees: ", degrees)
	}
	switch current {
	case "N":
		if direction == "L" {
			return getNextDir("W", "L", degrees-90)
		} else {
			return getNextDir("E", "R", degrees-90)
		}
	case "S":
		if direction == "L" {
			return getNextDir("E", "L", degrees-90)
		} else {
			return getNextDir("W", "R", degrees-90)
		}
	case "E":
		if direction == "L" {
			return getNextDir("N", "L", degrees-90)
		} else {
			return getNextDir("S", "R", degrees-90)
		}
	case "W":
		if direction == "L" {
			return getNextDir("S", "L", degrees-90)
		} else {
			return getNextDir("N", "R", degrees-90)
		}
	default:
		log.Fatal("bad direction: ", direction)
		return ""
	}
}

func normalMoves(instructions []*Instruction) {
	x := 0
	y := 0
	facing := "E"
	for _, ins := range instructions {
		dir := ins.dir
		if dir == "F" {
			dir = facing
		}
		if dir == "R" || dir == "L" {
			facing = getNextDir(facing, ins.dir, ins.amount)
			continue
		}
		switch dir {
		case "N":
			y += ins.amount
		case "S":
			y -= ins.amount
		case "E":
			x += ins.amount
		case "W":
			x -= ins.amount
		}
	}
	fmt.Println(x, y, math.Abs(float64(x))+math.Abs(float64(y)))
}

func waypoints(instructions []*Instruction) {
	waypointX := 10
	waypointY := 1
	shipX := 0
	shipY := 0
	for _, ins := range instructions {
		fmt.Print(ins, ", ship (", shipX, shipY, "), waypoint (", waypointX, waypointY, ")")
		switch ins.dir {
		case "N":
			waypointY += ins.amount
		case "S":
			waypointY -= ins.amount
		case "E":
			waypointX += ins.amount
		case "W":
			waypointX -= ins.amount
		case "L":
			// x = 1, y= 2 -> x = -2, y = 1 -> x = -1, y = -2 -> x = 2, y = -1
			times := ins.amount / 90
			for i := 0; i < times; i++ {
				waypointX, waypointY = -1*waypointY, waypointX
			}
		case "R":
			times := ins.amount / 90
			for i := 0; i < times; i++ {
				waypointX, waypointY = waypointY, -1*waypointX
			}
		case "F":
			shipX += ins.amount * waypointX
			shipY += ins.amount * waypointY
		}
		fmt.Println(" ---->", "ship (", shipX, shipY, "), waypoint (", waypointX, waypointY, ")")
	}
	fmt.Println(shipX, shipY, math.Abs(float64(shipX))+math.Abs(float64(shipY)))
}

func main() {
	lines, err := getLines("day12_input.txt")
	assertOk(err)
	instructions := []*Instruction{}
	for _, line := range lines {
		dir := line[0]
		amount, err := strconv.Atoi(line[1:])
		assertOk(err)
		i := &Instruction{dir: string(dir), amount: amount}
		instructions = append(instructions, i)
	}

	//normalMoves(instructions)
	waypoints(instructions)

}
