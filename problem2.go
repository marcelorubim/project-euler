package main

import (
 "fmt"
 "math"
)

func main(){
	curr := 1
	prev := 1
	sum := 0
	for curr < 4000000 {
		fmt.Println(curr, sum)
		var newCurr = prev+curr
		prev = curr
		curr = newCurr
		var rem2 = math.Mod(float64(curr),float64(2))
		if rem2 == 0 {
			sum += curr
		}
	}
	fmt.Println(sum)
}