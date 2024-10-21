package main

import (
	"fmt"
)

type MaxHeap struct {
	array []int
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.heapifUp(len(h.array) - 1)
}

func (h *MaxHeap) Pop() int {
	if len(h.array) == 0 {
		panic("Heap is empty")
	}
	maxVal := h.array[0]
	last := len(h.array) - 1
	h.array[0] = h.array[last]
	h.array = h.array[:last]

	if len(h.array) > 0 {
		h.heapifDown(0)
	}

	return maxVal
}

func (h *MaxHeap) heapifDown(index int) {
	for {
		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2
		largestIndex := index

		if leftChildIndex < len(h.array) && h.array[leftChildIndex] >= h.array[largestIndex] {
			largestIndex = leftChildIndex
		}

		if rightChildIndex < len(h.array) && h.array[rightChildIndex] >= h.array[largestIndex] {
			largestIndex = rightChildIndex
		}
		if largestIndex == index {
			break
		}

		h.array[index], h.array[largestIndex] = h.array[largestIndex], h.array[index]
		index = largestIndex
	}
}

func (h *MaxHeap) heapifUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.array[parentIndex] >= h.array[index] {
			break
		}

		h.array[parentIndex], h.array[index] = h.array[index], h.array[parentIndex]
		index = parentIndex
	}
}

func celi(a, b int) int {
	if a%b == 0 {
		return a / b
	}
	return a/b + 1
}

func maxKelements(nums []int, k int) int64 {
	maxScore := 0
	maxHeap := NewMaxHeap()

	for _, num := range nums {
		maxHeap.Insert(num)
	}

	for i := 0; i < k; i++ {
		val := maxHeap.Pop()
		maxScore += val
		fmt.Println(val, celi(val, k), maxHeap.array)
		maxHeap.Insert(celi(val, 3))
	}
	return int64(maxScore)
}

func main() {
	fmt.Println(maxKelements([]int{756902131, 995414896, 95906472, 149914376, 387433380, 848985151}, 6))
}
