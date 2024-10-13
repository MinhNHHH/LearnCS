package main

import (
	"fmt"
)

func minSubarray(nums []int, p int) int {
	// count := 0
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}

	// target := nums[len(nums)-1] % p
	// subarr := map[int]int{}
	// for index, value := range nums {
	// 	div := value % p
	// }
	// fmt.Println(subarr)
	return 1
}

func main() {
	fmt.Println(minSubarray([]int{3, 1, 4, 2}, 6))
}
