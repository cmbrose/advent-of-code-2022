package main

import (
	"fmt"
	"main/util"
	"strings"
)

type point struct {
	x, y int
}

func simulateSand(grid *[][]rune, source *point) {
	p := *source

	tryStep := func() bool {
		nextRow := (*grid)[p.y+1]
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
		if p.x <= 1 {
			for y := 0; y < len(*grid); y += 1 {
				r := '.'
				if y == len(*grid)-1 {
					r = '#'
				}

				(*grid)[y] = append(make([]rune, 5), (*grid)[y]...)
				for x := 0; x < 5; x += 1 {
					(*grid)[y][x] = r
				}
			}

			p.x += 5
			source.x += 5
		}

		if p.x >= len((*grid)[0])-1 {
			for y := 0; y < len(*grid); y += 1 {
				r := '.'
				if y == len(*grid)-1 {
					r = '#'
				}

				(*grid)[y] = append((*grid)[y], make([]rune, 5)...)

				for x := 1; x <= 5; x += 1 {
					(*grid)[y][len((*grid)[0])-x] = r
				}
			}
		}
	}

	(*grid)[p.y][p.x] = 'o'
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

	maxY += 2

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
	for x := minX; x <= maxX; x += 1 {
		setPoint(x, maxY, '#')
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

	// util.PrintGrid(grid, "%c")
	// fmt.Println()

	source := point{500 - minX, 0}

	for grid[0][source.x] == '+' {
		simulateSand(&grid, &source)

		// util.PrintGrid(grid, "%c")
		// fmt.Println()

		cnt += 1
	}

	// util.PrintGrid(grid, "%c")
	// fmt.Println()

	fmt.Println(cnt)
}
