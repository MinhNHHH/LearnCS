package main

import (
	"fmt"
)

func maxCount(banned []int, n int, maxSum int) int {
	mapBanned := make(map[int]bool)
	for _, num := range banned {
		mapBanned[num] = true
	}
	count := 0
	sums := 0
	// Iterate over numbers from 1 to n
	for i := 1; i <= n; i++ {
		// Skip banned numbers
		if mapBanned[i] {
			continue
		}

		// Check if adding the current number exceeds maxSum
		if sums+i > maxSum {
			break
		}

		// Add the current number and increment the count
		sums += i
		count++
	}
	return count

}

func main() {
	fmt.Println(maxCount([]int{1, 6, 5}, 5, 6))
	fmt.Println(maxCount([]int{11}, 7, 50))

}
