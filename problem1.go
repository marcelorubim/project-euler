package main

import (
	"math"
)

func Problem1() int {
	sum := 0
	for i := 0; i < 1000; i++ {
		var rem3 = math.Mod(float64(i), float64(3))
		var rem5 = math.Mod(float64(i), float64(5))
		if rem3 == 0 || rem5 == 0 {
			sum += i
		}
	}
	return sum
}
