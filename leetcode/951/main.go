package main

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func insertLevelOrder(arr []interface{}, root *TreeNode, index int) *TreeNode {
	if index < len(arr) && arr[index] != nil {
		node := &TreeNode{Val: arr[index].(int)}
		root = node

		root.Left = insertLevelOrder(arr, root.Left, index*2+1)
		root.Right = insertLevelOrder(arr, root.Right, index*2+2)
	}

	return root

}

func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	return (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) ||
		(flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left))
}

func main() {
	arr1 := []interface{}{1, 2, 3, 4, 5, 6, nil, nil, nil, 7, 8}
	arr2 := []interface{}{1, 3, 2, nil, 6, 4, 5, nil, nil, nil, nil, 8, 7}
	var root1 *TreeNode
	var root2 *TreeNode
	root1 = insertLevelOrder(arr1, root1, 0)
	root2 = insertLevelOrder(arr2, root2, 0)
	fmt.Println(root1.Left, root1.Right)
	fmt.Println(flipEquiv(root1, root2))
}
