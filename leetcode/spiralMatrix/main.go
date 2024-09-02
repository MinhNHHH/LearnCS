package main

import "fmt"

func spiralOrder(matrix [][]int) []int {
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	movement := 1
	r, c := 0, 0
	res := []int{matrix[r][c]}
	rows := len(matrix)
	cols := len(matrix[0])
	for len(res) < rows*cols {
		for k, dir := range directions {
			dr, dc := dir[0], dir[1]
			for step := 0; step < movement; step++ {
				r += dr
				c += dc
				if r >= 0 && r < rows && c >= 0 && c < cols {
					res = append(res, matrix[r][c])
				}
			}
			if k == 1 || k == 3 {
				movement += 1
			}

		}
	}
	return res
}

func main() {
	fmt.Println(spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
}
