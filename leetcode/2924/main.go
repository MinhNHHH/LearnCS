package main

import "fmt"

func findChampion(n int, edges [][]int) int {
	set := map[int]bool{}
	for i := 0; i < n; i++ {
		set[i] = true
	}

	for _, edge := range edges {
		delete(set, edge[1])
	}
	if len(set) > 1 {
		return -1
	}
	for key, value := range set {
		if value {
			return key
		}
	}
	return -1
}

func main() {
	fmt.Println(findChampion(3, [][]int{{0, 1}, {1, 2}}))
	fmt.Println(findChampion(4, [][]int{{0, 2}, {1, 3}, {1, 2}}))

}
