package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func firstMissingPositive(nums []int) int {
	set := map[int]bool{}

	for _, num := range nums {
		set[num] = true
	}

	for i := 1; i <= len(nums); i++ {
		if _, ok := set[i]; !ok {
			return i
		}
	}

	return len(set) + 1
}

func main() {
	fmt.Println(firstMissingPositive([]int{0, 1, 2}))

	fmt.Println(firstMissingPositive([]int{1}))

	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
}
