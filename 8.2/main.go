package main

import (
	"fmt"
	"main/util"
)

func main() {
	maxScore := 0
	grid := util.ParseIntGrid()

	for i := 1; i < len(grid)-1; i += 1 {
		for j := 1; j < len(grid[i])-1; j += 1 {
			cell := grid[i][j]

			var x int

			for x = j - 1; x > 0 && grid[i][x] < cell; x -= 1 {
			}
			left := j - x

			for x = j + 1; x < len(grid[i])-1 && grid[i][x] < cell; x += 1 {
			}
			right := x - j

			for x = i - 1; x > 0 && grid[x][j] < cell; x -= 1 {
			}
			up := i - x

			for x = i + 1; x < len(grid)-1 && grid[x][j] < cell; x += 1 {
			}
			down := x - i

			// fmt.Printf("(%d, %d) [%d] => %d = %d * %d * %d * %d\n", j, i, cell, left*right*up*down, left, right, up, down)
			maxScore = util.MaxInt(maxScore, left*right*up*down)
		}
	}

	fmt.Println(maxScore)
}
