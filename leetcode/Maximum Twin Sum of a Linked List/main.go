// In a linked list of size n, where n is even, the ith node (0-indexed) of the linked list is known as the twin of the (n-1-i)th node, if 0 <= i <= (n / 2) - 1.

// For example, if n = 4, then node 0 is the twin of node 3, and node 1 is the twin of node 2. These are the only nodes with twins for n = 4.

// The twin sum is defined as the sum of a node and its twin.

// Given the head of a linked list with even length, return the maximum twin sum of the linked list.

// Input: head = [5,4,2,1]
// Output: 6
// Explanation:
// Nodes 0 and 1 are the twins of nodes 3 and 2, respectively. All have twin sum = 6.
// There are no other nodes with twins in the linked list.
// Thus, the maximum twin sum of the linked list is 6.

// Input: head = [4,2,2,3]
// Output: 7
// Explanation:
// The nodes with twins present in this linked list are:
// - Node 0 is the twin of node 3 having a twin sum of 4 + 3 = 7.
// - Node 1 is the twin of node 2 having a twin sum of 2 + 2 = 4.
// Thus, the maximum twin sum of the linked list is max(7, 4) = 7.

package main

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

//Solution 1
// Reverse a half a linkedlist.

func reverse(head *ListNode) *ListNode {
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func pairSum1(head *ListNode) int {
	slow, fast := head, head
	res := 0

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	reverseListNode := reverse(slow)
	point := head
	for reverseListNode != nil && point != nil {
		res = max(res, point.Val+reverseListNode.Val)
		point = point.Next
		reverseListNode = reverseListNode.Next
	}
	return res
}

func pairSum2(head *ListNode) int {
	slow, fast := head, head
	res := 0
	stack := []int{}
	for fast != nil && fast.Next != nil {
		stack = append(stack, slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}

	for slow != nil {
		popValue := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = max(res, slow.Val+popValue)
		slow = slow.Next
	}
	return res
}
