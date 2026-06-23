package main

import (
	"math/rand"
	"time"
)

func fillUniqueMatrix(rows, cols int) [][]int {
	rand.Seed(time.Now().UnixNano())
	matrix := make([][]int, rows)
	used := make(map[int]bool)

	for i := 0; i < rows; i++ {
		matrix[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			for {
				num := rand.Intn(1000)
				if !used[num] {
					matrix[i][j] = num
					used[num] = true
					break
				}
			}
		}
	}
	return matrix
}
