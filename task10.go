package main

import "fmt"

func Task10(lim int) int {
	sum := 0
	for i := 0; i < lim; i++ {
		if IsPrime(i) {
			fmt.Println(i)
			sum += i
		}
	}
	return sum
}
