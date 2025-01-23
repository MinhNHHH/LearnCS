package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxScoreSightseeingPair(values []int) int {
	maxRes := 0
	maxVal := values[0]

	for i := 1; i < len(values); i++ {
		maxRes = max(maxRes, maxVal+values[i]-i)
		maxVal = max(maxVal, values[i]+i)
	}
	return maxRes
}

func main() {
	fmt.Println(maxScoreSightseeingPair([]int{8, 1, 5, 2, 6}))
}
