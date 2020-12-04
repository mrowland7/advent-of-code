package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
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

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func validPassport1(p *Passport) bool {
	if p == nil {
		return false
	}
	// Length
	if !(len(p.byr) > 0 &&
		len(p.iyr) > 0 &&
		len(p.eyr) > 0 &&
		len(p.hgt) > 0 &&
		len(p.hcl) > 0 &&
		len(p.ecl) > 0 &&
		len(p.pid) > 0) {
		return false
	}

	return true
}

func validPassport2(p *Passport) bool {
	if p == nil {
		return false
	}
	// Required fields
	if !(len(p.byr) > 0 &&
		len(p.iyr) > 0 &&
		len(p.eyr) > 0 &&
		len(p.hgt) > 0 &&
		len(p.hcl) > 0 &&
		len(p.ecl) > 0 &&
		len(p.pid) > 0) {
		fmt.Println("missing field")
		return false
	}
	// 	byr (Birth Year) - four digits; at least 1920 and at most 2002.
	byr, err := strconv.Atoi(p.byr)
	if err != nil || byr < 1920 || byr > 2002 {
		fmt.Println("invalid byr", p.byr)
		return false
	}
	// 	iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	iyr, err := strconv.Atoi(p.iyr)
	if err != nil || iyr < 2010 || iyr > 2020 {
		fmt.Println("invalid iyr", p.iyr)
		return false
	}
	// 	eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	eyr, err := strconv.Atoi(p.eyr)
	if err != nil || eyr < 2020 || eyr > 2030 {
		fmt.Println("invalid eyr", p.eyr)
		return false
	}
	// 	hgt (Height) - a number followed by either cm or in:
	// 	  If cm, the number must be at least 150 and at most 193.
	// 	  If in, the number must be at least 59 and at most 76.
	if len(p.hgt) < 4 || len(p.hgt) > 5 {
		fmt.Println("invalid hgt lennn", p.hgt)
		return false
	}
	dim := p.hgt[len(p.hgt)-2:]
	switch dim {
	case "in":
		val, err := strconv.Atoi(p.hgt[0 : len(p.hgt)-2])
		if err != nil || val < 59 || val > 76 {
			fmt.Println("invalid inches", p.hgt)
			return false
		}
	case "cm":
		val, err := strconv.Atoi(p.hgt[0 : len(p.hgt)-2])
		if err != nil || val < 150 || val > 193 {
			fmt.Println("invalid cm", p.hgt)
			return false
		}
	default:
		return false
	}

	// 	hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	match, err := regexp.Match("#[[:xdigit:]]{6}", []byte(p.hcl))
	if err != nil {
		log.Fatal(err)
	}
	if !match {
		fmt.Println("invalid hcl", p.hcl)
		return false
	}
	// 	ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	if p.ecl != "amb" &&
		p.ecl != "blu" &&
		p.ecl != "brn" &&
		p.ecl != "gry" &&
		p.ecl != "grn" &&
		p.ecl != "hzl" &&
		p.ecl != "oth" {
		fmt.Println("invalid ecl", p.ecl)
		return false
	}
	// 	pid (Passport ID) - a nine-digit number, including leading zeroes.
	_, err = strconv.Atoi(p.pid)
	if len(p.pid) != 9 || err != nil {
		fmt.Println("invalid pid", p.pid)
		return false
	}

	return true
}

func main() {
	lines, err := getLines("day4_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	passports := []*Passport{}
	collected_line := ""
	for _, line := range lines {
		if len(line) > 0 {
			collected_line = collected_line + " " + line
			continue
		}
		// Have everything in one line; parse
		passport := &Passport{}
		fields := strings.Fields(collected_line)
		for _, field := range fields {
			parts := strings.Split(field, ":")
			if len(parts) != 2 {
				continue
			}
			switch parts[0] {
			case "byr":
				passport.byr = parts[1]
			case "iyr":
				passport.iyr = parts[1]
			case "eyr":
				passport.eyr = parts[1]
			case "hgt":
				passport.hgt = parts[1]
			case "hcl":
				passport.hcl = parts[1]
			case "ecl":
				passport.ecl = parts[1]
			case "pid":
				passport.pid = parts[1]
			case "cid":
				passport.cid = parts[1]
			}
		}
		passports = append(passports, passport)
		fmt.Println(collected_line)
		fmt.Println(passport)
		collected_line = ""
	}
	validCt := 0
	for _, p := range passports {
		if validPassport2(p) {
			validCt++
		}
	}
	fmt.Println("total", len(passports), "valid", validCt)
}
