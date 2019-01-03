package main

func Problem4() int {
	result := 0
	for i := 100; i < 999; i++ {
		for j := 100; j < 999; j++ {
			prod := i * j
			if isPalindrome(prod) && prod > result {
				result = prod
			}
		}
	}
	return result
}
func isPalindrome(a int) bool {
	rev := 0
	num := a
	for num != 0 {
		rem := num % 10
		rev = rev*10 + rem
		num /= 10
	}
	return a == rev
}
