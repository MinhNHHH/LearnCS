package main

import "fmt"

func toZero(num int) int {
	if num < 0 {
		return 0
	}
	return num
}

func minZeroArray(nums []int, queries [][]int) int {
	totalNums := 0
	res := 0
	for _, val := range nums {
		totalNums += val
	}
	if totalNums == 0 {
		return 0
	}
	for _, query := range queries {
		l, r, val := query[0], query[1], query[2]
		for i := l; i <= r; i++ {
			totalNums -= toZero(nums[i] - val)
			nums[i] = toZero(nums[i] - val)
			fmt.Println(nums[i], totalNums)
		}
		res++
		if totalNums == 0 {
			return res
		}
	}

	return -1
}

func main() {
	fmt.Println(minZeroArray([]int{2, 0, 2}, [][]int{{0, 2, 1}, {0, 2, 1}, {1, 1, 3}}))
	// fmt.Println(minZeroArray([]int{0}, [][]int{{0, 0, 2}, {0, 0, 4}, {0, 0, 4}, {0, 0, 3}, {0, 0, 5}}))
}
