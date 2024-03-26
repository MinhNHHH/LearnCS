package main

import "fmt"

func search(nums []int, target int, conditionBias bool) int {
	left, right := 0, len(nums)-1
	index := -1
	for left <= right {
		mid := left + (right-left)/2
		if target > nums[mid] {
			left = mid + 1
		} else if target < nums[mid] {
			right = mid - 1
		} else {
			index = mid
			if conditionBias {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return index
}

func searchRange(nums []int, target int) []int {
	firstPosition := search(nums, target, false)
	if firstPosition == -1 {
		return []int{-1, -1}
	}
	lastPosition := search(nums, target, true)
	return []int{firstPosition, lastPosition}
}

func main() {
	// fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	// fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 6))
	// fmt.Println(searchRange([]int{}, 2))
	fmt.Println(searchRange([]int{2, 2}, 2))
}
