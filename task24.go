package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Task24() []int {
	var numbers []int
	m := make(map[string]bool)
	for i := 100000000; i <= 9999999999; i++ {
		s := strconv.Itoa(i)
		if i < 1000000000 {
			s = strings.Join([]string{"0", s}, "")
		}
		if len(s) == 10 && !HasRepetead(s) {
			m[s] = true
			fmt.Println(len(m), s)
		}
		if len(m) == 1000000 {
			fmt.Println(i)
			return numbers
		}
	}
	return numbers
}

func HasRepetead(s string) bool {
	c := strings.Split(s, "")
	m := make(map[string]bool)
	for _, x := range c {
		m[x] = true
	}
	return len(m) < len(c)
}
