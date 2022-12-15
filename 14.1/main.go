package main

import (
	"fmt"
	"main/util"
	"strings"
)

type point struct {
	x, y int
}

func simulateSand(grid [][]rune, p point) bool {
	tryStep := func() bool {
		nextRow := grid[p.y+1]
		if nextRow[p.x] == '.' {
			p.y += 1
			return true
		}
		if p.x == 0 || nextRow[p.x-1] == '.' {
			p.x -= 1
			p.y += 1
			return true
		}
		if p.x+1 == len(nextRow) || nextRow[p.x+1] == '.' {
			p.x += 1
			p.y += 1
			return true
		}
		return false
	}

	for tryStep() {
		if p.y == len(grid)-1 || p.x < 0 || p.x == len(grid[0]) {
			return false
		}
	}

	grid[p.y][p.x] = 'o'

	return true
}

func main() {
	cnt := 0

	var (
		minX = 500
		maxX = 500
		minY = 0
		maxY = 0
	)

	var vectors [][]point

	for _, line := range util.ReadInputLines("input.txt") {
		pointStrs := strings.Split(line, " -> ")

		points := util.Map(pointStrs, func(str string) point {
			pair := strings.Split(str, ",")

			x, y := util.AssertInt(pair[0]), util.AssertInt(pair[1])

			minX, maxX = util.MinInt(minX, x), util.MaxInt(maxX, x)
			minY, maxY = util.MinInt(minY, y), util.MaxInt(maxY, y)

			return point{x, y}
		})

		vectors = append(vectors, points)
	}

	var (
		width  = maxX - minX + 1
		height = maxY - minY + 1
	)

	grid := util.FillGrid(width, height, '.')
	// getPoint := func(x, y int) rune {
	// 	return grid[y-minY][x-minX]
	// }
	setPoint := func(x, y int, r rune) {
		grid[y-minY][x-minX] = r
	}

	setPoint(500, 0, '+')

	for _, vector := range vectors {
		for i := 0; i < len(vector)-1; i += 1 {
			c, n := vector[i], vector[i+1]

			util.Step(c.x, c.y, n.x, n.y, func(x, y int) {
				setPoint(x, y, '#')
			})
		}
	}

	util.PrintGrid(grid, "%c")
	fmt.Println()

	for simulateSand(grid, point{500 - minX, 0}) {
		// util.PrintGrid(grid, "%c")
		// fmt.Println()

		cnt += 1
	}

	util.PrintGrid(grid, "%c")
	fmt.Println()

	fmt.Println(cnt)
}
