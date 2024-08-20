package main

import (
	"fmt"
	"sort"
)

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool { return citations[i] > citations[j] })
	h, n := citations[0], len(citations)

	for i := 0; i < n; i++ {
		if citations[i] >= i {
			h = i
		}
	}
	return h + 1

}

func main() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
}
