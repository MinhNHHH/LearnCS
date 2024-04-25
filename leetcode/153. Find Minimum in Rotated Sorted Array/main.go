package main

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2
		if nums[0] <= nums[right] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}

func main() {
	// fmt.Println(findMin([]int{3, 1, 2}))
	fmt.Println(findMin([]int{4, 3, 1, 2}))
}
