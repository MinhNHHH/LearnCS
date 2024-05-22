// There is an undirected graph with n nodes, where each node is numbered between 0 and n - 1.
// You are given a 2D array graph, where graph[u] is an array of nodes that node u is adjacent to.
// More formally, for each v in graph[u], there is an undirected edge between node u and node v. The graph has the following properties:

// There are no self-edges (graph[u] does not contain u).
// There are no parallel edges (graph[u] does not contain duplicate values).
// If v is in graph[u], then u is in graph[v] (the graph is undirected).
// The graph may not be connected, meaning there may be two nodes u and v such that there is no path between them.

// A graph is bipartite if the nodes can be partitioned into two independent sets A and B such that every edge in the graph connects a node in set A and a node in set B.

// Return true if and only if it is bipartite.

// Input: graph = [[1,2,3],[0,2],[0,1,3],[0,2]]
// Output: false
// Explanation: There is no way to partition the nodes into two independent sets such that every edge connects a node in one and a node in the other.

// Input: graph = [[1,3],[0,2],[1,3],[0,2]]
// Output: true
// Explanation: We can partition the nodes into two sets: {0, 2} and {1, 3}.

package main

import "fmt"

func main() {
	fmt.Println(isBipartite([][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}))
	fmt.Println(isBipartite([][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}))
}

func bfs(graph [][]int, node int, color []int) bool {
	queue := []int{node}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, neighboor := range graph[current] {
			if color[neighboor] == -1 {
				color[neighboor] = 1 + color[current]
				queue = append(queue, neighboor)
			} else if color[neighboor] == color[current] {
				return false
			}

		}
	}
	return true
}
func isBipartite(graph [][]int) bool {
	// Initialize a array color
	color := make([]int, len(graph))
	for i := 0; i < len(graph); i++ {
		// -1 is uncolor
		color[i] = -1
	}

	for n := 0; n < len(graph); n++ {
		if color[n] == -1 {
			if !bfs(graph, n, color) {
				return false
			}
		}
	}

	return true
}
