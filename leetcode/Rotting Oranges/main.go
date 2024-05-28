// You are given an m x n grid where each cell can have one of three values:

// 0 representing an empty cell,
// 1 representing a fresh orange, or
// 2 representing a rotten orange.

// Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.

// Return the minimum number of minutes that must elapse until no cell has a fresh orange. If this is impossible, return -1.

// [2 1 1]
// [1 1 0]
// [0 1 1]

// Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
// Output: 4

// Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
// Output: -1
// Explanation: The orange in the bottom left corner (row 2, column 0) is never rotten, because rotting only happens 4-directionally

// [2 1 1]
// [0 1 1]
// [1 0 1]

// Input: grid = [[0,2]]
// Output: 0
// Explanation: Since there are already no fresh oranges at minute 0, the answer is just 0.
package main

import "fmt"

func main() {
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
}

func orangesRotting(grid [][]int) int {
	r := len(grid)
	c := len(grid[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	queue := [][]int{}
	times, fresh := 0, 0

	for row := 0; row < r; row++ {
		for col := 0; col < c; col++ {
			if grid[row][col] == 1 {
				fresh++
			} else if grid[row][col] == 2 {
				queue = append(queue, []int{row, col})
			}
		}
	}
	if fresh == 0 {
		return 0
	}

	for len(queue) > 0 && fresh > 0 {
		newQueue := [][]int{}
		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			di, dj := current[0], current[1]
			for _, dr := range directions {
				drx, dry := dr[0], dr[1]
				newI, newJ := di+drx, dj+dry
				if newI < 0 || newI >= r || newJ < 0 || newJ >= c || grid[newI][newJ] != 1 {
					continue
				}
				grid[newI][newJ] = 2
				newQueue = append(newQueue, []int{newI, newJ})
				fresh--
			}
		}
		queue = newQueue
		if len(queue) > 0 {
			times++
		}
	}
	if fresh > 0 {
		return -1
	}
	return times
}
