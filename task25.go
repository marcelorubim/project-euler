package main

import (
	"fmt"
	"math/big"
)

func Task25(length int) int {
	idx := 2
	x := big.NewInt(1)
	y := big.NewInt(1)
	for len(y.Text(10)) < length {
		idx++
		temp := y
		y = x.Add(x, y)
		x = temp
		if idx%100 == 0 {
			fmt.Println(idx, y)
		}
	}
	return idx
}
