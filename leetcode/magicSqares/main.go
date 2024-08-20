package main

import "fmt"

//A 3 x 3 magic square is a 3 x 3 grid filled with distinct numbers from 1 to 9 such that each row, column, and both diagonals all have the same sum.

//Given a row x col grid of integers, how many 3 x 3 contiguous magic square subgrids are there?

//Note: while a magic square can only contain numbers from 1 to 9, grid may contain numbers up to 15.

// 4 3 8 4
// 9 5 1 9
// 2 7 6 2

// 5 5 5
// 5 5 5
// 5 5 5
func sum(arr []int) int {
	sums := 0
	for _, value := range arr {
		sums += value
	}
	return sums
}

func isMagicSquare(matrix [][]int, i, j int) int {
	if i+3 > len(matrix) || j+3 > len(matrix[0]) {
		return 0
	}
	set := map[int]bool{}
	rowSums := make([]int, 3)
	colSums := make([]int, 3)
	diagonalSumsa := matrix[i][j] + matrix[i+1][j+1] + matrix[i+2][j+2]
	diagonalSumsb := matrix[i][j+2] + matrix[i+1][j+1] + matrix[i+2][j]
	for r := i; r < i+3; r++ {
		for c := j; c < j+3; c++ {
			value := matrix[r][c]
			if value >= 1 && value <= 9 && !set[value] {
				set[value] = true
				rowSums[r-i] += value
				colSums[c-j] += value
			} else {
				return 0
			}
		}
	}
	if diagonalSumsa != diagonalSumsb {
		return 0
	}
	for _, value := range rowSums {
		if diagonalSumsa != value || diagonalSumsb != value {
			return 0
		}
	}
	for _, value := range colSums {
		if diagonalSumsa != value || diagonalSumsb != value {
			return 0
		}
	}

	return 1
}
func numMagicSquaresInside(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	res := 0
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			res += isMagicSquare(grid, r, c)
		}
	}
	return res
}

// 4 7 8
// 9 5 1
// 2 3 6
func main() {
	fmt.Println(numMagicSquaresInside([][]int{{4, 3, 8, 4}, {9, 5, 1, 9}, {2, 7, 6, 2}}))
	fmt.Println(numMagicSquaresInside([][]int{{5, 5, 5}, {5, 5, 5}, {5, 5, 5}}))
	fmt.Println(numMagicSquaresInside([][]int{{7, 0, 5}, {2, 4, 6}, {3, 8, 1}}))

}
