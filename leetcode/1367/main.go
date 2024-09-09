package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubPath(head *ListNode, root *TreeNode) bool {
	var checkPath func(path *ListNode, root *TreeNode) bool
	checkPath = func(path *ListNode, root *TreeNode) bool {
		if path == nil {
			return true
		}
		if root == nil || path.Val != root.Val {
			return false
		}
		return checkPath(path.Next, root.Left) || checkPath(path.Next, root.Right)
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if checkPath(head, current) {
			return true
		}

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
	return false
}

func createListNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return &ListNode{}
	}
	res := &ListNode{}
	current := res

	for _, num := range arr {
		newNode := &ListNode{
			Val: num,
		}
		current.Next = newNode
		current = newNode
	}

	return res.Next
}

func createTreeFromList(lst []interface{}) *TreeNode {
	if len(lst) == 0 || lst[0] == nil {
		return nil
	}

	// Root node
	root := &TreeNode{Val: lst[0].(int)}
	queue := []*TreeNode{root}
	index := 1

	for index < len(lst) {
		node := queue[0]  // Take the node at the front of the queue
		queue = queue[1:] // Pop the node from the queue

		// Process left child
		if index < len(lst) && lst[index] != nil {
			leftNode := &TreeNode{Val: lst[index].(int)}
			node.Left = leftNode
			queue = append(queue, leftNode)
		}
		index++

		// Process right child
		if index < len(lst) && lst[index] != nil {
			rightNode := &TreeNode{Val: lst[index].(int)}
			node.Right = rightNode
			queue = append(queue, rightNode)
		}
		index++
	}

	return root
}
func main() {
	lst := []interface{}{1, 4, 4, nil, 2, 2, nil, 1, nil, 6, 8, nil, nil, nil, nil, 1, 3}
	tree := createTreeFromList(lst)
	arr := []int{4, 2, 8}
	head := createListNode(arr)
	fmt.Println(isSubPath(head, tree))
}
