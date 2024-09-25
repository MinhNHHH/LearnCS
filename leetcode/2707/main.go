package main

import "fmt"

func minExtraChar(s string, dictionary []string) int {
	dp := make([]int, len(s))
	dp[0] = 0

	hash := map[string]bool{}
	for _, str := range dictionary {
		hash[str] = true
	}

	j := 0
	i := 0
	for j < len(s) && i < len(s) {
		fmt.Println(hash[string(s[j:i])], i, j)
		if _, ok := hash[string(s[j:i])]; !ok {
			dp[i] = dp[j] + len(string(s[j:i]))
		} else {
			dp[i] = dp[j] + 0
			j = i
		}
		i++
	}
	fmt.Println(dp)
	return dp[j]
}

func main() {
	fmt.Println(minExtraChar("sayhelloworld", []string{"hello", "world"}))
}
