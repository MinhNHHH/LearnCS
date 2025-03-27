package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func longestMonotonicSubarray(nums []int) int {
	increase := 1
	decrease := 1

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			increase++
			decrease = 1
		} else if nums[i] < nums[i+1] {
			decrease++
			increase = 1
		} else {
			increase = 1
			decrease = 1
		}

	}

	return max(increase, decrease)
}

func main() {
	fmt.Println(longestMonotonicSubarray([]int{1, 4, 3, 3, 2}))
}
