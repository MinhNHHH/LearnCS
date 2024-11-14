package main

import (
	"fmt"
	"sort"
)

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	res := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		left := i + 1
		right := n - 1

		// Find the smallest j such that nums[i] + nums[j] >= lower
		for left <= right {
			mid := (left + right) / 2
			if nums[i]+nums[mid] >= lower {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		lowerBound := left

		// Reset right pointer to find the largest j such that nums[i] + nums[j] <= upper
		left = i + 1
		right = n - 1
		for left <= right {
			mid := (left + right) / 2
			if nums[i]+nums[mid] <= upper {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		upperBound := right

		// Count pairs within bounds
		if lowerBound <= upperBound {
			res += upperBound - lowerBound + 1
		}
	}

	return int64(res)
}

func main() {
	fmt.Println(countFairPairs([]int{0, 1, 7, 4, 4, 5}, 3, 6))
}
