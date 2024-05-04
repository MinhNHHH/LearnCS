package main

import (
	"fmt"
)

// MinHeap represents a min heap data structure
type MinHeap struct {
	array []int
}

// NewMinHeap creates a new instance of MinHeap
func NewMinHeap() *MinHeap {
	return &MinHeap{
		array: []int{},
	}
}

// Insert inserts a new element into the heap
func (h *MinHeap) Insert(value int) {
	h.array = append(h.array, value)
	h.heapifyUp(len(h.array) - 1)
}

// heapifyUp restores the heap property from a given index upwards
func (h *MinHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.array[parentIndex] > h.array[index] {
			// Swap parent and child if parent is greater
			h.array[parentIndex], h.array[index] = h.array[index], h.array[parentIndex]
			index = parentIndex
		} else {
			// If parent is smaller or equal, heap property is satisfied
			break
		}
	}
}

// DeleteRoot removes the root (minimum) element from the heap
func (h *MinHeap) DeleteRoot() int {
	if len(h.array) == 0 {
		panic("Heap is empty")
	}

	// Remove the root element
	root := h.array[0]
	last := len(h.array) - 1
	h.array[0] = h.array[last]
	h.array = h.array[:last]

	// Restore the heap property starting from the root
	h.heapifyDown(0)

	return root
}

// heapifyDown restores the heap property from a given index downwards
func (h *MinHeap) heapifyDown(index int) {
	for {
		leftChild := 2*index + 1
		rightChild := 2*index + 2
		smallest := index

		// Find the smallest child
		if leftChild < len(h.array) && h.array[leftChild] < h.array[smallest] {
			smallest = leftChild
		}
		if rightChild < len(h.array) && h.array[rightChild] < h.array[smallest] {
			smallest = rightChild
		}

		// Swap with the smallest child if necessary
		if smallest != index {
			h.array[index], h.array[smallest] = h.array[smallest], h.array[index]
			index = smallest
		} else {
			break
		}
	}
}

func main() {
	// Create a new min heap
	heap := NewMinHeap()

	// Insert elements into the heap
	heap.Insert(10)
	heap.Insert(20)
	heap.Insert(15)
	heap.Insert(5)
	fmt.Println(heap.array)
	// Delete the root element (minimum)
	fmt.Println("Deleted root:", heap.DeleteRoot())

	// Output the remaining elements in the heap
	fmt.Println("Remaining elements in the heap:", heap.array)
}
