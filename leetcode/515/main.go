package main

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateBNT(arr []int, index int) *TreeNode {
	if index >= len(arr) {
		return nil
	}

	node := &TreeNode{Val: arr[index]}

	node.Left = generateBNT(arr, index*2+1)
	node.Right = generateBNT(arr, index*2+2)

	return node
}

type Element struct {
	node  *TreeNode
	level int
}

func largestValues(root *TreeNode) []int {

	queue := []*Element{{node: root, level: 0}}
	temp := [][]int{}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if len(temp) <= current.level {
			temp = append(temp, []int{})
		}

		temp[current.level] = append(temp[current.level], current.node.Val)

		if current.node.Left != nil {
			queue = append(queue, &Element{node: current.node.Left, level: current.level + 1})
		}

		if current.node.Right != nil {
			queue = append(queue, &Element{node: current.node.Right, level: current.level + 1})
		}
	}

	res := []int{}
	for _, arr := range temp {
		sort.Ints(arr)
		res = append(res, arr[len(arr)-1])
	}

	return res
}

func main() {
	arr := []int{1, 3, 2, 5, 3, 9}
	tree := generateBNT(arr, 0)
	largestValues(tree)
}
