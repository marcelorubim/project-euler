package main

func Task6(lim int) int {
	sumSquares := 0
	squareSum := 0
	for i := 1; i <= lim; i++ {
		sumSquares += i * i
		squareSum += i
	}
	squareSum = squareSum * squareSum
	return squareSum - sumSquares
}
