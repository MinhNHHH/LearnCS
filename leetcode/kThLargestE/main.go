package main

type KthLargest struct {
	k   int
	arr []int
}

func Constructor(k int, nums []int) KthLargest {
	kTh := KthLargest{
		k:   k,
		arr: []int{},
	}
	for _, value := range nums {
		kTh.Add(value)
	}
	return kTh
}

func (this *KthLargest) Add(val int) int {
	this.arr = append(this.arr, val)
	this.heapifyUp(len(this.arr))

	return this.arr[this.k-1]
}

func (this *KthLargest) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if this.arr[parentIndex] >= this.arr[index] {
			break
		}
		this.arr[parentIndex], this.arr[index] = this.arr[index], this.arr[parentIndex]
		parentIndex = index
	}
}
func main() {}
