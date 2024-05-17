// Given the root of a binary tree, the level of its root is 1, the level of its children is 2, and so on.

// Return the smallest level x such that the sum of all the values of nodes at level x is maximal.

// Input: root = [1,7,0,7,-8,null,null]
// Output: 2
// Explanation:
// Level 1 sum = 1.
// Level 2 sum = 7 + 0 = 7.
// Level 3 sum = 7 + -8 = -1.
// So we return the level with the maximum sum which is level 2.

package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

func bfs(root *TreeNode) int {
	stack := []*TreeNode{root}
	maxSum := math.MinInt64
	level := 0
	res := 0

	for len(stack) > 0 {
		sum := 0
		level += 1
		subStack := []*TreeNode{}
		for _, node := range stack {
			sum += node.Val
			if node.Left != nil {
				subStack = append(subStack, node.Left)
			}
			if node.Right != nil {
				subStack = append(subStack, node.Right)
			}
		}
		stack = subStack
		if sum > maxSum {
			maxSum = sum
			res = level
		}
	}
	return res
}

func maxLevelSum(root *TreeNode) int {
	return bfs(root)
}
