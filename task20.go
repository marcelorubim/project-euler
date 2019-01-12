package main

import (
	"fmt"
	"math/big"
)

func Task20(n int) int64 {
	var sum int64
	factorial := big.NewInt(int64(n))

	for i := n - 1; i > 0; i-- {
		factorial = factorial.Mul(factorial, big.NewInt(int64(i)))
	}
	fmt.Println(factorial.Int64())
	for factorial.Text(10) != "0" {
		rem := big.NewInt(int64(0)).Mod(factorial, big.NewInt(int64(10)))
		sum += rem.Int64()
		factorial = factorial.Div(factorial, big.NewInt(int64(10)))
		fmt.Println(factorial)
	}

	return sum
}
