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
	//Read the file line by line
	for scanner.Scan() {
		s := scanner.Text()
		lines = append(lines, strings.Split(s, " "))
	}

	var rows [][]int
	//Convert the string value to int
	for _, r := range lines {
		row := []int{}
		for _, c := range r {
			v, _ := strconv.Atoi(c)
			row = append(row, v)
		}
		rows = append(rows, row)
	}
	return sumNodes(0, 0, 0, rows)
}

func solution2(rows [][]int) int {
	for i := len(rows) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			rows[i][j] += int(math.Max(float64(rows[i+1][j]), float64(rows[i+1][j+1])))
		}
	}
	return int(rows[0][0])
}

func sumNodes(sum int, index int, nextLine int, lines [][]int) int {
	if nextLine >= len(lines) {
		return sum
	}
	line := lines[nextLine]
	sum += line[index]
	sumL := sumNodes(sum, index, nextLine+1, lines)
	sumR := sumNodes(sum, index+1, nextLine+1, lines)
	return int(math.Max(float64(sumL), float64(sumR)))
}
