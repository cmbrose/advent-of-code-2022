package main

import (
	"container/heap"
	"fmt"
	"main/util"
)

type node struct {
	elev int

	x, y int

	neighbors []*node
}

func printGrid(grid [][]*node) {
	for _, row := range grid {
		var top, mid, bot string

		for _, n := range row {
			var u, d, l, r bool

			for _, o := range n.neighbors {
				u = u || o.y < n.y
				d = d || o.y > n.y
				l = l || o.x < n.x
				r = r || o.x > n.x
			}

			if u {
				top += " ^ "
			} else {
				top += "   "
			}

			if l {
				mid += "<"
			} else {
				mid += " "
			}

			mid += string('a' + n.elev)

			if r {
				mid += ">"
			} else {
				mid += " "
			}

			if d {
				bot += " v "
			} else {
				bot += "   "
			}
		}

		fmt.Println(top)
		fmt.Println(mid)
		fmt.Println(bot)
	}
}

func main() {
	var sx, sy, ex, ey int

	var nodes [][]*node

	for y, line := range util.ReadInputLines("./input.txt") {
		var row []*node

		for x, r := range line {
			elev := int(r - 'a')

			if r == 'S' {
				elev = 0
				sx = x
				sy = y
			} else if r == 'E' {
				elev = 25
				ex = x
				ey = y
			}

			n := &node{
				elev,
				x, y,
				nil,
			}

			if x > 0 {
				left := row[x-1]
				if left.elev-n.elev >= -1 {
					left.neighbors = append(left.neighbors, n)
				}
				if n.elev-left.elev >= -1 {
					n.neighbors = append(n.neighbors, left)
				}
			}

			if y > 0 {
				up := nodes[y-1][x]
				if up.elev-n.elev >= -1 {
					up.neighbors = append(up.neighbors, n)
				}
				if n.elev-up.elev >= -1 {
					n.neighbors = append(n.neighbors, up)
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

	start := nodes[sy][sx]
	heap.Push(q, scoredNode{start, 0})

	for q.Len() > 0 {
		n := heap.Pop(q).(scoredNode)

		if seen[n.y][n.x] {
			continue
		}

		scores[n.y][n.x] = n.score
		seen[n.y][n.x] = true

		if n.y == ey && n.x == ex {
			break
		}

		for _, o := range n.neighbors {
			heap.Push(q, scoredNode{o, n.score + 1})
		}
	}

	fmt.Println(scores[ey][ex])
}
