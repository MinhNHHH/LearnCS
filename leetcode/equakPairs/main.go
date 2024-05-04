// Given a 0-indexed n x n integer matrix grid, return the number of pairs (ri, cj) such that row ri and column cj are equal.

// A row and column pair is considered equal if they contain the same elements in the same order (i.e., an equal array).

// Input: grid = [[3,2,1],[1,7,6],[2,7,7]]
// Output: 1
// Explanation: There is 1 equal row and column pair:
// - (Row 2, Column 1): [2,7,7]

// Input: grid = [[3,1,2,2],[1,4,4,5],[2,4,2,2],[2,4,2,2]]
// Output: 3
// Explanation: There are 3 equal row and column pairs:
// - (Row 0, Column 0): [3,1,2,2]
// - (Row 2, Column 2): [2,4,2,2]
// - (Row 3, Column 2): [2,4,2,2]

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(equalPairs([][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}))
	fmt.Println(equalPairs([][]int{{13, 13}, {13, 13}}))
}

func equalPairs(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	res := 0
	hashMap := map[string]int{}

	for r := 0; r < row; r++ {
		temp := ""
		for c := 0; c < col; c++ {
			result := strconv.Itoa(grid[r][c])
			temp += result + ","
		}
		hashMap[temp]++
	}

	for c := 0; c < col; c++ {
		temp := ""
		for r := 0; r < row; r++ {
			result := strconv.Itoa(grid[r][c])
			temp += result + ","
		}
		if value, ok := hashMap[temp]; ok {
			res += value
		}
	}

	return res
}
