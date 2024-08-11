package main

// You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

// Find two lines that together with the x-axis form a container, such that the container contains the most water.

// Return the maximum amount of water a container can store.

// Notice that you may not slant the container.

import (
	"fmt"
)

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

// 1, 8, 6, 2, 5, 4, 8, 3, 7

func maxArea(height []int) int {
	left, right, maxArea := 0, len(height)-1, 0
	for left < right {
		maxArea = max(maxArea, min(height[left], height[right])*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func main() {
	fmt.Println("test1", maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println("test1", maxArea([]int{1, 1}))
}
