package main

import (
	//"regexp"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func earliestBus(buses []int, earliestTime int) {
outer:
	for i := earliestTime; i < earliestTime+1000; i++ {
		for _, b := range buses {
			if b != 1 && i%b == 0 {
				fmt.Println("can take bus", b, "at time", i, "product", b*(i-earliestTime))
				break outer
			}
		}
	}
}

func check(workerId int, numWorkers int, buses []int, wg *sync.WaitGroup) {
	start := 101088700000000
	//start := 100000000000000
	//start := 0
	// too high = 2990035291265711
	fmt.Println("worker ", workerId, "started")
	for i := start + workerId*buses[0]; i > -1; i += buses[0] * numWorkers {
		if i%100000000 == 0 {
			fmt.Println("i=", i)
		}
		valid := true
		for ind, b := range buses {
			if (i+ind)%b != 0 {
				valid = false
				break
			}
		}
		if valid {
			fmt.Println("Found it!", i)
			wg.Done()
			break
		}
	}
}

// brute force
func schedule(buses []int) {
	const numWorkers = 6
	wg := sync.WaitGroup{}
	wg.Add(1)
	for w := 1; w < numWorkers; w++ {
		go check(w, numWorkers, buses, &wg)
	}
	wg.Wait()
}

func findInverse(x, n int) int {
	x = x % n
	for i := 1; i < n; i++ {
		if (x*i)%n == 1 {
			return i
		}
	}
	return -1
}

// for (0, n1), (1, x), (2, x), (3, n2), (4, x), (5, n3)
// Want to find a value Z such that
// Z ~= 0 mod n1
// Z + 3 ~= 0 mod n2
// Z + 5 ~= 0 mod n3
// And since they're coprime you can use the chinese remainder theorem:
// https://crypto.stanford.edu/pbc/notes/numbertheory/crt.html
func scheduleSmart(buses []int) {
	bigM := 1
	for _, b := range buses {
		bigM *= b
	}
	x := 0
	for i, n := range buses {
		if n == 1 {
			continue
		}
		a_i := n - i
		b_i := bigM / n
		b_i_prime := findInverse(b_i, n)
		fmt.Println("X ~=", i, "mod", n, "; a_i:", a_i, "b_i:", b_i, "b_i_prime:", b_i_prime)
		x += a_i * b_i * b_i_prime
	}
	result := x % bigM
	fmt.Println("result is", result)
	for i, n := range buses {
		if n == 1 {
			continue
		}
		fmt.Println("bus", n, "mod result", result, "=", result%n, ", want", i)
	}
}

func main() {
	lines, err := getLines("day13_input.txt")
	//	lines, err := getLines("day13_dbg_2.txt")
	assertOk(err)
	//earliestTime, err := strconv.Atoi(lines[0])
	// assertOk(err)
	busesRaw := strings.Split(lines[1], ",")
	buses := []int{}
	for _, b := range busesRaw {
		if b != "x" {
			parsed, err := strconv.Atoi(b)
			assertOk(err)
			buses = append(buses, parsed)
		} else {
			buses = append(buses, 1)
		}
	}
	// earliestBus(buses, earliestTime)
	//schedule(buses)
	scheduleSmart(buses)
}
