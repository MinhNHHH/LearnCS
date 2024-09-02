package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Ints(candidates)

	var backTracking func(candidates []int, remaining_nums []int, target int)
	backTracking = func(listCandidates []int, remaining_nums []int, target int) {

		if target == 0 {
			copied := make([]int, len(listCandidates))
			copy(copied, listCandidates)
			fmt.Println(copied)
			res = append(res, copied)
		} else {
			for i := 0; i < len(candidates); i++ {
				if i > 0 && remaining_nums[i] == remaining_nums[i-1] {
					continue
				}
				if remaining_nums[i] > target {
					break
				}
				listCandidates = append(listCandidates, remaining_nums[i])
				backTracking(listCandidates, remaining_nums[i+1:], target-remaining_nums[i])
				listCandidates = listCandidates[:len(listCandidates)-1]
			}
		}
	}
	backTracking([]int{}, candidates, target)
	return res
}
func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
