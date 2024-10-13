package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxWidthRamp(nums []int) int {
	maxWidth := 0
	monotonicStack := []int{0}

	for i := 1; i < len(nums); i++ {
		index := monotonicStack[len(monotonicStack)-1]
		if nums[i] <= nums[index] {
			monotonicStack = append(monotonicStack, i)
		}
	}

	// Traverse from right to left and find the maximum width ramp
	for i := len(nums) - 1; i >= 0; i-- {
		// Check if the current number forms a valid ramp
		for len(monotonicStack) > 0 && nums[monotonicStack[len(monotonicStack)-1]] <= nums[i] {
			maxWidth = max(maxWidth, i-monotonicStack[len(monotonicStack)-1])
			// Pop from the stack
			monotonicStack = monotonicStack[:len(monotonicStack)-1]
		}
	}
	return maxWidth
}

func main() {
	fmt.Println(maxWidthRamp([]int{2, 2, 1}))
}
