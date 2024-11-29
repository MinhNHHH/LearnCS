package main

import "fmt"

func main() {
	fmt.Println(findLengthOfShortestSubarray([]int{1, 2, 3, 10, 4, 2, 3, 5}))
}

func findLengthOfShortestSubarray(arr []int) int {
	res := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			temp := arr[i]
			for j := i + 1; j < len(arr)-1; j++ {
				if temp < arr[j] {
					res++
				} else {
					break
				}
			}
		}
	}

	return res
}
