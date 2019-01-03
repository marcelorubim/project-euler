package main

func Task7(position int) int {
	var primes []int
	num := 2
	for len(primes) < position {
		if IsPrime(num) {
			primes = append(primes, num)
		}
		num++
	}
	return primes[len(primes)-1]
}
