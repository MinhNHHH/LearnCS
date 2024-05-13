// Consider all the leaves of a binary tree, from left to right order, the values of those leaves form a leaf value sequence.
// For example, in the given tree above, the leaf value sequence is (6, 7, 4, 9, 8).

// Two binary trees are considered leaf-similar if their leaf value sequence is the same.

// Return true if and only if the two given trees with head nodes root1 and root2 are leaf-similar.

// Input: root1 = [3,5,1,6,2,9,8,null,null,7,4], root2 = [3,5,1,6,7,4,2,null,null,null,null,null,null,9,8]
// Output: true

//	       		3																			3
//			 5			   		1													5            1
//	  6  	2		 		9    8                    6   7        4    2
//	    7     4																							9   8
package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Function to create a new TreeNode
func NewTreeNode(val int) *TreeNode {
	return &TreeNode{Val: val, Left: nil, Right: nil}
}

// Function to insert a new node into the binary tree
func (root *TreeNode) Insert(val int) *TreeNode {
	if root == nil {
		return NewTreeNode(val)
	}

	if val < root.Val {
		root.Left = root.Left.Insert(val)
	} else {
		root.Right = root.Right.Insert(val)
	}

	return root
}

// Function to print the binary tree in inorder traversal
func (root *TreeNode) Inorder() {
	if root != nil {
		root.Left.Inorder()
		fmt.Printf("%d ", root.Val)
		root.Right.Inorder()
	}
}

func main() {
	// Initializing a binary tree with root node value 5
	root := NewTreeNode(5)
	root.Insert(3)
	root.Insert(7)
	root.Insert(2)
	root.Insert(4)
	root.Insert(6)
	root.Insert(8)

	root2 := NewTreeNode(5)
	root2.Insert(3)
	root2.Insert(7)
	root2.Insert(2)
	root2.Insert(4)
	root2.Insert(6)
	root2.Insert(8)

	leafSimilar(root, root2)
}

func dfs(root *TreeNode, leafInt []int) []int {
	if root == nil {
		return leafInt
	}
	if root.Left == nil && root.Right == nil {
		leafInt = append(leafInt, root.Val)
	}
	leafInt = dfs(root.Left, leafInt)
	leafInt = dfs(root.Right, leafInt)
	return leafInt
}

func listsEqual(list1, list2 []int) bool {
	// Check if lengths are equal
	if len(list1) != len(list2) {
		return false
	}

	// Iterate through elements and compare
	for i := 0; i < len(list1); i++ {
		if list1[i] != list2[i] {
			return false
		}
	}

	return true
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaf1 := dfs(root1, []int{})
	leaf2 := dfs(root2, []int{})

	return listsEqual(leaf1, leaf2)
}
