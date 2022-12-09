package main

import (
	"fmt"
	"main/util"
)

func main() {
	grid := util.ParseIntGrid()

	// These are the heights of the tallest tree seen so far looking
	// from the given direction, not including the current cell
	//
	// Grid:     Left:
	// 30373     X3337
	// 25512     X2555
	// 65332     X6666
	// 33549     X3355
	// 35390     X3559
	var left, right, up, down [][]int

	util.PrintIntGrid(grid)
	fmt.Println()

	for _, row := range grid {
		left = append(left, make([]int, len(row)))
		right = append(right, make([]int, len(row)))
		up = append(up, make([]int, len(row)))
		down = append(down, make([]int, len(row)))
	}

	for i, row := range grid {
		i2 := len(grid) - 1 - i
		for j := range row {
			j2 := len(row) - 1 - j

			if j == 0 {
				left[i][j] = -1
				right[i][j2] = -1
			} else {
				left[i][j] = util.MaxInt(left[i][j-1], grid[i][j-1])
				right[i][j2] = util.MaxInt(right[i][j2+1], grid[i][j2+1])
			}

			if i == 0 {
				up[i][j] = -1
				down[i2][j] = -1
			} else {
				up[i][j] = util.MaxInt(up[i-1][j], grid[i-1][j])
				down[i2][j] = util.MaxInt(down[i2+1][j], grid[i2+1][j])
			}
		}
	}

	util.PrintIntGrid(left)
	fmt.Println()
	util.PrintIntGrid(right)
	fmt.Println()
	util.PrintIntGrid(up)
	fmt.Println()
	util.PrintIntGrid(down)
	fmt.Println()

	cnt := 0

	for i := range grid {
		for j := range grid[i] {
			cell := grid[i][j]

			if cell > left[i][j] || cell > right[i][j] || cell > up[i][j] || cell > down[i][j] {
				cnt += 1
			}
		}
	}

	fmt.Println(cnt)
}
