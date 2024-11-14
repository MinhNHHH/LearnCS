package main

import (
	"fmt"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	hash := make(map[[2]int]int)
	var dfs func(row, col int) int
	dfs = func(row, col int) int {
		directions := [][]int{{-1, 1}, {0, 1}, {1, 1}}
		if val, found := hash[[2]int{row, col}]; found {
			return val
		}
		maxMove := 0

		for _, dir := range directions {
			newRow, newCol := row+dir[0], col+dir[1]

			if newRow >= 0 && newRow < m && newCol < n && grid[newRow][newCol] > grid[row][col] {
				maxMove = max(maxMove, dfs(newRow, newCol)+1)
			}
		}
		hash[[2]int{row, col}] = maxMove
		return maxMove
	}
	res := 0
	for row := 0; row < m; row++ {
		res = max(res, dfs(row, 0))
	}

	return res
}

func main() {
	fmt.Println(maxMoves([][]int{{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}}))
}
