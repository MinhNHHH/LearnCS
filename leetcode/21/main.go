package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	point := res

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			point.Next = list1
			list1 = list1.Next
		} else {
			point.Next = list2
			list2 = list2.Next
		}
		point = point.Next
	}

	if list1 != nil {
		point.Next = list1
	} else {
		point.Next = list2
	}

	return res.Next
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
	return res
}
func main() {
	list1 := createListNode([]int{1, 2, 4})
	list2 := createListNode([]int{1, 3, 4})
	fmt.Println(mergeTwoLists(list1.Next, list2.Next))
}
