package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maximumSubarraySum(nums []int, k int) int64 {
	// Initialize variables
	visited := make(map[int]bool)
	start, currentSum, maxSum := 0, 0, 0

	for end := 0; end < len(nums); end++ {
		// If the current number is already in the window, shrink the window from the left
		for visited[nums[end]] {
			// Remove nums[start] from the visited map and subtract it from the current sum
			visited[nums[start]] = false
			currentSum -= nums[start]
			start++
		}

		// Add the current number to the window
		visited[nums[end]] = true
		currentSum += nums[end]

		// If the window size equals k, update the maximum sum
		if end-start+1 == k {
			if currentSum > maxSum {
				maxSum = currentSum
			}

			// Slide the window by removing the leftmost element
			visited[nums[start]] = false
			currentSum -= nums[start]
			start++
		}
	}

	return int64(maxSum)
}
func main() {
	fmt.Println(maximumSubarraySum([]int{1, 5, 4, 2, 9, 9, 9}, 3))
	fmt.Println(maximumSubarraySum([]int{9, 9, 9, 1, 2, 3}, 3))
}
