package main

import (
	"fmt"
	"slices"
	"sort"
)

func check(nums []int) bool {
	temp := make([]int, len(nums))
	copy(temp, nums)
	sort.Ints(temp)
	n := len(nums)

	for i := 0; i < n; i++ {
		formed := []int{nums[i]}
		for j := (i + 1) % n; j != i; {
			formed = append(formed, nums[j])
			j = (j + 1) % n
		}
		if slices.Equal(temp, formed) {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(check([]int{3, 4, 5, 1, 2}))
}
