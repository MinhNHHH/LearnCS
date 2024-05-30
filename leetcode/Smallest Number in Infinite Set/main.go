// You have a set which contains all positive integers [1, 2, 3, 4, 5, ...].

// Implement the SmallestInfiniteSet class:

// SmallestInfiniteSet() Initializes the SmallestInfiniteSet object to contain all positive integers.
// int popSmallest() Removes and returns the smallest integer contained in the infinite set.
// void addBack(int num) Adds a positive integer num back into the infinite set, if it is not already in the infinite set.

// Input
// ["SmallestInfiniteSet", "addBack", "popSmallest", "popSmallest", "popSmallest", "addBack", "popSmallest", "popSmallest", "popSmallest"]
// [[], [2], [], [], [], [1], [], [], []]
// Output
// [null, null, 1, 2, 3, null, 1, 4, 5]

// Explanation
// SmallestInfiniteSet smallestInfiniteSet = new SmallestInfiniteSet();
// smallestInfiniteSet.addBack(2);    // 2 is already in the set, so no change is made.
// smallestInfiniteSet.popSmallest(); // return 1, since 1 is the smallest number, and remove it from the set.
// smallestInfiniteSet.popSmallest(); // return 2, and remove it from the set.
// smallestInfiniteSet.popSmallest(); // return 3, and remove it from the set.
// smallestInfiniteSet.addBack(1);    // 1 is added back to the set.
// smallestInfiniteSet.popSmallest(); // return 1, since 1 was added back to the set and
//                                    // is the smallest number, and remove it from the set.
// smallestInfiniteSet.popSmallest(); // return 4, and remove it from the set.
// smallestInfiniteSet.popSmallest(); // return 5, and remove it from the set.

package main

func main() {

}

type SmallestInfiniteSet struct {
	array []int
	added map[int]bool
}

func Constructor() SmallestInfiniteSet {
	arr := []int{}
	for i := 1; i <= 1000; i++ {
		arr = append(arr, i)
	}
	return SmallestInfiniteSet{
		array: arr,
		added: map[int]bool{},
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	smallest := this.array[0]
	this.added[smallest] = false
	temp := this.array[len(this.array)-1]
	this.array[0] = temp
	this.array = this.array[:len(this.array)-1]
	this.heapifyDown(0)
	return smallest
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num == 0 {
		return
	}
	if !this.added[num] {
		this.array = append(this.array, num)
		this.added[num] = true
	}
	this.heapifyUp(len(this.array) - 1)
}

func (this *SmallestInfiniteSet) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if this.array[parentIndex] <= this.array[index] {
			break
		}
		this.array[index], this.array[parentIndex] = this.array[parentIndex], this.array[index]
		index = parentIndex
	}

}
func (this *SmallestInfiniteSet) heapifyDown(index int) {
	for {
		leftChildIndex := index*2 + 1
		rightChildIndex := index*2 + 2
		smallestIndex := index
		if leftChildIndex < len(this.array) && this.array[smallestIndex] > this.array[leftChildIndex] {
			smallestIndex = leftChildIndex
		}

		if rightChildIndex < len(this.array) && this.array[smallestIndex] > this.array[rightChildIndex] {
			smallestIndex = rightChildIndex
		}

		if smallestIndex == index {
			break
		}

		this.array[smallestIndex], this.array[index] = this.array[index], this.array[smallestIndex]
		index = smallestIndex
	}
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
