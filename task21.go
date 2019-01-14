package main

func Task21() int {
	sum := 0
	var amicables []int
	for i := 1; i < 10000; i++ {
		n := SumProperDivisors(i)
		sumProperN := SumProperDivisors(n)

		if i != n && sumProperN == i {
			amicables = append(amicables, i)
		}
	}
	for _, x := range amicables {
		sum += x
	}
	return sum
}

func SumProperDivisors(n int) int {
	sum := 0
	for i := n / 2; i > 0; i-- {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}
