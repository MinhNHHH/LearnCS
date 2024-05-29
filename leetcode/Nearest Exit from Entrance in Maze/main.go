// You are given an m x n matrix maze (0-indexed) with empty cells (represented as '.') and walls (represented as '+').
// You are also given the entrance of the maze, where entrance = [entrancerow, entrancecol] denotes the row and column of the cell you are initially standing at.

// In one step, you can move one cell up, down, left, or right.
// You cannot step into a cell with a wall, and you cannot step outside the maze.
// Your goal is to find the nearest exit from the entrance. An exit is defined as an empty cell that is at the border of the maze.
// The entrance does not count as an exit.

// Return the number of steps in the shortest path from the entrance to the nearest exit, or -1 if no such path exists.

// Input: maze = [["+","+",".","+"],[".",".",".","+"],["+","+","+","."]], entrance = [1,2]
// Output: 1
// Explanation: There are 3 exits in this maze at [1,0], [0,2], and [2,3].
// Initially, you are at the entrance cell [1,2].
// - You can reach [1,0] by moving 2 steps left.
// - You can reach [0,2] by moving 1 step up.
// It is impossible to reach [2,3] from the entrance.
// Thus, the nearest exit is [0,2], which is 1 step away.

// [+ + . +]
// [. . x +]
// [+ + + .]

// Input: maze = [["+","+","+"],[".",".","."],["+","+","+"]], entrance = [1,0]

// Input: maze = [[".","+"]], entrance = [0,0]
// Output: -1
// Explanation: There are no exits in this maze.

package main

import (
	"fmt"
)

func main() {
	fmt.Println(nearestExit([][]byte{{43, 43, 46, 43}, {46, 46, 46, 43}, {43, 43, 43, 46}}, []int{1, 2}))
	fmt.Println(nearestExit([][]byte{{46, 43}}, []int{0, 0}))
}

func nearestExit(maze [][]byte, entrance []int) int {

	//                       L        U       R        D
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	queue := [][]int{{entrance[0], entrance[1], 0}}
	maze[entrance[0]][entrance[1]] = 43
	r := len(maze)
	c := len(maze[0])
	for len(queue) > 0 {
		current := queue[0]
		di, dj, step := current[0], current[1], current[2]
		queue = queue[1:]
		if di == 0 || dj == 0 || di == r-1 || dj == c-1 {
			if step > 0 {
				return step
			}
		}

		for _, dr := range directions {
			drx, dry := dr[0], dr[1]
			if drx+di >= 0 && drx+di < r && dj+dry >= 0 && dj+dry < c && maze[drx+di][dry+dj] == 46 {
				maze[drx+di][dry+dj] = 43
				queue = append(queue, []int{drx + di, dj + dry, step + 1})
			}
		}
	}

	return -1
}
