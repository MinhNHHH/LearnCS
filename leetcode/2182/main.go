package main

import "fmt"

type Element struct {
	char  rune
	count int
}

type MaxHeap struct {
	array []Element
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (h *MaxHeap) Insert(data interface{}) {
	h.array = append(h.array, data.(Element))
	h.heapifUp(len(h.array) - 1)
}

func (h *MaxHeap) heapifUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.array[parentIndex].char > h.array[index].char {
			break
		}
		h.array[parentIndex], h.array[index] = h.array[index], h.array[parentIndex]
		index = parentIndex
	}
}

func (h *MaxHeap) Pop() Element {
	if len(h.array) == 0 {
		panic("Heap is empty")
	}

	maxHeap := h.array[0]
	last := len(h.array) - 1
	h.array[0] = h.array[last]
	h.array = h.array[:last]

	if len(h.array) > 0 {
		h.heapifDown(0)
	}
	return maxHeap
}

func (h *MaxHeap) heapifDown(index int) {
	for {
		leftChildIndex := index*2 + 1
		rightChildIndex := index*2 + 2
		largestIndex := index

		if leftChildIndex < len(h.array) && h.array[largestIndex].char <= h.array[leftChildIndex].char {
			largestIndex = leftChildIndex
		}
		if rightChildIndex < len(h.array) && h.array[largestIndex].char <= h.array[rightChildIndex].char {
			largestIndex = rightChildIndex
		}
		if largestIndex == index {
			break
		}
		h.array[largestIndex], h.array[index] = h.array[index], h.array[largestIndex]
		index = largestIndex
	}
}

func repeatLimitedString(s string, repeatLimit int) string {
	hashMap := map[rune]int{}
	for _, char := range s {
		hashMap[char]++
	}

	heap := NewMaxHeap()
	for key, value := range hashMap {
		heap.Insert(Element{char: key, count: value})
	}

	newS := []rune{}

	for len(heap.array) > 0 {
		current := heap.Pop()

		k := 0
		for k < repeatLimit && current.count > 0 {
			newS = append(newS, current.char)
			k++
			current.count--
		}

		if current.count > 0 {
			if len(heap.array) == 0 {
				break
			}
			second := heap.Pop()
			newS = append(newS, second.char)
			second.count--
			if second.count > 0 {
				heap.Insert(second)
			}
			heap.Insert(current)
		}
	}

	return string(newS)
}

func main() {
	fmt.Println(repeatLimitedString("cczazcc", 3))
}
