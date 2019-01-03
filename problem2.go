package main

import (
	"math"
)

func Problem2(lim int) int {
	curr := 1
	prev := 1
	sum := 0
	for curr < lim {
		var newCurr = prev + curr
		prev = curr
		curr = newCurr
		var rem2 = math.Mod(float64(curr), float64(2))
		if rem2 == 0 {
			sum += curr
		}
	}
	return sum
}
