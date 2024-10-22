package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	maxLevel := []int{}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		sums := 0
		level := len(stack)
		for i := 0; i < level; i++ {
			node := stack[0]
			stack = stack[1:]
			sums += node.Val
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
		}
		maxLevel = append(maxLevel, sums)
	}
	sort.Slice(maxLevel, func(i, j int) bool {
		return maxLevel[i] > maxLevel[j]
	})

	if len(maxLevel) < k {
		return int64(maxLevel[0])
	}
	return int64(maxLevel[k-1])
}

// insertLevelOrder builds a binary tree from an array.
func insertLevelOrder(arr []int, root *TreeNode, i int) *TreeNode {
	// Base case for recursion
	if i < len(arr) {
		temp := &TreeNode{Val: arr[i]}
		root = temp

		// Insert left child
		root.Left = insertLevelOrder(arr, root.Left, 2*i+1)

		// Insert right child
		root.Right = insertLevelOrder(arr, root.Right, 2*i+2)
	}
	return root
}

// inorderTraversal performs inorder traversal of the binary tree.
func inorderTraversal(root *TreeNode) {
	if root != nil {
		inorderTraversal(root.Left)
		fmt.Print(root.Val, " ")
		inorderTraversal(root.Right)
	}
}

func main() {
	var root *TreeNode
	arr := []int{1, 2, 3}
	root = insertLevelOrder(arr, root, 0)
	fmt.Println(kthLargestLevelSum(root, 2))
}
