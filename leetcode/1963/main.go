package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func minSwaps(s string) int {
	n := len(s)
	imbalance := 0
	max_imbalace := 0
	for i := 0; i < n; i++ {
		if string(s[i]) == "]" {
			imbalance++
		} else {
			imbalance--
		}
		max_imbalace = max(max_imbalace, imbalance)
	}

	return (max_imbalace + 1) / 2
}

func main() {
	fmt.Println(minSwaps("[]]["))
}
