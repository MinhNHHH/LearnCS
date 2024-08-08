package main

import "fmt"

type MaxHeap struct {
	array []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.heapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) DeleteMax() int {
	if len(h.array) == 0 {
		panic("Heap is empty")
	}

	max := h.array[0]
	last := len(h.array) - 1
	h.array[0] = h.array[last]
	h.array = h.array[:last]

	if len(h.array) > 0 {
		h.heapifyDown(0)
	}

	return max
}

func (h *MaxHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.array[parentIndex] >= h.array[index] {
			break
		}
		h.array[index], h.array[parentIndex] = h.array[parentIndex], h.array[index]
		index = parentIndex
	}
}

func (h *MaxHeap) heapifyDown(index int) {
	for {
		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2
		largestIndex := index

		if leftChildIndex < len(h.array) && h.array[leftChildIndex] > h.array[largestIndex] {
			largestIndex = leftChildIndex
		}

		if rightChildIndex < len(h.array) && h.array[rightChildIndex] > h.array[largestIndex] {
			largestIndex = rightChildIndex
		}

		if largestIndex == index {
			break
		}

		h.array[index], h.array[largestIndex] = h.array[largestIndex], h.array[index]
		index = largestIndex
	}
}

func main() {
	heap := NewMaxHeap()
	heap.Insert(30)
	heap.Insert(20)
	heap.Insert(12)
	heap.Insert(15)
	heap.Insert(10)
	heap.Insert(40)

	fmt.Println("Heap array after insertion:", heap.array)

	max := heap.DeleteMax()
	fmt.Println("Max element removed from heap:", max)
	fmt.Println("Heap array after deletion:", heap.array)
}
