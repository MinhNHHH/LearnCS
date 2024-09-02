package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func maxDistance(arrays [][]int) int {
	minA := arrays[0][0]
	maxA := arrays[0][len(arrays[0])-1]
	res := 0
	for _, arr := range arrays {
		res = max(res, max(arr[len(arr)-1]-minA, maxA-arr[0]))
		minA = min(minA, arr[0])
		maxA = max(maxA, arr[len(arr)-1])
	}
	return res
}

func main() {
	fmt.Println(maxDistance([][]int{{1, 4}, {0, 5}}))
}
