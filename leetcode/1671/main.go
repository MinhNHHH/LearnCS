package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minimumMountainRemovals(nums []int) int {

	inscres := make([]int, len(nums))
	decres := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				inscres[i] = max(inscres[i], inscres[j]+1)
			}
		}
	}

	for i := len(nums) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if nums[i] < nums[j] {
				decres[i] = max(decres[i], decres[j]+1)
			}
		}
	}
	fmt.Println(inscres)
	fmt.Println(decres)
	return 1
}

func main() {
	fmt.Println(minimumMountainRemovals([]int{2, 1, 1, 5, 6, 2, 3, 1}))
}
