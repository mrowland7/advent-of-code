package main

import (
	"log"
	//"regexp"
	"fmt"
	"strconv"
	"strings"
)

type Operator string

const (
	Mul  Operator = "*"
	Add  Operator = "+"
	Iden Operator = "."
)

type Expr struct {
	// Always set
	op Operator

	// Then either these two are set
	left  *Expr
	right *Expr

	// Or this is
	value int
}

func parseExpr(line string) *Expr {
	// fmt.Println("parsing line>", line)
	if len(line) == 1 {
		val, err := strconv.Atoi(line)
		assertOk(err)
		return &Expr{value: val, op: Iden}
	}
	// Not a paren start
	if line[0] != ')' {
		return &Expr{left: parseExpr(string(line[0])), right: parseExpr(line[4:]), op: Operator(string(line[2]))}
	}
	// Parenthesis start.

	// If it's the whole line, simplify it
	var outerParenIndex int
	parenStackCount := 1
	for i := 1; i < len(line); i++ {
		if line[i] == '(' {
			parenStackCount--
			if parenStackCount == 0 {
				outerParenIndex = i
				break
			}
		} else if line[i] == ')' {
			parenStackCount++
		}
	}
	if outerParenIndex == len(line)-1 {
		return parseExpr(line[1:outerParenIndex])
	}

	// Otherwise the paren expr is the left, rest is the right
	return &Expr{
		left:  parseExpr(line[1:outerParenIndex]),
		right: parseExpr(line[outerParenIndex+4:]),
		op:    Operator(string(line[outerParenIndex+2])),
	}
}

func evaluateExpr(expr *Expr) int {
	if expr == nil {
		fmt.Println("got nil input!")
		return 0
	}
	fmt.Println("evaluating", expr)
	switch expr.op {
	case Mul:
		return evaluateExpr(expr.left) * evaluateExpr(expr.right)
	case Add:
		return evaluateExpr(expr.left) + evaluateExpr(expr.right)
	case Iden:
		return expr.value
	}
	return 0
}

func getIndexes(s, sign string) []int {
	x := strings.Index(s, sign)
	if x == -1 {
		return []int{}
	}
	return append(getIndexes(s[x+1:], sign), x)
}

// put parens around every addition
func preprocess(raw string) string {
	fmt.Println("raw>", raw)
	result := raw
	//indexes := getIndexes(raw, "+")
	//numIndexes := len(indexes)
	//numPairsAdded := 0
	for resIdx := 0; resIdx < len(result); resIdx++ {
		if result[resIdx] != '+' {
			continue
		}
		fmt.Printf("resIdx %v... ", resIdx)
		// walk back and front to add parens
		frontParenStack := 0
		for x := resIdx + 1; x < len(result); x++ {
			fmt.Printf("x = %v... ", x)
			if result[x] == ' ' {
				continue
			}
			if result[x] == '(' {
				frontParenStack++
			} else if result[x] == ')' {
				frontParenStack--
			}
			if frontParenStack == 0 {
				// Success, add a paren as the next index
				result = result[0:x+1] + ")" + result[x+1:]
				fmt.Println("Added front paren at", x, ", result is now: ", result)
				break
			}
		}
		backParenStack := 0
		for x := resIdx - 1; x >= 0; x-- {
			fmt.Printf("x = %v >%v<... ", x, string(result[x]))
			if result[x] == ' ' {
				continue
			}
			if result[x] == ')' {
				backParenStack++
			} else if result[x] == '(' {
				backParenStack--
			}
			if backParenStack == 0 {
				// Success, add a paren as the next index
				result = result[0:x] + "(" + result[x:]
				fmt.Println("Added back paren, result is now: ", result)
				break
			}
		}
		resIdx++
	}

	fmt.Println("result>", result)
	return result
}

func main() {
	lines, err := getLines("day18_input.txt")
	// lines, err := getLines("day18_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	totalSum := 0
	for _, line := range lines {
		fmt.Println("===")
		if line[0] == '#' {
			continue
		}
		line = preprocess(line)
		expr := parseExpr(reverse(line))
		result := evaluateExpr(expr)
		fmt.Println(line, "=", result)
		totalSum += result
	}
	fmt.Println("total sum:", totalSum)
}
