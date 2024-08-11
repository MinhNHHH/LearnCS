package main

import "fmt"

func main() {
	rotate([][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}})
}

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func rotate(matrix [][]int) {
	row := len(matrix)
	col := len(matrix[0])

	for r := 0; r < row; r++ {
		for c := r; c < col; c++ {
			matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
		}
	}
	for _, m := range matrix {
		reverseSlice(m)
	}
	fmt.Println(matrix)

}
