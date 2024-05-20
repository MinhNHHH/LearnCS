// Number of Provinces
// There are n cities. Some of them are connected, while some are not. If city a is connected directly with city b, and city b is connected directly with city c, then city a is connected indirectly with city c.

// A province is a group of directly or indirectly connected cities and no other cities outside of the group.

// You are given an n x n matrix isConnected where isConnected[i][j] = 1 if the ith city and the jth city are directly connected, and isConnected[i][j] = 0 otherwise.

// Return the total number of provinces.

// Input: isConnected = [[1,1,0],[1,1,0],[0,0,1]]
// Output: 2

// Input: isConnected = [[1,0,0],[0,1,0],[0,0,1]]
// Output: 3

package main

import "fmt"

func main() {
	fmt.Println(findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}}))
}

func findCircleNum(isConnected [][]int) int {
	visited := make([]bool, len(isConnected))
	provinces := 0

	var dfs = func(city int) {}
	dfs = func(city int) {
		for neighboor, isConnect := range isConnected[city] {
			if isConnect == 1 && !visited[neighboor] {
				visited[neighboor] = true
				dfs(neighboor)
			}
		}
	}

	for city := 0; city < len(isConnected); city++ {
		if !visited[city] {
			dfs(city)
			provinces++
		}
	}

	return provinces
}
