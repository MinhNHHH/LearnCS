package main

import "fmt"

//You are given an integer array cost where cost[i] is the cost of ith step on a staircase. Once you pay the cost, you can either climb one or two steps.

//You can either start from the step with index 0, or the step with index 1.

//Return the minimum cost to reach the top of the floor.

//Input: cost = [10,15,20]
//Output: 15
//Explanation: You will start at index 1.
//- Pay 15 and climb two steps to reach the top.
//The total cost is 15.

//Input: cost = [1,100,1,1,1,100,1,1,100,1]
//Output: 6
//Explanation: You will start at index 0.
//- Pay 1 and climb two steps to reach index 2.
//- Pay 1 and climb two steps to reach index 4.
//- Pay 1 and climb two steps to reach index 6.
//- Pay 1 and climb one step to reach index 7.
//- Pay 1 and climb two steps to reach index 9.
//- Pay 1 and climb one step to reach the top.
//The total cost is 6.

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	} else if len(cost) == 1 {
		return cost[0]
	}

	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]
	for index := 2; index < len(cost); index++ {
		dp[index] = cost[index] + min(dp[index-1], dp[index-2])
	}

	return min(dp[len(cost)-1], dp[len(cost)-2])
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
}
