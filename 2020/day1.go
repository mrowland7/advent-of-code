package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readNums() map[int]int {
	file, err := os.Open("day1_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	nums := map[int]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		if _, ok := nums[num]; !ok {
			nums[num] = 0
		}
		nums[num]++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nums
}

func part1(nums map[int]int) {
	fmt.Printf("nums %v\n", nums)

	for num1, _ := range nums {
		for num2, _ := range nums {
			if num1+num2 == 2020 {
				fmt.Printf("%v + %v = %v, %v * %v = %v\n", num1, num2, num1+num2, num1, num2, num1*num2)
			}
		}
	}
}

func part2(nums map[int]int) {
	fmt.Printf("nums %v\n", nums)

	for num1, _ := range nums {
		for num2, _ := range nums {
			for num3, _ := range nums {
				if num1+num2+num3 == 2020 {
					fmt.Printf("%v + %v + %v = %v, %v * %v * %v= %v\n", num1, num2, num3, num1+num2+num3, num1, num2, num3, num1*num2*num3)
					return
				}
			}
		}
	}
}

func main() {
	nums := readNums()
	part2(nums)
}
