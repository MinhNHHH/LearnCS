package main

import "fmt"

func highestPeak(isWater [][]int) [][]int {
	heightMatrix := make([][]int, len(isWater))
	for i := range isWater {
		heightMatrix[i] = make([]int, len(isWater[i]))
		for j := range heightMatrix[i] {
			heightMatrix[i][j] = -1
		}
	}

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	queue := [][]int{}

	for r := 0; r < len(isWater); r++ {
		for c := 0; c < len(isWater[0]); c++ {
			if isWater[r][c] == 1 {
				queue = append(queue, []int{r, c})
				heightMatrix[r][c] = 0
			}
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, direct := range directions {
			newR, newC := (current[0] + direct[0]), (current[1] + direct[1])
			if newR >= 0 && newR < len(heightMatrix) && newC >= 0 && newC < len(heightMatrix[0]) && heightMatrix[newR][newC] == -1 {
				heightMatrix[newR][newC] = heightMatrix[current[0]][current[1]] + 1
				queue = append(queue, []int{newR, newC})
			}
		}
	}

	return heightMatrix
}

func main() {
	fmt.Println(highestPeak([][]int{{0, 1}, {0, 0}}))
}
