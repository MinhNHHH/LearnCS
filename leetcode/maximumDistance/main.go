package main

import (
	"fmt"
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxDistance(arr [][]int) int {
	minA := math.MaxInt
	maxA := math.MinInt
	for _, a := range arr {
		minA = min(minA, a[0])
		maxA = max(maxA, a[len(a)-1])
	}
	return maxA - minA
}
func main() {
	fmt.Println(maxDistance([][]int{{1, 2, 3}, {4, 5}, {1, 2, 3}}))
}
