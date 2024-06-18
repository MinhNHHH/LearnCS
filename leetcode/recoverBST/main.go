package main

import "math"

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var dfs func(root *TreeNode)
	var first *TreeNode
	var second *TreeNode
	prev := &TreeNode{Val: math.MinInt}

	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)

		if first == nil && prev.Val > root.Val {
			first = prev
		}
		if second == nil && prev.Val > root.Val {
			second = prev
		}

		dfs(root.Right)
	}

	dfs(root)
	temp := first
	first.Val = second.Val
	second.Val = temp.Val

}
