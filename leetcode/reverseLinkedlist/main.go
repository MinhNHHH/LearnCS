// Given the head of a singly linked list, reverse the list, and return the reversed list.

// Input: head = [1,2,3,4,5]
// Output: [5,4,3,2,1]

package main

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var dummyNode *ListNode
	curr := head

	for curr != nil {
		tmp := curr.Next
		curr.Next = dummyNode
		dummyNode = curr
		curr = tmp
	}

	return dummyNode
}
