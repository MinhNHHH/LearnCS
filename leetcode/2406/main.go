package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func minGroups(intervals [][]int) int {
	start := []int{}
	end := []int{}

	for _, interval := range intervals {
		start = append(start, interval[0])
		end = append(end, interval[1])
	}

	sort.Ints(start)
	sort.Ints(end)

	i, j := 0, 0
	res := 0
	groups := 0
	for i < len(start) {
		if start[i] <= end[j] {
			i++
			groups++
		} else {
			j++
			groups--
		}
		res = max(res, groups)
	}
	return res
}

func main() {
	fmt.Println(minGroups([][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}}))

}
