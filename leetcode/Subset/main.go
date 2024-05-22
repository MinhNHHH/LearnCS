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
	fmt.Println(subSetsBFS([]int{1, 2, 3}))
}

func subsetsDFS(nums []int) [][]int {
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

// Helper function to find the index of an element in nums
func indexOf(nums []int, target int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}
	return -1
}

func subSetsBFS(nums []int) [][]int {
	res := [][]int{}
	queue := [][]int{{}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		res = append(res, current)
		start := 0
		if len(current) > 0 {
			start = indexOf(nums, current[len(current)-1]) + 1
		}

		for i := start; i < len(nums); i++ {
			newSubset := append([]int{}, current...)
			newSubset = append(newSubset, nums[i])
			queue = append(queue, newSubset)
		}
		fmt.Println(queue)
	}
	return res
}
