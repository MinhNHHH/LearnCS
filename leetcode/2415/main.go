package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeLevel struct {
	node  *TreeNode
	level int
}

// ConvertToBinaryTreeRecursive recursively builds a binary tree from a list of integers
func ConvertToBinaryTreeRecursive(values []int, index int) *TreeNode {
	if index >= len(values) {
		return nil
	}

	// Create the current node
	node := &TreeNode{Val: values[index]}

	// Recursively create left and right children
	node.Left = ConvertToBinaryTreeRecursive(values, 2*index+1)
	node.Right = ConvertToBinaryTreeRecursive(values, 2*index+2)

	return node
}

// ConvertToBinaryTree is a wrapper function to simplify usage
func ConvertToBinaryTree(values []int) *TreeNode {
	return ConvertToBinaryTreeRecursive(values, 0)
}

func reverseArr(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func flatArr(arr [][]int) []int {
	res := []int{}

	for _, a := range arr {
		res = append(res, a...)
	}

	return res
}

func reverseOddLevels(root *TreeNode) *TreeNode {

	stack := []*NodeLevel{{node: root, level: 0}}
	traverseValues := [][]int{}
	for len(stack) > 0 {
		current := stack[0]
		stack = stack[1:]

		if len(traverseValues) <= current.level {
			traverseValues = append(traverseValues, []int{})
		}
		traverseValues[current.level] = append(traverseValues[current.level], current.node.Val)
		if current.node.Left != nil {
			stack = append(stack, &NodeLevel{node: current.node.Left, level: current.level + 1})
		}
		if current.node.Right != nil {
			stack = append(stack, &NodeLevel{node: current.node.Right, level: current.level + 1})
		}
	}

	for index, arr := range traverseValues {
		if index%2 != 0 {
			arr = reverseArr(arr)
		}
	}

	flatArr := flatArr(traverseValues)
	return ConvertToBinaryTree(flatArr)
}

func main() {
	ex := ConvertToBinaryTree([]int{2, 3, 5, 8, 13, 21, 34})
	fmt.Println(reverseOddLevels(ex))
}
