package main

import (
	"fmt"
)

// A swap is defined as taking two distinct positions in an array and swapping the values in them.
// A circular array is defined as an array where we consider the first element and the last element to be adjacent.
// Given a binary circular array nums, return the minimum number of swaps required to group all 1's present in the array together at any location.

// Input: nums = [0,1,0,1,1,0,0]

// Output: 1
// Explanation: Here are a few of the ways to group all the 1's together:

// [0,0,1,1,1,0,0] using 1 swap.
// [0,1,1,1,0,0,0] using 1 swap.
// [1,1,0,0,0,0,1] using 2 swaps (using the circular property of the array).
// There is no way to group all 1's together with 0 swaps.
// Thus, the minimum number of swaps required is 1.

func minSwaps(nums []int) int {
	k := 0
	for _, num := range nums {
		if num == 1 {
			k++
		}
	}
	slidingWindows := 0

	// Calculate max 1 of first sliding
	for i := 0; i < k; i++ {
		slidingWindows += nums[i]
	}

	for i := k; i < len(nums); i++ {
		slidingWindows += nums[i-k] + nums[i]
	}

	return k - slidingWindows
}

func main() {
	fmt.Println(minSwaps([]int{0, 1, 0, 1, 1, 0, 0}))
	fmt.Println(minSwaps([]int{0, 1, 1, 1, 0, 0, 1, 1, 0}))

}
