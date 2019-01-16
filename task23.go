package main

import (
	"fmt"
	"math"
	"sort"
)

func Task23() int {
	sum := 0
	sumAllAbundantNums := GetAllSumTowAbundant()
	for i := 1; i <= 28123; i++ {
		if !IsSumTwoAbundant(i, sumAllAbundantNums) {
			sum += i
		}
	}
	return sum
}

func IsSumTwoAbundant(num int, sumAllAbundantNums []int) bool {
	for _, i := range sumAllAbundantNums {
		if i == num {
			return true
		}
		if i > num {
			return false
		}
	}
	return false
}

func GetAllSumTowAbundant() []int {
	abundantNums := GetAllAbundant()
	fmt.Println(abundantNums)
	m := make(map[int]bool)

	var sumAllAbundantNums []int
	for i := 0; i < len(abundantNums)-1; i++ {
		for j := i; j < len(abundantNums); j++ {
			sum := abundantNums[i] + abundantNums[j]
			_, prs := m[sum]
			if sum <= 28123 && !prs {
				sumAllAbundantNums = append(sumAllAbundantNums, sum)
				m[sum] = true
			}
		}
	}
	sort.Ints(sumAllAbundantNums)
	return sumAllAbundantNums
}

func GetAllAbundant() []int {
	var abundantNums []int
	for i := 1; i <= 28123; i++ {
		if Abundant(i) {
			abundantNums = append(abundantNums, i)
		}
	}
	return abundantNums
}

func Abundant(num int) bool {
	sumFactors := 0
	for i := 1; i <= int(math.Floor(float64(num)/2)); i++ {
		if num%i == 0 {
			sumFactors += i
		}
	}
	return sumFactors > num
}
