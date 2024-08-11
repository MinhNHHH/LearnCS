package main

import "fmt"

// [1,2,3,4,5,6,7] 3 => [5,6,7,1,2,3,4]

//
func rotate(nums []int, k int) {
	n := len(nums)
	for k != 0 {
		temp := nums[n-1]
		for i := n - 1; i > 0; i-- {
			nums[i] = nums[i-1]
		}
		nums[0] = temp
		k--
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums)
}
