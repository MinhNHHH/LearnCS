package main

import (
	"fmt"
	"sort"
)

func sum(arr []int) int {
	s := 0
	for _, num := range arr {
		s += num
	}
	return s
}
func dividePlayers(skill []int) int64 {
	n := len(skill)
	total := sum(skill)
	if total%(n/2) != 0 {
		return -1
	}

	hash := make([]int, n/2)

	sort.Ints(skill)
	total = 0
	i, j := 0, n-1
	for i <= j {
		hash[i] = skill[i] + skill[j]
		if i > 0 && hash[i] != hash[i-1] {
			return -1
		}
		total += skill[i] * skill[j]
		i++
		j--
	}

	return int64(total)
}

func main() {
	fmt.Println(dividePlayers([]int{1, 2, 2, 5}))
}
