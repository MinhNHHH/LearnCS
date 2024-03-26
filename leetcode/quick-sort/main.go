package main

import "fmt"

func main() {
	fmt.Println(QuickSort([]int{3, 5, 2, 1, 6}))
}

// Partitioning: Choose a pivot element from the array.
// Rearrange the elements of the array such that all elements less than the pivot come before it, and all elements greater than the pivot come after it.
// After this partitioning, the pivot element is in its correct sorted position.

// Recursion: Recursively apply the partitioning step to the sub-arrays formed by the partitioning step until the entire array is sorted.

// Base case: The base case of the recursion is when the sub-array has zero or one element, as a single element is considered sorted.

// Combine: Since the elements are rearranged in place, no explicit combining step is required. The array is sorted when the recursion unwinds.

func QuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	p := patition(nums)
	QuickSort(nums[:p])
	QuickSort(nums[p+1:])
	return nums
}

func patition(nums []int) int {
	pivot := nums[len(nums)-1]
	index := -1

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < pivot {
			index += 1
			nums[i], nums[index] = nums[index], nums[i]
		}
	}
	fmt.Println(index)
	nums[index+1], nums[len(nums)-1] = nums[len(nums)-1], nums[index+1]
	return index + 1
}
