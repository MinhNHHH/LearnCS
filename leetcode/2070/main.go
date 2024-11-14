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

func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	priceBeauty := make([][2]int, 0)
	maxBeauty := 0
	for _, item := range items {
		price, beauty := item[0], item[1]
		maxBeauty = max(maxBeauty, beauty)
		priceBeauty = append(priceBeauty, [2]int{price, maxBeauty})
	}

	result := make([]int, len(queries))
	for i, query := range queries {
		idx := sort.Search(len(priceBeauty), func(j int) bool {
			return priceBeauty[j][0] > query
		}) - 1
		if idx >= 0 {
			result[i] = priceBeauty[idx][1]
		} else {
			result[i] = 0
		}
	}
	return result
}

func main() {
	fmt.Println(maximumBeauty([][]int{{1, 2}, {3, 2}, {2, 4}, {5, 6}, {3, 5}}, []int{1, 2, 3, 4, 5, 6}))
	fmt.Println(maximumBeauty([][]int{{10, 1000}}, []int{5}))
}
