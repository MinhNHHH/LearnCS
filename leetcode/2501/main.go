package main

import (
	"fmt"
	"sort"
)

func longestSquareStreak(nums []int) int {
	sort.Ints(nums)
	numMap := make(map[int]bool)

	for _, val := range nums {
		numMap[val] = true
	}

	longestSteak := 0

	for _, num := range nums {
		steak := 1
		curr := num
		for {
			next := curr * curr
			if numMap[next] {
				steak++
				curr = next
			} else {
				break
			}
		}
		if steak >= 2 && steak >= longestSteak {
			longestSteak = steak
		}
	}
	if longestSteak < 2 {
		return -1
	}
	return longestSteak
}

func main() {
	fmt.Println(longestSquareStreak([]int{4, 3, 6, 16, 8, 2}))
}
