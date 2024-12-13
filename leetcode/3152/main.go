package main

import "fmt"

func isArraySpecial(nums []int, queries [][]int) []bool {
	n := len(nums)

	preComputePair := make([]bool, n)
	for i := 1; i < n; i++ {
		preComputePair[i] = (nums[i]%2 != nums[i-1]%2)
	}

	preFixSum := make([]int, n)
	for i := 1; i < n; i++ {
		preFixSum[i] = preFixSum[i-1]
		if !preComputePair[i] {
			preFixSum[i]++
		}
	}

	result := make([]bool, len(queries))
	for i, query := range queries {
		from, to := query[0], query[1]
		if from == to {
			result[i] = true
		} else {
			specialCount := preFixSum[to] - preFixSum[from]
			result[i] = (specialCount == 0)
		}
	}
	return result
}

func main() {
	fmt.Println(isArraySpecial([]int{3, 4, 1, 2, 6}, [][]int{{0, 4}}))
}
