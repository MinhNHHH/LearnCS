// Given an integer array nums of unique elements, return all possible
// subsets
// (the power set).

// The solution set must not contain duplicate subsets. Return the solution in any order.

// Input: nums = [1,2,3]
// Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

// Input: nums = [0]
// Output: [[],[0]]

package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	res := [][]int{}
	var dfs func(start int, path []int)

	dfs = func(start int, path []int) {
		// Append a copy of the current path to the result
		pathCopy := make([]int, len(path))
		copy(pathCopy, path)
		res = append(res, pathCopy)

		// Explore further elements to generate new subsets
		for i := start; i < len(nums); i++ {
			// Include nums[i] in the current subset
			path = append(path, nums[i])
			// Move on to the next element
			dfs(i+1, path)
			// Backtrack to explore subsets without nums[i]
			path = path[:len(path)-1]
		}
	}

	// Start the DFS with an empty path
	dfs(0, []int{})
	return res
}
