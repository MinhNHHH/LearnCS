package main

import (
	"fmt"
)

type MinHeap struct {
	array []int
}

func NewHeap() *MinHeap {
	return &MinHeap{}
}

func (h *MinHeap) insert(key int) {
	h.array = append(h.array, key)
	h.heapifyUp(len(h.array) - 1)
}

func (h *MinHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.array[parentIndex] > h.array[index] {
			h.array[parentIndex], h.array[index] = h.array[index], h.array[parentIndex]
			index = parentIndex
		} else {
			return
		}
	}
}

func (h *MinHeap) pop() int {
	if len(h.array) == 0 {
		panic("Heap is empty")
	}

	current := h.array[0]
	last := len(h.array) - 1
	h.array[0] = h.array[last]
	h.array = h.array[:last]

	h.heapifyDown(0)

	return current
}

func (h *MinHeap) heapifyDown(index int) {
	for {
		leftChilIndex := 2*index + 1
		rightChildIndex := 2*index + 2
		smallestIndex := index

		if leftChilIndex < len(h.array) && h.array[leftChilIndex] < h.array[smallestIndex] {
			smallestIndex = leftChilIndex
		}
		if rightChildIndex < len(h.array) && h.array[rightChildIndex] < h.array[smallestIndex] {
			smallestIndex = rightChildIndex
		}

		if smallestIndex != index {
			h.array[smallestIndex], h.array[index] = h.array[index], h.array[smallestIndex]
			index = smallestIndex
		} else {
			break
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minOperations(nums []int, k int) int {
	heap := NewHeap()
	for _, num := range nums {
		heap.insert(num)
	}
	count := 0
	for {
		num1 := heap.pop()
		if num1 >= k {
			return count
		}
		num2 := heap.pop()
		num := min(num1, num2)*2 + max(num1, num2)
		count++
		heap.insert(num)

	}
}

func main() {
	fmt.Println(minOperations([]int{1, 1, 2, 4, 9}, 20))
}
