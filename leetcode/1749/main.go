package main

import "fmt"

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxAbsoluteSum(nums []int) int {
	temp := make([]int, len(nums))
	copy(temp, nums)
	maxSum := temp[0]
	for i := 1; i < len(nums); i++ {
		if temp[i-1] > 0 {
			temp[i] = temp[i-1] + temp[i]
		}
		maxSum = max(maxSum, temp[i])
	}

	minSum := temp[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] < 0 {
			nums[i] = nums[i-1] + nums[i]
		}
		minSum = min(minSum, nums[i])
	}
	return max(abs(minSum), abs(maxSum))
}

func main() {
	fmt.Println(maxAbsoluteSum([]int{2, -5, 1, -4, 3, -2}))
}
