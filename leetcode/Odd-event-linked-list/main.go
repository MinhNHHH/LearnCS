// Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.

// The first node is considered odd, and the second node is even, and so on.

// Note that the relative order inside both the even and odd groups should remain as it was in the input.

// You must solve the problem in O(1) extra space complexity and O(n) time complexity.

// EX1
// Input: head = [1,2,3,4,5]
// Output: [1,3,5,2,4]

// EX2
// Input: head = [2,1,3,5,6,4,7]
// Output: [2,3,6,7,1,5,4]

package main

func main() {
	values := []int{2, 1, 3, 5, 6, 4, 7}

	res := &ListNode{}
	point := res
	for _, val := range values {
		point.Next = &ListNode{Val: val}
		point = point.Next
	}
	oddEvenList(res.Next)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	count := 1

	odd, even := &ListNode{}, &ListNode{}
	currOdd := odd
	currEven := even

	for head != nil {
		if count%2 == 0 {
			currEven.Next = &ListNode{Val: head.Val}
			currEven = currEven.Next
		} else {
			currOdd.Next = &ListNode{Val: head.Val}
			currOdd = currOdd.Next
		}
		head = head.Next
		count++
	}
	currOdd.Next = even.Next

	return odd.Next
}
