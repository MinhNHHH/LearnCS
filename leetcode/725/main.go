package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func createNode(arr []int) *ListNode {
	if len(arr) == 0 {
		return &ListNode{}
	}

	node := &ListNode{}
	current := node

	for _, num := range arr {
		newNode := &ListNode{
			Val: num,
		}
		current.Next = newNode
		current = newNode
	}
	return node.Next
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	res := make([]*ListNode, k)
	count := 0
	current := head
	for current != nil {
		count++
		current = current.Next
	}

	splitNodes := count / k
	extraNode := count % k
	i := 0
	current = head
	for current != nil {
		res[i] = current
		currentPathLength := splitNodes
		if extraNode > 0 {
			currentPathLength++
			extraNode--
		}

		for i := 0; i < currentPathLength; i++ {
			current = current.Next
		}

		temp := current.Next
		current.Next = nil
		current = temp
		i++

	}

	return res

}

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	head := createNode(arr)
	fmt.Println(splitListToParts(head, 3))
}
