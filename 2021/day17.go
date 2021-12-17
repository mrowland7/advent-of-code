package main

import (
	//"regexp"
	"fmt"
	//	"strconv"
	//	"strings"
)

func main() {
	xTargetMin := 241
	xTargetMax := 275
	yTargetMin := -75
	yTargetMax := -49

	biggestY := -1
	argmaxX := -1
	argmaxY := -1
	numWork := 0
	for xs := 20; xs <= 275; xs++ {
		for ys := -75; ys <= 10000; ys++ {
			//fmt.Println("trying x=", xs, "y=", ys)
			xv := xs
			yv := ys
			x := 0
			y := 0
			maxY := 0
			for step := 1; step < 200; step++ {
				x = x + xv
				y = y + yv
				if y > maxY {
					maxY = y
				}
				if x >= xTargetMin && x <= xTargetMax && y >= yTargetMin && y <= yTargetMax {
					fmt.Println("---- xs=", xs, "ys=", ys, "hit at x=", x, "y=", y)
					numWork++
					if maxY > biggestY {
						argmaxX = xs
						argmaxY = ys
						biggestY = maxY
					}
					break
				}
				if xv > 0 {
					xv -= 1
				}
				yv -= 1
				if x > xTargetMax && y < yTargetMin {
					//	fmt.Println("---- too far")
					break
				}
			}
		}
	}
	fmt.Println("biggest is x=", argmaxX, "y=", argmaxY, "which hit a height of", biggestY)
	fmt.Println("num that worked", numWork)
}
