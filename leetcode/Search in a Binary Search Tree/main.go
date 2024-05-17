// You are given the root of a binary search tree (BST) and an integer val.

// Find the node in the BST that the node's value equals val and return the subtree rooted with that node. If such a node does not exist, return null.
// Input: root = [4,2,7,1,3], val = 2
// Output: [2,1,3]

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if val < root.Val {
		searchBST(root.Left, val)
	} else if val > root.Val {
		searchBST(root.Right, val)
	} else {
		return root
	}
	return root
}
