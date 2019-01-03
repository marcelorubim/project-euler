package main

import (
	"fmt"
	"math"
)

func Problem3(lim int) []int {
	//lim := 600851475143
	var primes []int
	for i := 2; i <= lim; i++ {
		if IsPrime(i) {
			primes = append(primes, i)
		}
	}
	fmt.Println("Number of primes:", len(primes))
	var factors []int
	count := 0
	var percent float64
	for _, num := range primes {
		var perc = math.Floor((float64(count) / float64(len(primes))) * 100)
		if perc != percent {
			fmt.Println(perc, "%")
			percent = perc
		}
		var rem = math.Mod(float64(lim), float64(num))
		if rem == 0 {
			div := lim / num
			factors = append(factors, num)
			if IsPrime(div) {
				factors = append(factors, div)
			}
		}
		count++
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
