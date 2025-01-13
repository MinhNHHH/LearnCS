package main

import "fmt"

func waysToSplitArray(nums []int) int {
	left := 0
	tempLeft := 0
	total := 0
	for _, value := range nums {
		total += value
	}
	res := 0
	for left < len(nums)-1 {
		tempLeft += nums[left]
		tempRight := total - tempLeft
		if tempLeft >= tempRight {
			res++
		}
		left++
	}

	return res
}

func main() {
	fmt.Println(waysToSplitArray([]int{10, 4, -8, 7}))
}
