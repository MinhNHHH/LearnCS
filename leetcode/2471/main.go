package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Element struct {
	node  *TreeNode
	level int
}

func convertArrtoBinary(arr []int, index int) *TreeNode {
	if index >= len(arr) {
		return nil
	}

	node := &TreeNode{Val: arr[index]}

	node.Left = convertArrtoBinary(arr, 2*index+1)
	node.Right = convertArrtoBinary(arr, 2*index+2)

	return node
}

func minSwapsToSort(arr []int) int {
	n := len(arr)
	arrPos := make([][2]int, n)

	// Pair elements with indices
	for i, value := range arr {
		arrPos[i] = [2]int{value, i}
	}

	// Sort based on the array values
	sort.Slice(arrPos, func(i, j int) bool {
		return arrPos[i][0] < arrPos[j][0]
	})

	visited := make([]bool, n)
	swaps := 0

	for i := 0; i < n; i++ {
		// Skip already visited elements or elements already in correct position
		if visited[i] || arrPos[i][1] == i {
			continue
		}

		// Compute the size of the cycle
		cycleSize := 0
		j := i
		for !visited[j] {
			visited[j] = true
			j = arrPos[j][1]
			cycleSize++
		}

		// Add (cycleSize - 1) swaps for the current cycle
		if cycleSize > 1 {
			swaps += (cycleSize - 1)
		}
	}

	return swaps
}
func minimumOperations(root *TreeNode) int {
	stack := []*Element{{node: root, level: 0}}
	temp := [][]int{}
	for len(stack) > 0 {
		current := stack[0]
		stack = stack[1:]

		if len(temp) <= current.level {
			temp = append(temp, []int{})
		}

		temp[current.level] = append(temp[current.level], current.node.Val)

		if current.node.Left != nil {

			stack = append(stack, &Element{node: current.node.Left, level: current.level + 1})
		}

		if current.node.Right != nil {
			stack = append(stack, &Element{node: current.node.Right, level: current.level + 1})
		}
	}

	swaps := 0
	for _, arr := range temp {
		swaps += minimunSwaps(arr)
	}
	return swaps
}

func main() {
	arr := []int{332, 463, 103, 417, 150, 409, 41, 135, 129, 117, 474, 263, 328, 456, 347, 167, 383, 422, 493, 489, 275, 72, 425, 89, 162, 18, 363, 290, 106, 260, 468, 432, 323, 36, 302, 190, 280, 488, 446, 75}
	tree := convertArrtoBinary(arr, 0)
	fmt.Println(minimumOperations(tree))
}
