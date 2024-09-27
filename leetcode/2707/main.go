package main

import (
	"fmt"
	"math"
)

func minExtraChar(s string, dictionary []string) int {
	dp := make([]int, len(s)+1)
	dp[0] = 0
	hash := map[string]bool{}
	for _, str := range dictionary {
		hash[str] = true
	}

	for i := 1; i <= len(s); i++ {
		dp[i] = dp[i-1] + 1
		for l := i; l > 0; l-- {
			if hash[string(s[i-l:i])] {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-l])))
			}
		}
	}
	return dp[len(dp)-1]
}

func main() {
	fmt.Println(minExtraChar("sayhelloworld", []string{"hello", "world"}))
}
