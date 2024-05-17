// Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.

// Input: root = [1,2,3,null,5,null,4]
// Output: [1,3,4]

package main

import "fmt"

func main() {
	arr := []interface{}{1, 2, 3, nil, 5, nil, 4}
	var root *TreeNode
	root = insertLevelOrder(arr, root, 0)
	// root2 := insertLevelOrder([]interface{}{1, 2, 3, 4}, root, 0)
	fmt.Println(rightSideView(root))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
func bfs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{root}
	res := []int{}
	for len(queue) > 0 {
		subqueue := []*TreeNode{}
		l := len(queue)
		for index, node := range queue {
			if index == l-1 {
				res = append(res, node.Val)
			}
			if node.Left != nil {
				subqueue = append(subqueue, node.Left)
			}
			if node.Right != nil {
				subqueue = append(subqueue, node.Right)
			}
		}
		queue = subqueue
	}
	return res
}

func rightSideView(root *TreeNode) []int {

	return bfs(root)
}
