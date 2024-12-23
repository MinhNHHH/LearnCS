package main

import "fmt"

func monotonicStack(arr []int) []int {
	stack := []int{}
	for _, value := range arr {
		if len(stack) == 0 {
			stack = append(stack, value)
		} else {
			maxVal := stack[0]
			for len(stack) > 0 && stack[0] > value {
				stack = stack[1:]
			}
			stack = append(stack, maxVal)
		}
	}
	return stack
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxChunksToSorted(arr []int) int {
	arr = monotonicStack(arr)
	return len(arr)
}

func main() {
	fmt.Println(maxChunksToSorted([]int{1, 0, 2, 3, 4}))
}
