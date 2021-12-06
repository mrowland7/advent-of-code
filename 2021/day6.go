package main

import (
	//	"log"
	//"regexp"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := strings.Split("5,4,3,5,1,1,2,1,2,1,3,2,3,4,5,1,2,4,3,2,5,1,4,2,1,1,2,5,4,4,4,1,5,4,5,2,1,2,5,5,4,1,3,1,4,2,4,2,5,1,3,5,3,2,3,1,1,4,5,2,4,3,1,5,5,1,3,1,3,2,2,4,1,3,4,3,3,4,1,3,4,3,4,5,2,1,1,1,4,5,5,1,1,3,2,4,1,2,2,2,4,1,2,5,5,1,4,5,2,4,2,1,5,4,1,3,4,1,2,3,1,5,1,3,4,5,4,1,4,3,3,3,5,5,1,1,5,1,5,5,1,5,2,1,5,1,2,3,5,5,1,3,3,1,5,3,4,3,4,3,2,5,2,1,2,5,1,1,1,1,5,1,1,4,3,3,5,1,1,1,4,4,1,3,3,5,5,4,3,2,1,2,2,3,4,1,5,4,3,1,1,5,1,4,2,3,2,2,3,4,1,3,4,1,4,3,4,3,1,3,3,1,1,4,1,1,1,4,5,3,1,1,2,5,2,5,1,5,3,3,1,3,5,5,1,5,4,3,1,5,1,1,5,5,1,1,2,5,5,5,1,1,3,2,2,3,4,5,5,2,5,4,2,1,5,1,4,4,5,4,4,1,2,1,1,2,3,5,5,1,3,1,4,2,3,3,1,4,1,1", ",")
	fish := []int{}
	dayCt := map[int]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0, 6: 0, 7: 0, 8: 0}
	for _, inputStr := range input {
		n, err := strconv.Atoi(inputStr)
		assertOk(err)
		fish = append(fish, n)
		dayCt[n]++
	}

	fmt.Printf("dayCt %v\n", dayCt)
	//part1(fish)
	for day := 0; day < 256; day++ {
		newDayCt := map[int]int{}
		newDayCt[0] = dayCt[1]
		newDayCt[1] = dayCt[2]
		newDayCt[2] = dayCt[3]
		newDayCt[3] = dayCt[4]
		newDayCt[4] = dayCt[5]
		newDayCt[5] = dayCt[6]
		newDayCt[6] = dayCt[7] + dayCt[0]
		newDayCt[7] = dayCt[8]
		newDayCt[8] = dayCt[0]
		dayCt = newDayCt
	}
	fmt.Printf("how many fish? %v\n", dayCt[0]+dayCt[1]+dayCt[2]+dayCt[3]+dayCt[4]+dayCt[5]+dayCt[6]+dayCt[7]+dayCt[8])
}

func part1(fish []int) {
	for day := 0; day < 80; day++ {
		fmt.Printf("======== day %v: num fish %v\n", day, len(fish))
		startLen := len(fish)
		for i := 0; i < startLen; i++ {
			if fish[i] == 0 {
				fish[i] = 6
				fish = append(fish, 8)
			} else {
				fish[i] = fish[i] - 1
			}
		}
	}
	fmt.Printf("how many fish? %v\n", len(fish))
}
