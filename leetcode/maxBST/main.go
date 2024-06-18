// Given a binary tree root, return the maximum sum of all keys of any sub-tree which is also a Binary Search Tree (BST).

// Assume a BST is defined as follows:

// The left subtree of a node contains only nodes with keys less than the node's key.
// The right subtree of a node contains only nodes with keys greater than the node's key.
// Both the left and right subtrees must also be binary search trees.

// Input: root = [1,4,3,2,4,2,5,null,null,null,null,null,null,4,6]
// Output: 20
// Explanation: Maximum sum in a valid Binary search tree is obtained in root node with key equal to 3.

// Input: root = [4,3,null,1,2]
// Output: 2
// Explanation: Maximum sum in a valid Binary search tree is obtained in a single root node with key equal to 2.

// Input: root = [-4,-2,-5]
// Output: 0
// Explanation: All values are negatives. Return an empty BST.
package main

type TreeNode struct {
	Val   int
	Right *TreeNode
	Left  *TreeNode
}

func main() {}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func validateAndSum(node *TreeNode) (bool, int, int, int) {
	if node == nil {
		return true, 0, 0, 0
	}

	isLeftBST, totalSumLeft, minLeft, maxLeft := validateAndSum(node.Left)
	isRightBST, totalSumRight, minRight, maxRight := validateAndSum(node.Right)

	if isLeftBST && isRightBST && maxLeft < node.Val && node.Val < minRight {
		total := node.Val + totalSumLeft + totalSumRight
		return true, total, min(minLeft, node.Val), max(maxRight, node.Val)
	}
	return false, 0, 0, 0
}

func maxSumBST(root *TreeNode) int {
	res := 0

	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		isBst, sumTree, _, _ := validateAndSum(root)
		if isBst {
			res = max(sumTree, res)
		}
		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(root)
	return res
}
