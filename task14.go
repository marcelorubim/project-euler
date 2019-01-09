package main

func Task14(num int) []int {
	var sequence []int
	nextNum := num
	for nextNum > 1 {
		if nextNum%2 == 0 {
			nextNum = nextNum / 2
		} else {
			nextNum = 3*nextNum + 1
		}
		sequence = append(sequence, nextNum)
	}
	return sequence
}
