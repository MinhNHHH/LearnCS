package main

import (
	"fmt"
	"sort"
)

// Give an 2D integer gird size m x n. In one operation, you can add x to or subtract x from any element in the gird
// A uni-value gird is a gird where all the elements of it are equal
// Return the minimun number of operations to make the gird uni-value.

// The solution:
// Frist check all reminder in gird is the same if all is the same continue process else return -1
// Flatten 2D to 1D then find median
// compute ∣num−median∣/x∣num−median∣/x.

func flat(grids [][]int) []int {
	res := []int{}
	for _, grid := range grids {
		res = append(res, grid...)
	}
	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func minOperations(grid [][]int, x int) int {
	newArr := flat(grid)
	sort.Ints(newArr)
	reminder := newArr[0] % x
	for i := 1; i < len(newArr); i++ {
		temp := newArr[i] % x
		if temp != reminder {
			return -1
		}
		reminder = temp
	}
	mid := len(newArr) / 2
	if len(newArr)%2 == 0 {
		mid = len(newArr)/2 - 1
	}

	median := newArr[mid]
	total := 0
	for _, num := range newArr {
		total += abs(num - median)
	}
	return total / x
}

func main() {
	fmt.Println(minOperations([][]int{{2, 4}, {6, 8}}, 2))
	fmt.Println(minOperations([][]int{{1, 2}, {3, 4}}, 2))
}
