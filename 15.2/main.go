package main

import (
	"fmt"
	"main/util"
	"os"
	"regexp"
)

type point struct {
	x, y int
}

func main() {
	var (
		maxX = 4000000
		minX = 0
		maxY = 4000000
		minY = 0
	)

	re := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)

	sensors := make(map[point]int)
	beacons := make(map[point]bool)

	for _, line := range util.ReadInputLines("input.txt") {
		match := re.FindStringSubmatch(line)

		var (
			sx = util.AssertInt(match[1])
			sy = util.AssertInt(match[2])
			bx = util.AssertInt(match[3])
			by = util.AssertInt(match[4])
		)

		dist := util.AbsInt(sx-bx) + util.AbsInt(sy-by)

		sensors[point{sx, sy}] = dist
		beacons[point{bx, by}] = true
	}

	for s, dist := range sensors {
		points := []point{
			{s.x, s.y - dist - 1}, // Top
			{s.x - dist - 1, s.y}, // Left
			{s.x, s.y + dist + 1}, // Bottom
			{s.x + dist + 1, s.y}, // Right
			{s.x, s.y - dist - 1}, // Top again for a loop
		}

		for i := 0; i < 4; i += 1 {
			p1, p2 := points[i], points[i+1]

			util.Step(p1.x, p1.y, p2.x, p2.y, func(x, y int) {
				if x < minX || x > maxX || y < minY || y > maxY {
					return
				}

				if beacons[point{x, y}] {
					return
				}

				for s, dist := range sensors {
					test := util.AbsInt(x-s.x) + util.AbsInt(y-s.y)

					if test <= dist {
						return
					}
				}

				fmt.Printf("%d (%d, %d)\n", x*4000000+y, x, y)
				os.Exit(0)
			})
		}
	}
}
