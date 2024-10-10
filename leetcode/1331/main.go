package main

import (
	"fmt"
	"sort"
)

func arrayRankTransform(arr []int) []int {
	temp := make([]int, len(arr))
	copy(temp, arr)
	sort.Ints(arr)

	hash := map[int]int{}
	rank := 1
	for _, v := range arr {
		if _, ok := hash[v]; !ok {
			hash[v] = rank
			rank++
		}
	}

	for i, v := range temp {
		temp[i] = hash[v]
	}

	return temp
}

func main() {
	fmt.Println(arrayRankTransform([]int{37, 12, 28, 9, 100, 56, 80, 5, 12}))
}
