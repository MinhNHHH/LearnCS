package main

import (
	"fmt"
	"math"
)

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
func findMaxAverage(nums []int, k int) float64 {
	maxAver := float64(math.MinInt64)
	count := 0
	sum := 0
	for index, _ := range nums {
		for count < k && index+count < len(nums) {
			sum += nums[index+count]
			count++
		}

		if count == k {
			maxAver = max(maxAver, float64(sum)/float64(k))
		}
		count--
		sum -= nums[index]
	}
	return maxAver
}

func main() {
	test1 := []int{1, 12, -5, -6, 50, 3}
	testk1 := 4
	fmt.Println(findMaxAverage(test1, testk1))
}
