package main

import (
	//"regexp"
	"fmt"
	//	"strconv"
	//	"strings"
)

func main() {
	lines, err := getLines("day7_input.txt")
	assertOk(err)
	for _, line := range lines {
		fmt.Println(line)
	}
}
