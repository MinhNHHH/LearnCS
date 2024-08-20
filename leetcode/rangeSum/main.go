package main

import (
	"fmt"
	"sort"
)

func rangeSum(nums []int, n int, left int, right int) int {
	// newLen := n * (n+1) /2
	newArr := []int{}
	sum := 0
	i := 0
	for i < len(nums) {
		sumSubArr := 0
		for j := i; j < len(nums); j++ {
			sumSubArr += nums[j]
			newArr = append(newArr, sumSubArr)
		}
		i++
	}
	sort.Ints(newArr)
	// (Số Hạng Đầu + Số Hạng Cuối) x Số Số Hạng / 2
	sum = (newArr[left-1] + newArr[right-1]) * (right - left + 1) / 2
	return sum
}

func main() {
	fmt.Println(rangeSum([]int{1, 2, 3, 4}, 4, 1, 5))
}
