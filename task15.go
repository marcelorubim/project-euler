package main

import "fmt"

func Task15(gridSize int) int64 {
	var grid [21][21]int64
	grid[0][0] = 1
	for i := 0; i < 21; i++ {
		for j := 0; j < 21; j++ {
			if i == 0 && j == 0 {
				grid[i][j] = 1
			} else {
				if i > 0 {
					grid[i][j] += grid[i-1][j]
				}
				if j > 0 {
					grid[i][j] += grid[i][j-1]
				}
			}
			fmt.Println(grid[i][j], i, j)
		}
	}
	fmt.Println(len(grid))
	return grid[20][20]
}
