package main

import (
	//"regexp"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines, err := getLines("day4_input.txt")
	assertOk(err)
	ct := 0
	ct2 := 0
	for _, line := range lines {
		pairs := strings.Split(line, ",")
		e1 := strings.Split(pairs[0], "-")
		e2 := strings.Split(pairs[1], "-")
		e1lo, err := strconv.Atoi(e1[0])
		assertOk(err)
		e1hi, err := strconv.Atoi(e1[1])
		assertOk(err)
		e2lo, err := strconv.Atoi(e2[0])
		assertOk(err)
		e2hi, err := strconv.Atoi(e2[1])
		assertOk(err)
		e2Ine1 := e2lo >= e1lo && e2hi <= e1hi
		e1Ine2 := e1lo >= e2lo && e1hi <= e2hi
		if e2Ine1 || e1Ine2 {
			ct++
			ct2++
			continue
		}
		// -----xxxxxxx----
		// ---------xxxxxx-
		case1 := e2lo >= e1lo && e1hi >= e2lo
		// -----xxxxxxx----
		// ---xxxxxx------
		case2 := e1lo >= e2lo && e2hi >= e1lo
		if case1 || case2 {
			ct2++
		}
	}
	fmt.Println(ct)
	fmt.Println(ct2)
}
