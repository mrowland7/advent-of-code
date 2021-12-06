package main

import (
	"log"
	//"regexp"
	"fmt"
	//	"strconv"
	//	"strings"
)

func main() {
	lines, err := getLines("day7_input.txt")
	asserOk(err)
	for _, line := range lines {
		fmt.Println(line)
	}
}
