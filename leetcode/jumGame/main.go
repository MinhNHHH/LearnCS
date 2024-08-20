package main

import "fmt"

// Input: nums = [2,3,1,1,4]
// Output: true
// Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.

func canJump(nums []int) bool {
	loopback := nums[0]
	for i := 1; i < len(nums); i++ {
		if loopback < i {
			return false
		} else {
			loopback = i + nums[i]
		}
	}
	return true
}

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
	fmt.Println(canJump([]int{1, 3, 2}))
}
