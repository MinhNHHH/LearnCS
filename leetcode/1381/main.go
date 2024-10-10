package main

type CustomStack struct {
	index    int
	capacity int
	stack    []int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		stack:    make([]int, maxSize),
		capacity: maxSize,
		index:    0,
	}
}

func (this *CustomStack) Push(x int) {
	if this.index == this.capacity {
		return
	}

	this.stack[this.index] = x
	this.index++
}

func (this *CustomStack) Pop() int {
	popValue := this.stack[this.index]
	this.stack[this.index] = 0
	this.index--
	return popValue
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (this *CustomStack) Increment(k int, val int) {
	k = min(k, len(this.stack))
	for i := 0; i < k; i++ {
		this.stack[i] = this.stack[i] + val
	}
}
func main() {

}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
