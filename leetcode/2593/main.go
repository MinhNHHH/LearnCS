package main

import "fmt"

type MinHeap struct {
	array []int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func (h *MinHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.heapifUp(len(h.array) - 1)
}

func (h *MinHeap) heapifUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.array[parentIndex] > h.array[index] {
			h.array[parentIndex], h.array[index] = h.array[index], h.array[parentIndex]
			index = parentIndex
		} else {
			break
		}
	}
}

func (h *MinHeap) Delete() int {
	if len(h.array) == 0 {
		panic("Heap is empty")
	}

	min := h.array[0]
	last := len(h.array) - 1
	h.array[0] = h.array[last]
	h.array = h.array[:last]

	if len(h.array) > 0 {
		h.heapifDown(0)
	}
	return min
}

func (h *MinHeap) heapifDown(index int) {
	for {
		leftChildIndex := index*2 + 1
		rightChildIndex := index*2 + 2
		smallestIndex := index

		if leftChildIndex < len(h.array) && h.array[leftChildIndex] < h.array[smallestIndex] {
			smallestIndex = leftChildIndex
		}
		if rightChildIndex < len(h.array) && h.array[rightChildIndex] < h.array[smallestIndex] {
			smallestIndex = rightChildIndex
		}

		if smallestIndex == index {
			break
		}

		h.array[smallestIndex], h.array[index] = h.array[index], h.array[smallestIndex]
		index = smallestIndex
	}
}

func findScore(nums []int) int64 {
	heap := NewMinHeap()

	for _, val := range nums {
		heap.Insert(val)
	}

	fmt.Println(heap.array)
	return int64(0)
}

func main() {
	fmt.Println(findScore([]int{2, 1, 3, 4, 5, 2}))
}
