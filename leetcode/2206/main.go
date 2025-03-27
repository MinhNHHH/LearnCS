package main

import "fmt"

func divideArray(nums []int) bool {
	totalPair := len(nums) / 2
	hashMap := map[int]int{}

	for _, num := range nums {
		hashMap[num]++
	}
	count := 0
	fmt.Println(hashMap)
	for _, val := range hashMap {
		count += val / 2
	}
	return totalPair == count
}

func main() {
	fmt.Println(divideArray([]int{3, 2, 3, 2, 2, 2}))
}
