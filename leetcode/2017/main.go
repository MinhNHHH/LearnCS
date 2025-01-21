package main

import (
	"fmt"
	"math"
)

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func gridGame(grid [][]int) int64 {
	minRes := int64(math.MaxInt)
	sumRow1 := int64(0)

	for _, value := range grid[0] {
		sumRow1 += int64(value)
	}

	sumRow2 := int64(0)
	for i, value := range grid[1] {
		sumRow1 -= int64(grid[0][i])

		if sumRow1 < sumRow2 {
			minRes = min(minRes, sumRow2)
		} else {
			minRes = min(minRes, sumRow1)
		}
		sumRow2 += int64(value)
	}
	return int64(minRes)
}

func main() {
	fmt.Println(gridGame([][]int{{2, 5, 4}, {1, 5, 1}}))
}
