package main

import (
	"log"
	//"regexp"
	"fmt"
	"sort"
	"strconv"
	//	"strings"
)

func main() {
	lines, err := getLines("day9_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	nums := []int64{}
	var invalidNum int64 = -1
	for i, line := range lines {
		x, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		n := int64(x)
		if len(nums) < 25 {
			nums = append(nums, n)
			continue
		}
		valid := false
	outer:
		for j := i - 25; j < len(nums); j++ {
			for k := j; k < len(nums); k++ {
				sum := nums[j] + nums[k]
				if sum == n {
					valid = true
					break outer
				}
			}
		}
		if !valid {
			invalidNum = n
			fmt.Println(invalidNum, "is not valid")
		}
		nums = append(nums, n)
	}
	// sums[i][j] is the sum of nums i through j.
	sums := make([][]int64, len(nums))
	for i, _ := range nums {
		sums[i] = make([]int64, len(nums))
		for j := i; j < len(nums); j++ {
			if j == i {
				sums[i][j] = nums[j]
			} else {
				sums[i][j] = sums[i][j-1] + nums[j]
			}
			if sums[i][j] == invalidNum {
				fmt.Println("invalid range is ", i, "to", j)
				subRange := nums[i : j+1]
				sort.Slice(subRange, func(i, j int) bool {
					return subRange[i] < subRange[j]
				})
				fmt.Println(subRange, subRange[0]+subRange[len(subRange)-1])
			}
		}
	}

}
