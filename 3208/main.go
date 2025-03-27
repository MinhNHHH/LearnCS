package main

import (
	"fmt"
)

func main() {
	fmt.Println(numberOfAlternatingGroups([]int{0, 1, 0, 0, 1, 0, 1}, 6))
}

func numberOfAlternatingGroups(colors []int, k int) int {
	colors = append(colors, colors[:k-1]...)
	res := 0
	left := 0

	for i := 0; i < len(colors); i++ {
		if i > 0 && colors[i] == colors[i-1] {
			left = i
		}
		fmt.Println(i, left)
		if i-left+1 >= k {
			res++
		}
	}

	return res
}
