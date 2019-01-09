package main

import (
	"math/big"
	"strconv"
	"strings"
)

func Task16(exp int) int {
	num := big.NewInt(2)
	num.Exp(big.NewInt(2), big.NewInt(1000), big.NewInt(0))
	sum := 0
	numbers := strings.Split(num.String(), "")
	for _, x := range numbers {
		val, _ := strconv.Atoi(x)
		sum += val
	}
	return sum
}
