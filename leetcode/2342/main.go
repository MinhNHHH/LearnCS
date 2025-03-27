package main

import (
	"fmt"
	"sort"
)

func sumDigits(num int) int {
	res := 0
	for num > 0 {
		res += num % 10
		num = num / 10
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maximumSum(nums []int) int {
	hashMap := map[int][]int{}
	for _, num := range nums {
		key := sumDigits(num)
		hashMap[key] = append(hashMap[key], num)
	}
	res := -1
	for _, arr := range hashMap {
		if len(arr) >= 2 {
			sort.Ints(arr)
			res = max(res, arr[len(arr)-1]+arr[len(arr)-2])
		}
	}

	return res
}

func main() {
	fmt.Println(maximumSum([]int{18, 43, 36, 13, 7}))
}
