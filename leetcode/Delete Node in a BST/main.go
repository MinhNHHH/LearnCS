// Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.

// Basically, the deletion can be divided into two stages:

// Search for a node to remove.
// If the node is found, delete the node.

// Input: root = [5,3,6,2,4,null,7], key = 3
// Output: [5,4,6,2,null,null,7]
// Explanation: Given key to delete is 3. So we find the node with value 3 and delete it.
// One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
// Please notice that another valid answer is [5,2,6,null,4,null,7] and it's also accepted.

// Input: root = [5,3,6,2,4,null,7], key = 0
// Output: [5,3,6,2,4,null,7]
// Explanation: The tree does not contain a node with value = 0.

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertLevelOrder(arr []interface{}, root *TreeNode, i int, n int) *TreeNode {
	if i < n {
		var temp *TreeNode
		if arr[i] != nil {
			temp = &TreeNode{Val: arr[i].(int)}
			root = temp

			// insert left child
			root.Left = insertLevelOrder(arr, root.Left, 2*i+1, n)

			// insert right child
			root.Right = insertLevelOrder(arr, root.Right, 2*i+2, n)
		}
	}
	return root
}

// inorder prints the inorder traversal of the binary tree
func inorder(root *TreeNode) {
	if root != nil {
		inorder(root.Left)
		fmt.Printf("%d ", root.Val)
		inorder(root.Right)
	}
}

func main() {
	arr := []interface{}{5, 3, 6, 2, 4, nil, 7}
	n := len(arr)
	var root *TreeNode
	root = insertLevelOrder(arr, root, 0, n)
	deleteNode(root, 3)
	inorder(root)
}

func minNode(root *TreeNode) *TreeNode {
	// In a Binary Search Tree (BST), the smallest node is the leftmost node.
	current := root

	for current != nil && current.Left != nil {
		current = current.Left
	}
	return current
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else {
		// The node is a leaf (has no children) or The node has only one child
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		}
		// Find the node’s in-order successor (the smallest node in its right subtree) or in-order predecessor (the largest node in its left subtree).
		temp := minNode(root.Right)
		root.Val = temp.Val

		root.Right = deleteNode(root.Right, temp.Val)
	}
	return root
}
