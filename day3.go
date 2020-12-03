package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	//"strconv"
	//"strings"
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

func checkJump(lines []string, xJump int, yJump int) int {
	x := 0
	width := len(lines[0])
	treeHitCt := 0
	for i, line := range lines {
		if i%yJump != 0 {
			continue
		}
		xIdx := x % width
		//hit := false
		if line[xIdx] == byte('#') {
			//	hit = true
			treeHitCt += 1
		}
		//fmt.Printf("line %v: x %v, xIdx %v, hit %v, \t %v\n", i, x, xIdx, hit, line)
		x += xJump
	}
	return treeHitCt
}

func main() {
	lines, err := getLines("day3_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	first := checkJump(lines, 1, 1)
	fmt.Println("1/1", first)
	second := checkJump(lines, 3, 1)
	fmt.Println("3/1", second)
	third := checkJump(lines, 5, 1)
	fmt.Println("5/1", third)
	fourth := checkJump(lines, 7, 1)
	fmt.Println("7/1", fourth)
	fifth := checkJump(lines, 1, 2)
	fmt.Println("1/2", fifth)

	fmt.Println("combined product", first*second*third*fourth*fifth)
}
