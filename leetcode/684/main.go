package main

import "fmt"

func findRedundantConnection(edges [][]int) []int {
	graph := make([]int, len(edges)+1)
	findTailNode := func(node int) int {
		for graph[node] != 0 {
			node = graph[node]
		}
		return node
	}
	for _, edge := range edges {
		tail1, tail2 := findTailNode(edge[0]), findTailNode(edge[1])
		if tail1 == tail2 {
			return edge
		}
		graph[tail1] = tail2
	}
	return []int{}
}

func main() {
	fmt.Println(findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 3}}))
}
