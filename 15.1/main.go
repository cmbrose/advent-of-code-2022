package main

import (
	"fmt"
	"main/util"
	"math"
	"regexp"
)

type point struct {
	x, y int
}

func main() {
	var (
		maxX = math.MinInt
		minX = math.MaxInt
		maxY = math.MinInt
		minY = math.MaxInt
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

		minX = util.MinInt(minX, sx-dist)
		maxX = util.MaxInt(maxX, sx+dist)
		minY = util.MinInt(minY, sy-dist)
		maxY = util.MaxInt(maxY, sy+dist)
	}

	cnt := 0
	targetY := 2_000_000

	p := point{minX, targetY}

	for ; p.x <= maxX; p.x += 1 {
		if beacons[p] {
			continue
		}

		for s, dist := range sensors {
			test := util.AbsInt(p.x-s.x) + util.AbsInt(p.y-s.y)

			if test <= dist {
				cnt += 1
				break
			}
		}
	}

	fmt.Println(cnt)
}
