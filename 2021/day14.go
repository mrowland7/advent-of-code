package main

import (
	//"regexp"
	"fmt"
	//	"strconv"
	"strings"
)

func main() {
	lines, err := getLines("day14.txt")
	assertOk(err)
	polymer := "FSHBKOOPCFSFKONFNFBB"
	//	polymer := "NNCB"
	//	lines := []string{
	//		"CH -> B",
	//		"HH -> N",
	//		"CB -> H",
	//		"NH -> C",
	//		"HB -> C",
	//		"HC -> B",
	//		"HN -> C",
	//		"NN -> C",
	//		"BH -> H",
	//		"NC -> B",
	//		"NB -> B",
	//		"BN -> B",
	//		"BB -> N",
	//		"BC -> B",
	//		"CC -> N",
	//		"CN -> C",
	//	}
	rules := map[string]string{}
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = parts[1]
	}
	pairs := map[string]int64{}
	for i := 0; i < len(polymer)-1; i++ {
		s := polymer[i : i+2]
		if _, ok := pairs[s]; !ok {
			pairs[s] = 0
		}
		pairs[s]++
	}

	for step := 1; step <= 40; step++ {
		newPairs := map[string]int64{}
		for pair, ct := range pairs {
			ruleResult := rules[pair]
			new1 := string(pair[0]) + ruleResult
			new2 := ruleResult + string(pair[1])
			if _, ok := newPairs[new1]; !ok {
				newPairs[new1] = 0
			}
			if _, ok := newPairs[new2]; !ok {
				newPairs[new2] = 0
			}
			newPairs[new1] += ct
			newPairs[new2] += ct
		}
		pairs = newPairs
		counts := map[string]int64{}
		for k, v := range pairs {
			counts[string(k[0])] += v
		}
		counts["B"]++ // hackerman
		mx := int64(0)
		min := int64(1000000000000000000)
		ln := int64(0)
		for _, v := range counts {
			if v > mx {
				mx = v
			}
			if v < min {
				min = v
			}
			ln += v
		}
		fmt.Println("after step", step, "counts are", counts, "diff", mx-min, "len", ln)
	}

	//	// naive
	//	for step := 1; step <= 10; step++ {
	//		newPolymer := ""
	//		counts := map[string]int{}
	//		for i := 0; i < len(polymer)-1; i++ {
	//			start := string(polymer[i])
	//			ruleResult := rules[polymer[i:i+2]]
	//			if _, ok := counts[start]; !ok {
	//				counts[start] = 0
	//			}
	//			if _, ok := counts[ruleResult]; !ok {
	//				counts[ruleResult] = 0
	//			}
	//			counts[start]++
	//			counts[ruleResult]++
	//			newPolymer += start + ruleResult
	//		}
	//		last := string(polymer[len(polymer)-1])
	//		newPolymer += last
	//		counts[last]++
	//		polymer = newPolymer
	//		mx := 0
	//		min := 100000000
	//		ln := 0
	//		for _, v := range counts {
	//			if v > mx {
	//				mx = v
	//			}
	//			if v < min {
	//				min = v
	//			}
	//			ln += v
	//		}
	//		fmt.Println("after step", step, "counts are", counts, "diff", mx-min, "len", ln)
	//		//fmt.Println("string", newPolymer)
	//	}
}
