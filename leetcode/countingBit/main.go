package main

import "fmt"

//Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n),
// ans[i] is the number of 1's in the binary representation of i.

// Input: n = 2
// Output: [0,1,1]
// Explanation:
// 0 --> 0
// 1 --> 1
// 2 --> 10
// Input: n = 5
// Output: [0,1,1,2,1,2]
// Explanation:
// 0 --> 0
// 1 --> 1
// 2 --> 10
// 3 --> 11
// 4 --> 100
// 5 --> 101

func countBits(n int) []int {

	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = dp[i>>1] + (i & 1)
	}
	return dp
}
func main() {
	fmt.Println(countBits(2))
	fmt.Println(countBits(5))
}
