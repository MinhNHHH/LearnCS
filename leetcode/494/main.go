package main

import "fmt"

func findTargetSumWays(nums []int, target int) int {
	totalSum := 0

	for _, num := range nums {
		totalSum += num
	}

	if abs(target) > totalSum || (target+totalSum)%2 != 0 {
		return 0
	}

	sum := (totalSum + target) / 2
	dp := make([]int, sum+1)
	dp[0] = 1

	for _, num := range nums {
		for j := sum; j >= num; j-- {
			dp[j] += dp[j-num]
		}
	}

	return dp[sum]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(findTargetSumWays([]int{1, 1, 1, 1, 1}, 3))
}
