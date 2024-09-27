package main

import (
	"fmt"
	"strconv"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func convertListIntToListString(nums []int) []string {
	res := []string{}
	for _, value := range nums {
		res = append(res, strconv.Itoa(value))
	}
	return res
}

func mergeSort(strings []string) []string {
	if len(strings) <= 1 {
		return strings
	}

	mid := len(strings) / 2
	left := mergeSort(strings[:mid])
	right := mergeSort(strings[mid:])

	return merge(left, right)
}

func merge(left, right []string) []string {
	sorted := make([]string, len(left)+len(right))
	l, r, k := 0, 0, 0

	for l < len(left) && r < len(right) {
		// Compare concatenated results
		if left[l]+right[r] > right[r]+left[l] {
			sorted[k] = left[l]
			l++
		} else {
			sorted[k] = right[r]
			r++
		}
		k++
	}

	// Copy remaining elements
	for l < len(left) {
		sorted[k] = left[l]
		l++
		k++
	}

	for r < len(right) {
		sorted[k] = right[r]
		r++
		k++
	}

	return sorted
}

func sumNums(nums []int) int {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	return sum
}
func largestNumber(nums []int) string {
	if sumNums(nums) == 0 {
		res := strconv.Itoa(nums[0])
		return res
	}
	strs := convertListIntToListString(nums)
	strs = mergeSort(strs)
	res := ""
	for _, str := range strs {
		res = res + str
	}
	return res
}

func main() {
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
}
