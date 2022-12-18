package main

import (
	"container/heap"
	"fmt"
	"main/util"
	"strings"
)

type rock uint32

const fullRow byte = 0b01111111

func (r rock) isFarLeft() bool {
	return r&0b01000000010000000100000001000000 != 0
}

func (r rock) shiftLeft() rock {
	return (r & 0b10000000100000001000000010000000) | // height checkers
		(r&0b00111111001111110011111100111111)<<1 // actual shape
}

func (r rock) isFarRight() bool {
	return r&0b00000001000000010000000100000001 != 0
}

func (r rock) shiftRight() rock {
	return (r & 0b10000000100000001000000010000000) | // height checkers
		(r&0b01111110011111100111111001111110)>>1 // actual shape
}

func (r rock) height() int {
	return int((r & (1 << 31) >> 31) + (r & (1 << 23) >> 23) + (r & (1 << 15) >> 15) + (r & (1 << 7) >> 7))
}

func (r rock) getRow(row int) byte {
	return byte((r >> (8 * (3 - row))) & 0b01111111)
}

func makeRock(shape [][]byte) rock {
	var r rock
	for y := 0; y < 4; y += 1 {
		r = r << 8
		if y >= len(shape) {
			continue
		}

		for x := 0; x < len(shape[y]); x += 1 {
			if shape[y][x] == '.' {
				continue
			}
			r |= 1 << (4 - x)
		}

		if r&0b01111111 != 0 {
			r |= 1 << 7
		}
	}
	return r
}

func main() {
	wind := []byte(util.ReadInputLines("./input.txt")[0])

	rockBlocks := util.ReadInputBlocks("./rocks.txt")
	rocks := util.Map(rockBlocks, func(lines []string) rock {
		shape := util.Map(lines, func(line string) []byte {
			return []byte(line)
		})

		r := makeRock(shape)
		//fmt.Printf("%v, %032b, %d\n", lines, r, r.height())
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

	truncAmount := 0
	top := 0
	pit := []byte{}

	type state struct {
		pit string
		rc  int
		wc  int
	}
	type nextState struct {
		state
		top int
	}
	cache := make(map[state]nextState)

	// startState := state{
	// 	pit: hex.EncodeToString(pit),
	// 	rc:  rockCursor,
	// 	wc:  windCursor,
	// }

	cacheHits := 0
	target := 2022

	for y := 0; y < target; y += 1 {
		if y%100 == 0 {
			fmt.Println(y, truncAmount, top, len(cache), cacheHits)
			cacheHits = 0
		}

		// if cached, ok := cache[startState]; ok && false {
		// 	cacheHits += 1
		// 	//fmt.Println("Found cached")
		// 	windCursor = cached.wc
		// 	rockCursor = cached.rc
		// 	top = cached.top
		// 	newPit, _ := hex.DecodeString(cached.pit)

		// 	if len(newPit) < len(pit) {
		// 		truncAmount += len(pit) - len(newPit)
		// 	}
		// 	pit = newPit

		// 	startState = cached.state
		// 	continue
		// }

		if y%100 == 0 {
			trimY := searchForTrim(pit)
			if trimY > 0 {
				truncAmount += trimY + 1
				top -= (trimY + 1)
				pit = pit[trimY+1:]
			}
		}

		r := nextRock()
		requiredHeight := top + 3 + r.height()

		for height := len(pit); height < requiredHeight; height += 1 {
			pit = append(pit, 0)
		}

		testTop := fall(r, requiredHeight-1, pit, nextWind)
		top = util.MaxInt(top, testTop)

		// endState := nextState{
		// 	state: state{
		// 		pit: hex.EncodeToString(pit),
		// 		rc:  rockCursor,
		// 		wc:  windCursor,
		// 	},
		// 	top: top,
		// }

		// cache[startState] = endState
		// startState = endState.state
		//printPit(pit, 0, 0)
	}
	fmt.Println(target, truncAmount, top, len(cache), cacheHits)

	//printPit(pit, 0, 0)
	fmt.Println(top + truncAmount)
}

func printPit(pit []byte, r rock, rY int) {
	var bldr strings.Builder

	for y := len(pit) - 1; y >= 0; y -= 1 {
		bldr.WriteString(fmt.Sprintf("%4d ", y))

		var rockRow byte
		rockRowY := rY - y
		if rockRowY >= 0 && rockRowY < r.height() {
			rockRow = r.getRow(rockRowY)
		}

		for i := 6; i >= 0; i -= 1 {
			c := '.'
			if pit[y]&(1<<i) != 0 {
				c = '#'
			}
			if rockRow&(1<<i) != 0 {
				c = '@'
			}

			bldr.WriteRune(c)
		}

		bldr.WriteRune('\n')
	}

	fmt.Println(bldr.String())
}

func searchForTrim(pit []byte) int {
	var y int
	for y = len(pit) - 1; y > 0; y -= 1 {
		if pit[y]&0b01000000 != 0 {
			break
		}
	}

	if y <= 0 {
		return 0
	}

	type point struct{ x, y int }

	dist := make(map[point]int)
	prev := make(map[point]point)
	seen := make(map[point]bool)

	q := util.NewPriorityQueue(func(a, b point) bool {
		if a.x > b.x {
			return true
		}
		return a.y > b.y
	})

	start := point{0, y}
	end := point{-1, -1}

	dist[start] = 0
	heap.Push(q, start)

	for q.Len() > 0 {
		p := q.Pop().(point)

		if seen[p] {
			continue
		}
		seen[p] = true

		x, y := p.x, p.y
		if x == 7 {
			end = p
			break
		}

		alt := dist[p] + 1

		up := point{x, y + 1}
		if !seen[up] && up.y < len(pit) && pit[up.y]&(1<<(7-up.x)) != 0 {
			if d, ok := dist[up]; !ok || d > alt {
				dist[up] = alt
				prev[up] = p
			}
			heap.Push(q, up)
		}

		right := point{x + 1, y}
		if !seen[right] && pit[right.y]&(1<<(7-right.x)) != 0 {
			if d, ok := dist[right]; !ok || d > alt {
				dist[right] = alt
				prev[right] = p
			}
			heap.Push(q, right)
		}

		down := point{x, y - 1}
		if !seen[down] && down.y > 0 && pit[down.y]&(1<<(7-down.x)) != 0 {
			if d, ok := dist[down]; !ok || d > alt {
				dist[down] = alt
				prev[down] = p
			}
			heap.Push(q, down)
		}
	}

	if end.x == -1 {
		return 0
	}

	minY := end.y
	for end, ok := prev[end]; ok; end, ok = prev[end] {
		if end.y < minY {
			minY = end.y
		}
	}

	return minY
}

// Returns the new highest y
func fall(r rock, y int, pit []byte, nextWind func() byte) int {
	//fmt.Printf("START: %d\n", y)
	//printPit(pit, r, y)

	// y always points to the TOP of the rock

	var window uint32

	for {
		wind := nextWind()

		canMove := true
		shifted := r.shiftLeft()
		if wind == '>' {
			shifted = r.shiftRight()
		}

		if wind == '<' && r.isFarLeft() {
			canMove = false
		} else if wind == '>' && r.isFarRight() {
			canMove = false
		} else {
			// Constrict window to the height of the rock
			testWindow := window << (8 * (4 - r.height()))

			canMove = uint32(shifted)&testWindow == 0
			//fmt.Printf("WINDOW: %032b, %032b => %v, (%d, %d)\n", shifted, window, canMove, y, r.height())
		}

		if canMove {
			r = shifted

			//fmt.Printf("Move %c\n", wind)
			//printPit(pit, r, y)
		} else {
			//fmt.Printf("Can't move %c, %032b, %032b\n", wind, r, shifted)
		}

		if y == 0 {
			// On the ground
			break
		}

		window = window << 8
		if y-r.height() >= 0 {
			window |= uint32(pit[y-r.height()])
		} else {
			window |= uint32(fullRow)
		}
		// Constrict window to the height of the rock
		testWindow := window << (8 * (4 - r.height()))

		canFall := uint32(r)&testWindow == 0
		//fmt.Printf("WINDOW: %032b, %032b => %v (%d, %d)\n", r, testWindow, canFall, y, r.height())

		if canFall {
			y -= 1
			//fmt.Printf("Fall\n")
			//printPit(pit, r, y)
		} else {
			break
		}
	}

	//fmt.Printf("END: %d\n", y)

	// Lock in the rock
	for i := 0; i < r.height(); i += 1 {
		if y-i < 0 {
			continue
		}

		row := r.getRow(i)
		pit[y-i] |= row
	}

	return y + 1
}
