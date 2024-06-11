// Given the root of a binary search tree and the lowest and highest boundaries as low and high, trim the tree so that all its elements lies in [low, high].
// Trimming the tree should not change the relative structure of the elements that will remain in the tree (i.e., any node's descendant should remain a descendant).
// It can be proven that there is a unique answer.

// Return the root of the trimmed binary search tree. Note that the root may change depending on the given bounds.

// Input: root = [1,0,2], low = 1, high = 2
// Output: [1,null,2]

// Input: root = [3,0,4,null,2,null,null,1], low = 1, high = 3
// Output: [3,2,null,1]

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {}

func minValue(node *TreeNode) *TreeNode {
	cur := node
	for cur.Left != nil {
		cur = cur.Left
	}
	return cur
}

func deleteNode(node *TreeNode, key int) *TreeNode {
	if node == nil {
		return node
	}

	if key < node.Val {
		node.Left = deleteNode(node.Left, key)
	} else if key > node.Val {
		node.Right = deleteNode(node.Right, key)
	} else {
		if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		}
		temp := minValue(node.Right)
		node.Val = temp.Val
		node.Right = deleteNode(node.Right, temp.Val)
	}
	return node
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	if root.Val < low {
		return root.Right
	}
	if root.Val > high {
		return root.Left
	}

	return root
}
