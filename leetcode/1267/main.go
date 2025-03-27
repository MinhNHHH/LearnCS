package main

import "fmt"

func countServers(grid [][]int) int {
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := [][]int{grid[0]}
	count := 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, direct := range directions {
			newX, newY := current[0]+direct[0], current[1]+direct[1]
			if newX >= 0 && newY >= 0 && newX < len(grid) && newY < len(grid[0]) {
				if grid[current[0]][current[1]] == 1 && grid[newX][newY] == grid[current[0]][current[1]] {
					count++
					grid[current[0]][current[1]] = 2
				}
			}
		}

	}
	return count
}

func main() {
	fmt.Println(countServers([][]int{{1, 0}, {0, 1}}))
}
