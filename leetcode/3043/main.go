package main

import (
	"fmt"
	"strconv"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	hashSet1 := map[int]bool{}
	// hashSet2 := map[string]bool{}
	longestPrefix := 0

	for _, num := range arr1 {
		for num > 0 {
			hashSet1[num] = true
			num = num / 10
		}
	}

	for _, num := range arr2 {
		for num > 0 {
			if hashSet1[num] {
				longestPrefix = max(longestPrefix, num)
			}
			num = num / 10
		}
	}
	if longestPrefix == 0 {
		return 0
	}
	return len(strconv.Itoa(longestPrefix))
}

func main() {
	fmt.Println(longestCommonPrefix([]int{1, 2, 3}, []int{4}))
}
