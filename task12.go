package main

import (
	"fmt"
	"math"
)

func Task12(numDivisors int) int {
	count := 0
	triangule := 0
	max := 0
	numFactors := 0
	for numFactors < numDivisors {
		count++
		triangule += count
		numFactors = getNumFactors(triangule)
		if numFactors > max {
			max = numFactors
			fmt.Println(count, ":", triangule, ":", max)
		}
	}
	return triangule
}

func getNumFactors(num int) int {
	var divisors []int
	for i := 1; i <= int(math.Floor(float64(num)/2)); i++ {
		if num%i == 0 {
			divisors = append(divisors, i)
		}
	}
	//fmt.Println(num, divisors)
	return len(divisors) + 1

}
