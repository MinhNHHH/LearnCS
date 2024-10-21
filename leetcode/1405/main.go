package main

import (
	"fmt"
	"strings"
)

type Tupple struct {
	Str string
	Int int
}

type MaxHeap struct {
	array []Tupple
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{}
}

func (h *MaxHeap) Insert(val Tupple) {
	h.array = append(h.array, val)
	h.heapIfUp(len(h.array) - 1)
}

func (h *MaxHeap) heapIfUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if h.array[parentIndex].Int >= h.array[index].Int {
			break
		}
		h.array[index], h.array[parentIndex] = h.array[parentIndex], h.array[index]
		index = parentIndex
	}
}

func (h *MaxHeap) heapIfDown(index int) {
	for {
		leftChildIndex := index*2 + 1
		rightChildIndex := index*2 + 2
		largestIndex := index

		if leftChildIndex < len(h.array) && h.array[leftChildIndex].Int > h.array[largestIndex].Int {
			largestIndex = leftChildIndex
		}
		if rightChildIndex < len(h.array) && h.array[leftChildIndex].Int > h.array[largestIndex].Int {
			largestIndex = rightChildIndex
		}
		if largestIndex == index {
			break
		}
		h.array[index], h.array[largestIndex] = h.array[largestIndex], h.array[index]
		index = largestIndex
	}
}

func (h *MaxHeap) Pop() Tupple {
	if len(h.array) == 0 {
		panic("Heap is empty")
	}
	maxV := h.array[0]
	last := len(h.array) - 1
	h.array[0] = h.array[last]
	h.array = h.array[:last]

	if len(h.array) > 0 {
		h.heapIfDown(0)
	}
	return maxV
}

func longestDiverseString(a int, b int, c int) string {
	heap := NewMaxHeap()
	if a > 0 {
		heap.Insert(Tupple{
			Str: "a",
			Int: a,
		})
	}

	if b > 0 {
		heap.Insert(Tupple{
			Str: "b",
			Int: b,
		})
	}

	if c > 0 {
		heap.Insert(Tupple{
			Str: "c",
			Int: c,
		})
	}
	s := []string{}

	for len(heap.array) > 0 {
		current := heap.Pop()
		if len(s) >= 2 && s[len(s)-1] == current.Str && s[len(s)-2] == current.Str {
			if len(heap.array) == 0 {
				break
			}
			second := heap.Pop()
			s = append(s, second.Str)
			second.Int--
			if second.Int > 0 {
				heap.Insert(Tupple{
					Str: second.Str,
					Int: second.Int,
				})
			}
			heap.Insert(Tupple{
				Str: current.Str,
				Int: current.Int,
			})
		} else {
			s = append(s, current.Str)
			current.Int--
			if current.Int > 0 {
				heap.Insert(Tupple{
					Str: current.Str,
					Int: current.Int,
				})
			}

		}
	}

	return strings.Join(s, "")
}

func main() {
	fmt.Println(longestDiverseString(2, 2, 1))
}
