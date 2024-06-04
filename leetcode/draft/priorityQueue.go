package main

import (
	"container/heap"
	"fmt"
)

// Item represents an element in the priority queue.
type Item struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	index    int    // The index of the item in the heap.
}

// PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

// Len returns the number of items in the queue.
func (pq PriorityQueue) Len() int { return len(pq) }

// Less determines the order of items in the queue.
// We want Pop to give us the highest priority, so we use greater than here.
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

// Swap swaps the items with indexes i and j.
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push adds an item to the queue.
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop removes and returns the highest priority item from the queue.
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	// Create a priority queue and add some items.
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// Insert items with priorities.
	items := map[string]int{
		"task1": 3,
		"task2": 2,
		"task3": 1,
	}
	for value, priority := range items {
		item := &Item{
			value:    value,
			priority: priority,
		}
		heap.Push(&pq, item)
	}

	// Update the priority of an item.
	item := &Item{
		value:    "task2",
		priority: 5,
	}
	heap.Push(&pq, item)

	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d: %s\n", item.priority, item.value)
	}
}
