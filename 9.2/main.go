package main

import (
	"fmt"
	"strconv"
	"strings"

	"main/util"
)

type point struct {
	x, y int
}

func updateNext(r []point, i int) {
	dx := r[i].x - r[i+1].x
	dy := r[i].y - r[i+1].y

	if util.AbsInt(dx) < 2 && util.AbsInt(dy) < 2 {
		return
	}

	if dx < 0 {
		// t is to the right of h
		r[i+1].x -= 1
	} else if dx > 0 {
		// t is to the left of h
		r[i+1].x += 1
	}

	if dy < 0 {
		// t is above h
		r[i+1].y -= 1
	} else if dy > 0 {
		// t is below h
		r[i+1].y += 1
	}
}

var debug = false

func printState(r []point, visited map[point]bool) {
	if !debug {
		return
	}

	// TODO: don't hardcode size

	for y := 15; y >= -6; y -= 1 {
		for x := -11; x <= 14; x += 1 {
			inRope := false
			for i, p := range r {
				if p.x == x && p.y == y {
					if i == 0 {
						fmt.Print("H")
					} else if i == len(r)-1 {
						fmt.Print("T")
					} else {
						fmt.Print(i)
					}

					inRope = true
					break
				}
			}

			if inRope {
				// Do nothing
			} else if visited[point{x, y}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}

		fmt.Println()
	}

	fmt.Println()
}

func main() {
	var (
		visited = make(map[point]bool)
		length  = 10
		rope    = make([]point, length)
	)

	visited[point{0, 0}] = true

	for _, line := range util.ReadInputLines("./input.txt") {
		printState(rope, visited)

		pair := strings.Split(line, " ")

		dir := pair[0]
		mag, err := strconv.Atoi(pair[1])
		util.Check(err)

		var dx, dy int

		switch dir {
		case "R":
			dx = 1
		case "L":
			dx = -1
		case "U":
			dy = 1
		case "D":
			dy = -1
		}

		for mag > 0 {
			rope[0].x += dx
			rope[0].y += dy

			for i := 0; i < length-1; i += 1 {
				updateNext(rope, i)
			}

			visited[rope[length-1]] = true

			printState(rope, visited)

			mag -= 1
		}
	}

	fmt.Printf("%d\n", len(visited))
}
