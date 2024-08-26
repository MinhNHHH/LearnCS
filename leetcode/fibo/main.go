package main

import "fmt"

func fibo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1

	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
		fmt.Println("asdasd", dp)
	}
	return dp[n]
}

func main() {
	fmt.Println(fibo(3))
}
