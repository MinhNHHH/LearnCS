package main

import "fmt"

func moveZeros(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	l, r := 0, 0

	for l < len(arr) {
		if arr[l] == 0 {
			l++
		} else {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r++
		}
	}
	return arr
}

func main() {
	fmt.Println("test1", moveZeros([]int{0, 1, 0, 3, 12}))
	fmt.Println("test2", moveZeros([]int{0, 0, 0, 3, 12}))
}
