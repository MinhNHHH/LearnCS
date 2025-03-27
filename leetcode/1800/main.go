package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxAscendingSum(nums []int) int {
	res := nums[0]
	temp := nums[0]
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			temp += nums[i+1]
		} else {
			temp = nums[i+1]
		}
		res = max(res, temp)
	}

	return res
}

func main() {
	fmt.Println(maxAscendingSum([]int{10, 20, 30, 5, 10, 50}))
}
