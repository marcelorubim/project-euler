package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Task18(filename string) int {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines [][]string

	for scanner.Scan() {
		s := scanner.Text()
		lines = append(lines, strings.Split(s, " "))
	}

	var rows [][]int

	for _, r := range lines {
		row := []int{}
		for _, c := range r {
			v, _ := strconv.Atoi(c)
			row = append(row, v)
		}
		rows = append(rows, row)
	}

	val, err := strconv.Atoi(lines[0][0])
	check(err)
	sum, _ := sumNodes(val, 0, 1, rows, []int{val})
	sum1, _ := sumNodes(val, 1, 1, rows, []int{val})
	return int(math.Max(float64(sum), float64(sum1)))
}

func solution2(rows [][]int) int {
	for i := len(rows) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			rows[i][j] += int(math.Max(float64(rows[i+1][j]), float64(rows[i+1][j+1])))
		}
	}
	return int(rows[0][0])
}

func sumNodes(sum int, index int, nextLine int, lines [][]int, values []int) (int, []int) {
	if nextLine >= len(lines) {
		return sum, values
	}
	line := lines[nextLine]
	values = append(values, line[index])
	sum += line[index]
	sumL, valL := sumNodes(sum, index, nextLine+1, lines, values)
	sumR, valR := sumNodes(sum, index+1, nextLine+1, lines, values)
	if sumL > sumR {
		return sumL, valL
	}
	return sumR, valR
}
