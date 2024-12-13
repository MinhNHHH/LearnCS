package main

import (
	"fmt"
	"math"
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

func (h *MaxHeap) Delete() int {
	if len(h.array) == 0 {
		panic("Heap is empty")
	}

	max := h.array[0]
	last := len(h.array) - 1
	h.array[0] = h.array[last]
	h.array = h.array[:last]

	if len(h.array) > 0 {
		h.heapifDown(0)
	}

	return max
}

func (h *MaxHeap) heapifUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.array[parentIndex] >= h.array[index] {
			break
		}
		h.array[index], h.array[parentIndex] = h.array[parentIndex], h.array[index]
		index = parentIndex
	}
}

func (h *MaxHeap) heapifDown(index int) {
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

func pickGifts(gifts []int, k int) int64 {
	heap := NewMaxHeap()
	sum := 0
	for _, value := range gifts {
		heap.Insert(value)
	}

	for i := 0; i < k; i++ {
		maxGift := heap.Delete()
		fmt.Println(maxGift)
		sqr := int(math.Sqrt(float64(maxGift)))
		if sqr > 0 {
			heap.Insert(sqr)
		}
	}
	fmt.Println(heap.array)
	for _, value := range heap.array {
		sum += value
	}

	return int64(sum)
}

func main() {
	fmt.Println(pickGifts([]int{54, 6, 34, 66, 63, 52, 39, 62, 46, 75, 28, 65, 18, 37, 18, 13, 33, 69, 19, 40, 13, 10, 43, 61, 72}, 7))
}
