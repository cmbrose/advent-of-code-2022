package main

import (
	"fmt"
	"main/util"
	"strings"
)

const pitWidth = 7

type point struct {
	x, y int
}

type rock struct {
	shape [][]byte

	w, h int

	l, r, b []point
}

func makeRock(shape [][]byte) rock {
	return rock{
		shape,
		len(shape[0]),
		len(shape),
		getSide(shape, '<'),
		getSide(shape, '>'),
		getBottom(shape),
	}
}

func getSide(r [][]byte, dir byte) []point {
	var ps []point

	for y := 0; y < len(r); y += 1 {
		if dir == '>' {
			for x := len(r[y]) - 1; x >= 0; x -= 1 {
				if r[y][x] == '#' {
					ps = append(ps, point{x, y})
					break
				}
			}
		} else {
			for x := 0; x < len(r[y]); x += 1 {
				if r[y][x] == '#' {
					ps = append(ps, point{x, y})
					break
				}
			}
		}
	}

	return ps
}

func getBottom(r [][]byte) []point {
	minPoint := make(map[int]int) // x -> y

	for y := len(r) - 1; y >= 0; y -= 1 {
		for x := 0; x < len(r[y]); x += 1 {
			if _, ok := minPoint[x]; ok {
				continue
			}

			if r[y][x] == '#' {
				minPoint[x] = y
			}
		}

		if len(minPoint) == len(r[0]) {
			break
		}
	}

	var ps []point
	for x, y := range minPoint {
		ps = append(ps, point{x, y})
	}
	return ps
}

func main() {
	wind := []byte(util.ReadInputLines("./input.txt")[0])

	rockBlocks := util.ReadInputBlocks("./rocks.txt")
	rocks := util.Map(rockBlocks, func(lines []string) rock {
		shape := util.Map(lines, func(line string) []byte {
			return []byte(line)
		})

		r := makeRock(shape)

		//fmt.Println(lines, r)
		return r
	})

	rockCursor := 0
	nextRock := func() rock {
		r := rocks[rockCursor]
		rockCursor = (rockCursor + 1) % len(rocks)
		return r
	}

	windCursor := 0
	nextWind := func() byte {
		w := wind[windCursor]
		windCursor = (windCursor + 1) % len(wind)
		return w
	}

	top := 0
	pit := [][]byte{}

	target := 2022

	for i := 0; i < target; i += 1 {
		r := nextRock()
		requiredHeight := top + 3 + r.h

		for height := len(pit); height < requiredHeight; height += 1 {
			pit = append(pit, emptyRow())
		}

		testTop := fall(r, requiredHeight-1, pit, nextWind)
		top = util.MaxInt(top, testTop)
		//printPit(pit)
	}

	fmt.Println(top)
}

func printPit(pit [][]byte) {
	var bldr strings.Builder

	for y := len(pit) - 1; y >= 0; y -= 1 {
		bldr.WriteString(fmt.Sprintf("%4d ", y))
		for _, cell := range pit[y] {
			bldr.WriteRune(rune(cell))
		}
		bldr.WriteRune('\n')
	}

	fmt.Println(bldr.String())
}

func emptyRow() []byte {
	row := make([]byte, pitWidth)
	for i := range row {
		row[i] = '.'
	}
	return row
}

// Returns the new highest y
func fall(r rock, y int, pit [][]byte, nextWind func() byte) int {
	x := 2

	//fmt.Printf("START: (%d, %d)\n", x, y)

	// y always points to the TOP of the rock

	for {
		wind := nextWind()
		//fmt.Printf("(%d, %d) %c\n", x, y, wind)

		edge, dir := r.r, 1
		if wind == '<' {
			edge, dir = r.l, -1
		}

		canMove := true
		for _, p := range edge {
			testX := x + p.x + dir

			if testX < 0 || testX >= pitWidth {
				canMove = false
				break
			}

			if pit[y-p.y][testX] == '#' {
				canMove = false
				break
			}
		}

		if canMove {
			//fmt.Printf("Move %c\n", wind)
			x += dir
		}

		if y == 0 {
			// On the ground
			break
		}

		canFall := true
		for _, p := range r.b {
			testY := y - p.y - 1

			if pit[testY][x+p.x] == '#' {
				canFall = false
				break
			}
		}

		if canFall {
			y -= 1
		} else {
			break
		}
	}

	//fmt.Printf("END: (%d, %d)\n", x, y)

	// Lock in the rock
	for rY, row := range r.shape {
		for rX, c := range row {
			if c == '.' {
				continue
			}

			if pit[y-rY][x+rX] == '#' {
				printPit(pit)
				fmt.Println(r)
				panic(fmt.Sprintf("Invalid placement into (%d, %d) (%d, %d)", x+rX, y-rY, rX, rY))
			}

			pit[y-rY][x+rX] = '#'
		}
	}

	return y + 1
}
