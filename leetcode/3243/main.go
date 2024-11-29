package main

import (
	"fmt"
)

func shortestDistanceAfterQueries(n int, queries [][]int) []int {
	graphs := make([][]int, n)
	for i := range graphs {
		graphs[i] = append(graphs[i], i+1)
	}

	var shortestPath func() int
	shortestPath = func() int {
		queue := [][]int{{0, 0}}
		visited := map[int]bool{}
		visited[0] = true
		for len(queue) > 0 {
			q := queue[0]
			queue = queue[1:]
			currNode, length := q[0], q[1]

			if currNode == n-1 {
				return length
			}

			for _, node := range graphs[currNode] {
				if !visited[node] {
					queue = append(queue, []int{node, length + 1})
					visited[node] = true
				}
			}
		}
		return -1
	}

	res := []int{}
	for _, query := range queries {
		start, end := query[0], query[1]
		graphs[start] = append(graphs[start], end)
		res = append(res, shortestPath())
	}

	return res
}

func main() {
	fmt.Println(shortestDistanceAfterQueries(4, [][]int{{0, 2}, {0, 3}}))
}
