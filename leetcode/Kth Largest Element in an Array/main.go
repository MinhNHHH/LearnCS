package main

import "fmt"

func main() {
	fmt.Println(findKthLargest([]int{1, 2, 4, 3, 5, 7, 6}, 1))
}

// With length is <= 10^5. We cannot use QuickSort
func QuickSort(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	patitionIndex := patition(nums)
	QuickSort(nums[:patitionIndex])
	QuickSort(nums[patitionIndex+1:])
	return nums
}

func patition(nums []int) int {
	piviot := nums[len(nums)-1]
	index := -1

	for i, value := range nums {
		if value < piviot {
			index++
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
	nums[index+1], nums[len(nums)-1] = nums[len(nums)-1], nums[index+1]
	return index + 1
}

func findKthLargestWithQuickSort(nums []int, k int) int {
	nums = QuickSort(nums)
	kthLargest := nums[len(nums)-k]
	return kthLargest
}

type MaxHeap struct {
	array []int
}

func NewHeap() *MaxHeap {
	return &MaxHeap{}
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

func (h *MaxHeap) insert(num int) {
	h.array = append(h.array, num)
	h.heapifyUp(len(h.array) - 1)
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
		lastestIndex := index

		if leftChildIndex < len(h.array) && h.array[lastestIndex] < h.array[leftChildIndex] {
			lastestIndex = leftChildIndex
		}

		if rightChildIndex < len(h.array) && h.array[lastestIndex] < h.array[rightChildIndex] {
			lastestIndex = rightChildIndex
		}

		if lastestIndex == index {
			break
		}

		h.array[lastestIndex], h.array[index] = h.array[index], h.array[lastestIndex]
		index = lastestIndex
	}
}

func findKthLargest(nums []int, k int) int {
	heap := NewHeap()
	for _, value := range nums {
		heap.insert(value)
	}
	fmt.Println(heap.array)
	result := 0

	for k >= 0 {
		result = heap.DeleteMax()
		k--
	}
	return result
}
