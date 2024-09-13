package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	//res := &ListNode{}
	current := head

	for current.Next != nil {
		gcdNum := gcd(current.Val, current.Next.Val)
		newNode := &ListNode{
			Val: gcdNum,
		}

		newNode.Next = current.Next
		current.Next = newNode
		current = current.Next.Next
	}
	return head
}

func PrintNode(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
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
func main() {
	head := createListNode([]int{18, 6, 10, 3})
	news := insertGreatestCommonDivisors(head)
	PrintNode(news)
}
