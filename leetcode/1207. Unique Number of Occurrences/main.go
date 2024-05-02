// Given an array of integers arr, return true if the number of occurrences of each value in the array is unique or false otherwise.

// Example 1:

// Input: arr = [1,2,2,1,1,3]
// Output: true
// Explanation: The value 1 has 3 occurrences, 2 has 2 and 3 has 1. No two values have the same number of occurrences.

// Example 2:

// Input: arr = [1,2]
// Output: false

// Example 3:

// Input: arr = [-3,0,1,-3,1,1,1,-3,10,0]
// Output: true

package main

func main() {}

func uniqueOccurrences(arr []int) bool {
	hashMap := map[int]int{}
	s1 := map[int]bool{}

	for _, value := range arr {
		hashMap[value]++
	}

	for _, value := range hashMap {
		if _, ok := s1[value]; ok {
			return false
		}
		s1[value] = true
	}
	return true
}
