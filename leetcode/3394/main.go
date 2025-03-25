package main

import (
	"fmt"
	"sort"
)

func checkValidCuts(n int, rectangles [][]int) bool {
	xInterval := [][]int{}
	yInterval := [][]int{}

	for _, rectangle := range rectangles {
		xInterval = append(xInterval, []int{rectangle[0], rectangle[2]})
		yInterval = append(yInterval, []int{rectangle[1], rectangle[3]})
	}

	return checkInterval(xInterval) || checkInterval(yInterval)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func checkInterval(intervals [][]int) bool {
	count := 0
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	maxEnd := intervals[0][1]
	for _, interval := range intervals {
		if maxEnd <= interval[0] {
			count++
		}
		maxEnd = max(maxEnd, interval[1])
	}

	return count >= 2
}

func main() {
	fmt.Println(checkValidCuts(5, [][]int{{1, 0, 5, 2}, {0, 2, 2, 4}, {3, 2, 5, 3}, {0, 4, 4, 5}}))
}
