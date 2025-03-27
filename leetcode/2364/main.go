package main

import "fmt"

// i < j and i-j != nums[j] - nums[i]
// badPairs = totalPairs - goodPairs
// totalPairs = (len(nums)*len(nums)-1)/2

// nums[i] - i != nums[j] - j

func countBadPairs(nums []int) int64 {
	totalPairs := (len(nums) * (len(nums) - 1)) / 2
	goodPairs := 0
	hashMap := map[int]int{}

	for index, value := range nums {
		diff := index - value
		goodPairs += hashMap[diff]
		hashMap[diff]++
	}

	return int64(totalPairs - goodPairs)

}

func main() {
	fmt.Println(countBadPairs([]int{4, 1, 3, 3}))
}
