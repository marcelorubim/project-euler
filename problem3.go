package main

import (
	"math"
)

func Problem3(lim int) []int {
	var factors []int
	return task3(lim, factors)
}

func task3(lim int, factors []int) []int {
	if IsPrime(lim) {
		factors = append(factors, lim)
		return factors
	}
	for i := 2; i <= int(math.Floor(float64(lim)/2)); i++ {
		if lim%i == 0 && IsPrime(i) {
			x := lim / i
			factors = append(factors, i)
			return task3(x, factors)
		}
	}
	return factors
}

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

func IsPrimeC(value int, c chan<- bool) {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			c <- false
		}
	}
	c <- value > 1
}
