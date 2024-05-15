// Given the root of a binary tree and an integer targetSum, return the number of paths where the sum of the values along the path equals targetSum.

// The path does not need to start or end at the root or a leaf, but it must go downwards (i.e., traveling only from parent nodes to child nodes).

// Input: root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
// Output: 3
// Explanation: The paths that sum to 8 are shown.

// Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// Output: 3

package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Function to create a new TreeNode
// insertLevelOrder inserts nodes into the binary tree level by level
func insertLevelOrder(arr []interface{}, root *TreeNode, i int) *TreeNode {
	// Base case for recursion
	if i < len(arr) {
		var temp *TreeNode

		// Only insert non-null values
		if arr[i] != nil {
			temp = &TreeNode{Val: arr[i].(int)}
			root = temp

			// Insert left child
			root.Left = insertLevelOrder(arr, root.Left, 2*i+1)

			// Insert right child
			root.Right = insertLevelOrder(arr, root.Right, 2*i+2)
		}
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
	arr := []interface{}{10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1}
	var root *TreeNode
	root = insertLevelOrder(arr, root, 0)
	fmt.Println(pathSum(root, 8))
}

func pathSum(root *TreeNode, targetSum int) int {
	var dfs func(root *TreeNode, targetSum int, pathSum []int)
	count := 0
	dfs = func(root *TreeNode, targetSum int, pathSum []int) {
		if root == nil {
			return
		}
		sum := 0
		pathSum = append(pathSum, root.Val)
		for i := len(pathSum) - 1; i >= 0; i-- {
			sum += pathSum[i]
			if sum == targetSum {
				count++
			}
		}
		dfs(root.Left, targetSum, pathSum)
		dfs(root.Right, targetSum, pathSum)
	}
	dfs(root, targetSum, []int{})
	return count
}
