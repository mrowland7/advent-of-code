package main

import (
	//"regexp"
	"fmt"
	//	"strconv"
	//	"strings"
)

func getLoopSize(key int) int {
	value := 1
	loopSize := 0
	subjectNumber := 7
	for value != key {
		loopSize++
		value = (value * subjectNumber) % 20201227
	}
	return loopSize
}

func getKey(subjectNumber int, loopSize int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value = value * subjectNumber % 20201227
	}
	return value
}

func main() {
	cardKey := 3418282
	doorKey := 8719412

	fmt.Println(getLoopSize(5764801))
	fmt.Println(getLoopSize(17807724))

	cardLoopSize := getLoopSize(cardKey)
	doorLoopSize := getLoopSize(doorKey)
	fmt.Println("card", cardLoopSize)
	fmt.Println("door", doorLoopSize)

	key1 := getKey(doorKey, cardLoopSize)
	key2 := getKey(cardKey, doorLoopSize)
	fmt.Println("key1", key1)
	fmt.Println("key2", key2)
}
