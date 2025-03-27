package main

import "fmt"

type ProductOfNumbers struct {
	array []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		array: []int{},
	}
}

func (this *ProductOfNumbers) Add(num int) {

	this.array = append(this.array, num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	product := 1
	for i := 0; i < k; i++ {
		product = product * this.array[len(this.array)-i-1]
	}
	return product
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */

func main() {
	fmt.Println("vim-go")
}
