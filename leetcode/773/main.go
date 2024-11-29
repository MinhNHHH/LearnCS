package main

import (
	"fmt"
	"strings"
)

func boardToString(board [][]int) string {
	var sb strings.Builder

	for r := 0; r < len(board); r++ {
		for c := 0; c < len(board[0]); c++ {
			sb.WriteByte(byte(board[r][c]) + '0')
		}
	}
	return sb.String()
}

func slidingPuzzle(board [][]int) int {
	target := "123450"
	newBoards := boardToString(board)
	visited := map[string]bool{}
	queue := []string{newBoards}
	directions := [][2]int{
		{-1, 0}, // Up
		{1, 0},  // Down
		{0, -1}, // Left
		{0, 1},  // Right
	}
	steps := 1
	for len(queue) > 0 {

		for i := 0; i < len(queue); i++ {
			current := queue[0]
			queue = queue[1:]
			if current == target {
				return steps
			}
			spaceIndex := strings.Index(current, "0")
			row, col := spaceIndex/3, spaceIndex%3

			for _, dir := range directions {
				dx, dy := row+dir[0], col+dir[1]

				if dx >= 0 && dx < 2 && dy >= 0 && dy < 3 {
					newIndex := dx*3 + dy
					newState := []rune(current)
					newState[spaceIndex], newState[newIndex] = newState[newIndex], newState[spaceIndex]
					nextState := string(newState)
					if !visited[nextState] {
						visited[nextState] = true
						queue = append(queue, nextState)
					}
				}
			}
		}
		steps++
	}
	return -1
}

func main() {
	fmt.Println(slidingPuzzle([][]int{{1, 2, 3}, {4, 0, 5}}))
	fmt.Println(slidingPuzzle([][]int{{1, 2, 3}, {5, 4, 0}}))
	fmt.Println(slidingPuzzle([][]int{{4, 1, 2}, {5, 0, 3}}))
	fmt.Println(slidingPuzzle([][]int{{3, 2, 4}, {1, 5, 0}}))
}
