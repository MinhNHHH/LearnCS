package main

import "fmt"

func findThePrefixCommonArray(A []int, B []int) []int {
	C := make([]int, len(A))
	temp := make([]int, len(A)+1)
	count := 0
	for i := 0; i < len(C); i++ {
		temp[A[i]]++
		if temp[A[i]] == 2 {
			count++
		}
		temp[B[i]]++
		if temp[B[i]] == 2 {
			count++
		}
		C[i] = count
	}
	return C
}

func main() {
	fmt.Println(findThePrefixCommonArray([]int{1, 3, 2, 4}, []int{3, 1, 2, 4}))
}
