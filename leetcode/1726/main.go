package main

import "fmt"

func tupleSameProduct(nums []int) int {
	hashTable := map[int]int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			hashTable[nums[i]*nums[j]]++
		}
	}
	ans := 0
	for _, value := range hashTable {
		ans += (value - 1) * value * 4
	}
	return ans
}

func main() {
	fmt.Println(tupleSameProduct([]int{2, 3, 4, 6}))
}
