package main

import (
	"fmt"
)

// Graph represents an undirected graph using an adjacency list
type Graph struct {
	adjacencyList map[int][]int
}

// NewGraph creates a new Graph
func NewGraph() *Graph {
	return &Graph{adjacencyList: make(map[int][]int)}
}

// AddNode adds a node to the graph
func (g *Graph) AddNode(node int) {
	if _, exists := g.adjacencyList[node]; !exists {
		g.adjacencyList[node] = []int{}
	}
}

// AddEdge adds an edge between two nodes (undirected)
func (g *Graph) AddEdge(node1, node2 int) {
	g.adjacencyList[node1] = append(g.adjacencyList[node1], node2)
	g.adjacencyList[node2] = append(g.adjacencyList[node2], node1)
}

// PrintGraph prints the adjacency list of the graph
func (g *Graph) PrintGraph() {
	for node, neighbors := range g.adjacencyList {
		fmt.Printf("%d: %v\n", node, neighbors)
	}
}

func main() {
	graph := NewGraph()

	// Adding nodes
	graph.AddNode(1)
	graph.AddNode(2)
	graph.AddNode(3)
	graph.AddNode(4)

	// Adding edges
	graph.AddEdge(1, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(2, 4)
	graph.AddEdge(3, 4)

	// Print the graph
	graph.PrintGraph()
}
