package main

import (
	"fmt"
)

func check(board [][]int) int {
	for _, row := range board {
		for _, cell := range row {
			fmt.Println(cell)
		}
	}
	return 0
}

func main() {
	board := [][]int{{1, 1, 1}, {2, 2, 0}, {2, 0, 1}}
	check(board)
}
