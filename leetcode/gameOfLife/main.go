package main

// 0 1 0   0 0 0
// 0 0 1   1 0 1
// 1 1 1   0 1 1
// 0 0 0   0 1 0

// board = [[0,1,0],[0,0,1],[1,1,1],[0,0,0]]
// [[0,0,0],[1,0,1],[0,1,1],[0,1,0]]

// 1 - live if 1 in [2,3] else die
// 0 - live if count "1" == 3 else die

// Any live cell with fewer than two live neighbors dies as if caused by under-population.
// Any live cell with two or three live neighbors lives on to the next generation.
// Any live cell with more than three live neighbors dies, as if by over-population.
// Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

func countNeighbors(board [][]int, row, col int) int {
	neighbor := 0
	for r := row - 1; r < row+2; r++ {
		for c := col - 1; c < col+2; c++ {
			if r < 0 || c < 0 || (r == row && c == col) || r == len(board) || c == len(board[0]) {
				continue
			}
			if board[r][c] == 1 || board[r][c] == 3 {
				neighbor += 1
			}

		}
	}
	return neighbor
}

func gameOfLife(board [][]int) {
	// Original| New | State
	// 0       | 0   | 0
	// 0       | 0   | 0
	// 1       | 1   | 2
	// 1       | 1   | 3

	rows := len(board)
	cols := len(board[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			neighbor := countNeighbors(board, r, c)
			if board[r][c] > 0 {
				if neighbor >= 2 && neighbor <= 3 {
					board[r][c] = 3
				}
			} else if neighbor == 3 {
				board[r][c] = 2
			}
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] >= 2 && board[r][c] <= 3 {
				board[r][c] = 1
			} else if board[r][c] == 1 {
				board[r][c] = 0
			}
		}
	}
}

func main() {
	gameOfLife([][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}})
}
