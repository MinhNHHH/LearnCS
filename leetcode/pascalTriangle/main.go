package main

import "fmt"

// Given an integer numRows, return the first numRows of Pascal's triangle.
// In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:
// Input: numRows = 5
// Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

func generate(numRows int) [][]int {
	dp := make([][]int, numRows)

	for index := range dp {
		row := make([]int, index+1)
		row[0], row[index] = 1, 1
		for i := 1; i < index; i++ {
			row[i] = dp[index-1][i] + dp[index-1][i-1]
		}
		dp[index] = row
	}

	return dp
}

func main() {
	fmt.Println(generate(5))
}
