package main

import "fmt"

// 30 25 16 7   8 9
// 29 24 15 6   1 2
// 28 23 14 5   4 3
// 27 22 13 12 11 10
// 26 21 20 19 18 17

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	res := [][]int{{rStart, cStart}}
	movement := 1
	directions := [][]int{
		{0, 1},  // move right
		{1, 0},  // move down
		{0, -1}, // move up
		{-1, 0}, // move left
	}

	for len(res) < rows*cols {
		for k, i := range directions {
			dr, dc := i[0], i[1]
			for step := 0; step < movement; step++ {
				rStart += dr
				cStart += dc
				if (rStart >= 0 && rStart < rows) && cStart >= 0 && cStart < cols {
					res = append(res, []int{rStart, cStart})
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
	fmt.Println(spiralMatrixIII(5, 6, 1, 4))
}
