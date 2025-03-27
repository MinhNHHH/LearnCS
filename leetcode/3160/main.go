package main

import "fmt"

func queryResults(limit int, queries [][]int) []int {
	collorTracking := map[int]int{}
	ballTracking := map[int]int{}

	res := []int{}

	for _, query := range queries {
		ball, color := query[0], query[1]

		if _, exit := ballTracking[ball]; exit {
			pre_color := ballTracking[ball]
			collorTracking[pre_color] -= 1

			if collorTracking[pre_color] == 0 {
				delete(collorTracking, pre_color)
			}
		}

		ballTracking[ball] = color
		collorTracking[color]++
		res = append(res, len(collorTracking))

	}

	return res
}

func main() {
	// fmt.Println(queryResults(4, [][]int{{1, 4}, {2, 5}, {1, 3}, {3, 4}}))
	fmt.Println(queryResults(4, [][]int{{0, 1}, {1, 2}, {2, 2}, {3, 4}, {4, 5}}))
}
