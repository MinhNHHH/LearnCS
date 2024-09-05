package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	director := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	temp := [][]int{{sr, sc}}
	visited := make([][]bool, len(image))
	for i := range visited {
		visited[i] = make([]bool, len(image[0]))
	}

	for len(temp) > 0 {
		current := temp[0]
		temp = temp[1:]
		visited[current[0]][current[1]] = true
		for _, d := range director {
			dr, dc := current[0]+d[0], current[1]+d[1]
			fmt.Println(dr, dc)
			if dr >= 0 && dr < len(image) && dc >= 0 && dc < len(image[0]) && image[current[0]][current[1]] == image[dr][dc] && !visited[dr][dc] {
				temp = append(temp, []int{dr, dc})
			}
		}
	}

	for r := 0; r < len(visited); r++ {
		for c := 0; c < len(visited[0]); c++ {
			if visited[r][c] {
				image[r][c] = color
			}
		}
	}
	return image
}

func main() {
	fmt.Println(floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2))
}
