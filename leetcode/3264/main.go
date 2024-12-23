package main

import (
	"container/heap"
	"fmt"
)

type Element struct {
	value, index int
}

// MinHeap structure
type MinHeap []Element

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool {
	return h[i].value < h[j].value || (h[i].value == h[j].value && h[i].index < h[j].index)
}
func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// Modified getFinalState function
func getFinalState(nums []int, k int, multiplier int) []int {
	h := &MinHeap{}
	heap.Init(h)

	// Push all elements with their indices into the heap
	for i, num := range nums {
		heap.Push(h, Element{value: num, index: i})
	}

	// Perform k operations
	for i := 0; i < k; i++ {
		// Get the minimum element
		minElement := heap.Pop(h).(Element)

		// Update nums array
		nums[minElement.index] = minElement.value * multiplier

		// Push the updated value back into the heap
		heap.Push(h, Element{value: nums[minElement.index], index: minElement.index})
	}

	// Return the updated nums array
	return nums
}

func main() {
	// Example input
	nums := []int{4, 7, 9, 1, 3}
	k := 2
	multiplier := 3

	// Output the final state of the array
	result := getFinalState(nums, k, multiplier)
	fmt.Println("Final state:", result)
}
