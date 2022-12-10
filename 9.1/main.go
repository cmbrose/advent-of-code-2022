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

func updateTail(h, t *point) {
	dx := h.x - t.x
	dy := h.y - t.y

	if util.AbsInt(dx) < 2 && util.AbsInt(dy) < 2 {
		return
	}

	if dx < 0 {
		// t is to the right of h
		t.x -= 1
	} else if dx > 0 {
		// t is to the left of h
		t.x += 1
	}

	if dy < 0 {
		// t is above h
		t.y -= 1
	} else if dy > 0 {
		// t is below h
		t.y += 1
	}
}

var debug = false

func printState(h, t point, visited map[point]bool) {
	if !debug {
		return
	}

	// TODO: don't hardcode size

	for y := 6; y >= 0; y -= 1 {
		for x := 0; x <= 6; x += 1 {
			if h.x == x && h.y == y {
				fmt.Print("H")
			} else if t.x == x && t.y == y {
				fmt.Print("T")
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
		h       = point{0, 0}
		t       = point{0, 0}
	)

	visited[t] = true

	for _, line := range util.ReadInputLines("./input.txt") {
		printState(h, t, visited)

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
			h.x += dx
			h.y += dy

			updateTail(&h, &t)
			visited[t] = true

			printState(h, t, visited)

			mag -= 1
		}
	}

	fmt.Printf("%d\n", len(visited))
}
