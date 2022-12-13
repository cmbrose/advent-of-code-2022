package main

import (
	"container/heap"
	"fmt"
	"main/util"
)

type node struct {
	point

	elev int

	neighbors []*node
}

type point struct {
	x, y int
}

func main() {
	var ex, ey int

	var nodes [][]*node

	targets := make(map[point]bool) // all the 'a' points

	for y, line := range util.ReadInputLines("./input.txt") {
		var row []*node

		for x, r := range line {
			elev := int(r - 'a')

			if r == 'S' {
				elev = 0
			} else if r == 'E' {
				elev = 25
				ex = x
				ey = y
			}

			n := &node{
				point{x, y},
				elev,
				nil,
			}

			if n.elev == 0 {
				targets[n.point] = true
			}

			if x > 0 {
				left := row[x-1]
				if left.elev-n.elev >= -1 {
					n.neighbors = append(n.neighbors, left)
				}
				if n.elev-left.elev >= -1 {
					left.neighbors = append(left.neighbors, n)
				}
			}

			if y > 0 {
				up := nodes[y-1][x]
				if up.elev-n.elev >= -1 {
					n.neighbors = append(n.neighbors, up)
				}
				if n.elev-up.elev >= -1 {
					up.neighbors = append(up.neighbors, n)
				}
			}

			row = append(row, n)
		}

		nodes = append(nodes, row)
	}

	scores := util.Grid[int](len(nodes[0]), len(nodes))
	seen := util.Grid[bool](len(nodes[0]), len(nodes))

	type scoredNode struct {
		*node
		score int
	}

	q := util.NewPriorityQueue(func(a, b scoredNode) bool {
		return a.score < b.score
	})

	end := nodes[ey][ex]
	heap.Push(q, scoredNode{end, 0})

	var targetCopy []point

	for q.Len() > 0 {
		n := heap.Pop(q).(scoredNode)

		if seen[n.y][n.x] {
			continue
		}

		scores[n.y][n.x] = n.score
		seen[n.y][n.x] = true

		if _, ok := targets[n.point]; ok {
			delete(targets, n.point)
			targetCopy = append(targetCopy, n.point)

			if len(targets) == 0 {
				break
			}
		}

		for _, o := range n.neighbors {
			heap.Push(q, scoredNode{o, n.score + 1})
		}
	}

	min := util.MinInt(util.Map(targetCopy, func(p point) int {
		return scores[p.y][p.x]
	})...)

	fmt.Println(min)
}
