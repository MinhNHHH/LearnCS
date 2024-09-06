package main

import "fmt"

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

func modifiedList(nums []int, head *ListNode) *ListNode {
	visited := map[int]bool{}
	for _, num := range nums {
		visited[num] = true
	}
	res := &ListNode{}
	point := res
	point.Next = head
	for point != nil {
		if point.Next != nil && visited[point.Next.Val] {
			point.Next = point.Next.Next
		} else {
			point = point.Next
		}
	}
	return res.Next
}

func main() {
	ints := []int{1, 2, 3, 4, 5}
	nums := []int{1, 2, 3}
	head := &ListNode{}
	current := head
	for _, num := range ints {
		newNode := &ListNode{
			Val: num,
		}
		current.Next = newNode
		current = newNode
	}

	fmt.Println(modifiedList(nums, head.Next))
}
