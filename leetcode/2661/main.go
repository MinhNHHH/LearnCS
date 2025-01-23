package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func firstCompleteIndex(arr []int, mat [][]int) int {
	positionMatrix := map[int][]int{}
	rows := len(mat)
	cols := len(mat[0])
	rowCount := make([]int, rows)
	colCount := make([]int, cols)

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			positionMatrix[mat[r][c]] = []int{r, c}
		}
	}

	for r := 0; r < rows; r++ {
		rowCount[r] = cols
	}

	for c := 0; c < cols; c++ {
		colCount[c] = rows
	}

	for index, value := range arr {
		r, c := positionMatrix[value][0], positionMatrix[value][1]
		rowCount[r]--
		colCount[c]--

		if rowCount[r] == 0 || colCount[c] == 0 {
			return index
		}
	}
	return -1
}
