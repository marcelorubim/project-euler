package main

import (
	"fmt"
)

func Problem5() int {
	lim := 20
	i := 1
	num := 0
	for num != i {
		if divisable(i, lim) {
			num = i
			fmt.Println(i)
		} else {
			i++
		}
	}
	return num
}

func divisable(num int, lim int) bool {
	if num < lim {
		return false
	}
	for i := 1; i <= lim; i++ {
		if num%i != 0 {
			return false
		}
	}
	return true
}
