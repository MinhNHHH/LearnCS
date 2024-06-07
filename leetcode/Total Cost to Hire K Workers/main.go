// You are given a 0-indexed integer array costs where costs[i] is the cost of hiring the ith worker.
// You are also given two integers k and candidates. We want to hire exactly k workers according to the following rules:
// You will run k sessions and hire exactly one worker in each session.
// In each hiring session, choose the worker with the lowest cost from either the first candidates workers or the last candidates workers. Break the tie by the smallest index.
// For example, if costs = [3,2,7,7,1,2] and candidates = 2, then in the first hiring session, we will choose the 4th worker because they have the lowest cost [3,2,7,7,1,2].
// In the second hiring session, we will choose 1st worker because they have the same lowest cost as 4th worker but they have the smallest index [3,2,7,7,2]. Please note that the indexing may be changed in the process.
// If there are fewer than candidates workers remaining, choose the worker with the lowest cost among them. Break the tie by the smallest index.
// A worker can only be chosen once.

// Return the total cost to hire exactly k workers.

// Input: costs = [17,12,10,2,7,2,11,20,8], k = 3, candidates = 4
// Output: 11
// Explanation: We hire 3 workers in total. The total cost is initially 0.
// - In the first hiring round we choose the worker from [17,12,10,2,7,2,11,20,8]. The lowest cost is 2, and we break the tie by the smallest index, which is 3. The total cost = 0 + 2 = 2.
// - In the second hiring round we choose the worker from [17,12,10,7,2,11,20,8]. The lowest cost is 2 (index 4). The total cost = 2 + 2 = 4.
// - In the third hiring round we choose the worker from [17,12,10,7,11,20,8]. The lowest cost is 7 (index 3). The total cost = 4 + 7 = 11. Notice that the worker with index 3 was common in the first and last four workers.
// The total hiring cost is 11.

// Input: costs = [1,2,4,1], k = 3, candidates = 3
// Output: 4
// Explanation: We hire 3 workers in total. The total cost is initially 0.
// - In the first hiring round we choose the worker from [1,2,4,1]. The lowest cost is 1, and we break the tie by the smallest index, which is 0. The total cost = 0 + 1 = 1.
// Notice that workers with index 1 and 2 are common in the first and last 3 workers.
// - In the second hiring round we choose the worker from [2,4,1]. The lowest cost is 1 (index 2). The total cost = 1 + 1 = 2.
// - In the third hiring round there are less than three candidates. We choose the worker from the remaining workers [2,4]. The lowest cost is 2 (index 0). The total cost = 2 + 2 = 4.
// The total hiring cost is 4.

package main

import (
	"fmt"
)

func main() {
	// fmt.Println(totalCost([]int{1, 2, 4, 1}, 3, 3))
	// fmt.Println(totalCost([]int{17, 12, 10, 2, 7, 2, 11, 20, 8}, 3, 4))
	fmt.Println(totalCost([]int{2, 2, 2, 2, 2, 2, 1, 4, 5, 5, 5, 5, 5, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}, 7, 3))
}

// fmt.Println(totalCost([]int{31, 25, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 17, 58}, 11, 2)) 0
// fmt.Println(totalCost([]int{31, 72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 58}, 11, 2)) 1
// fmt.Println(totalCost([]int{72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 33, 58}, 11, 2)) 2
// fmt.Println(totalCost([]int{72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81, 58}, 11, 2)) 3
// fmt.Println(totalCost([]int{72, 79, 74, 65, 84, 91, 18, 59, 27, 9, 81}, 11, 2)) 4
// fmt.Println(totalCost([]int{72, 79, 74, 65, 84, 91, 18, 59, 27, 81}, 11, 2)) 5
// fmt.Println(totalCost([]int{72, 79, 74, 65, 84, 91, 18, 59, 81}, 11, 2)) 6
// fmt.Println(totalCost([]int{72, 79, 74, 65, 84, 91, 18, 81}, 11, 2)) 7
// fmt.Println(totalCost([]int{72, 79, 74, 65, 84, 91, 81}, 11, 2)) 8
// fmt.Println(totalCost([]int{79, 74, 65, 84, 91, 81}, 11, 2)) 9
// fmt.Println(totalCost([]int{79, 65, 84, 91, 81}, 11, 2)) 10
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func totalCost(costs []int, k int, candidates int) int64 {
	totalCost := int64(0)

	heap := NewMinHeap()

	for i := 0; i < candidates; i++ {
		heap.Insert([]int{costs[i], 0})
	}
	for i := len(costs) - candidates; i < len(costs); i++ {
		heap.Insert([]int{costs[i], 1})
	}
	left := candidates
	right := len(costs) - candidates - 1

	for i := 0; i < k; i++ {
		current := heap.popSmallest()
		totalCost = totalCost + int64(current[0])
		if left <= right {
			if current[1] == 0 {
				heap.Insert([]int{costs[left], 0})
				left++
			} else {
				heap.Insert([]int{costs[right], 1})
				right--
			}
		}
	}
	return totalCost
}

type MinHeap struct {
	array [][]int
}

func NewMinHeap() *MinHeap {
	return &MinHeap{}
}

func (this *MinHeap) Insert(listNums []int) {
	this.array = append(this.array, listNums)
	this.heapifyUp(len(this.array) - 1)
}

func (this *MinHeap) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if this.array[parentIndex][0] <= this.array[index][0] {
			break
		}
		this.array[index], this.array[parentIndex] = this.array[parentIndex], this.array[index]
		index = parentIndex
	}
}

func (this *MinHeap) heapifyDown(index int) {
	for {
		leftChildIndex := index*2 + 1
		rightChildIdnex := index*2 + 2
		smallestIndex := index

		if leftChildIndex < len(this.array) && this.array[smallestIndex][0] >= this.array[leftChildIndex][0] {
			smallestIndex = leftChildIndex
		}
		if rightChildIdnex < len(this.array) && this.array[smallestIndex][0] >= this.array[rightChildIdnex][0] {
			smallestIndex = rightChildIdnex
		}
		if smallestIndex == index {
			break
		}

		this.array[smallestIndex], this.array[index] = this.array[index], this.array[smallestIndex]
		index = smallestIndex
	}
}

func (this *MinHeap) popSmallest() []int {
	if len(this.array) == 0 {
		panic("Heap is empty")
	}
	smallest := this.array[0]
	this.array[0] = this.array[len(this.array)-1]
	this.array = this.array[:len(this.array)-1]
	this.heapifyDown(0)
	return smallest
}
