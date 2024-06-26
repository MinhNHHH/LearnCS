// Given a reference of a node in a connected undirected graph.

// Return a deep copy (clone) of the graph.

// Each node in the graph contains a value (int) and a list (List[Node]) of its neighbors.

// class Node {
//     public int val;
//     public List<Node> neighbors;
// }

// Test case format:

// For simplicity, each node's value is the same as the node's index (1-indexed). For example, the first node with val == 1, the second node with val == 2, and so on. The graph is represented in the test case using an adjacency list.

// An adjacency list is a collection of unordered lists used to represent a finite graph. Each list describes the set of neighbors of a node in the graph.

// The given node will always be the first node with val = 1. You must return the copy of the given node as a reference to the cloned graph.

// Input: adjList = [[2,4],[1,3],[2,4],[1,3]]
// Output: [[2,4],[1,3],[2,4],[1,3]]
// Explanation: There are 4 nodes in the graph.
// 1st node (val = 1)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
// 2nd node (val = 2)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).
// 3rd node (val = 3)'s neighbors are 2nd node (val = 2) and 4th node (val = 4).
// 4th node (val = 4)'s neighbors are 1st node (val = 1) and 3rd node (val = 3).

package main

func main() {}

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraphDFS(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[*Node]*Node)

	var dfs func(node *Node) *Node

	dfs = func(node *Node) *Node {
		if clonedNode, found := visited[node]; found {
			return clonedNode
		}
		copyNode := &Node{Val: node.Val, Neighbors: []*Node{}}
		visited[copyNode] = copyNode

		for _, neighboor := range node.Neighbors {
			copyNode.Neighbors = append(copyNode.Neighbors, dfs(neighboor))
		}
		return copyNode
	}
	return dfs(node)
}

func cloneGraphBFS(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[*Node]*Node)
	queue := []*Node{node}

	visited[node] = &Node{Val: node.Val, Neighbors: []*Node{}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for _, neighboor := range current.Neighbors {
			if _, found := visited[neighboor]; !found {
				// Copy a new Node
				visited[neighboor] = &Node{Val: neighboor.Val, Neighbors: []*Node{}}
				queue = append(queue, neighboor)
			}
			visited[current].Neighbors = append(visited[current].Neighbors, visited[neighboor])
		}
	}
	return visited[node]
}
