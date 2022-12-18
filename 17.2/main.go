package main

import (
	"encoding/hex"
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

	top := 0
	pit := []byte{}

	type state struct {
		pitSlice string
		rc       int
		wc       int
	}

	type snapshot struct {
		top int
		y   int
	}

	cache := make(map[state]snapshot)

	var loopLen int
	var loopHeight int

	target := 1_000_000_000_000

	y := 0
	for y = 0; y < target; y += 1 {
		r := nextRock()
		requiredHeight := top + 3 + r.height()

		for height := len(pit); height < requiredHeight; height += 1 {
			pit = append(pit, 0)
		}

		testTop := fall(r, requiredHeight-1, pit, nextWind)
		top = util.MaxInt(top, testTop)

		s := state{
			pitSlice: getPitSlice(pit, top),
			rc:       rockCursor,
			wc:       windCursor,
		}

		if seen, ok := cache[s]; ok {
			loopLen = y - seen.y
			loopHeight = top - seen.top
			y += 1 // This iteration finished but we'll skip the increment
			break
		}

		snap := snapshot{
			top: top,
			y:   y,
		}

		cache[s] = snap
	}

	//fmt.Printf("Each loop of length %d adds %d\n", loopLen, loopHeight)

	loops := (target - y) / loopLen
	y += loops * loopLen
	incrTop := loops * loopHeight

	//fmt.Printf("Simulated %d loops => %d, %d\n", loops, y, incrTop)

	for ; y < target; y += 1 {
		r := nextRock()
		requiredHeight := top + 3 + r.height()

		for height := len(pit); height < requiredHeight; height += 1 {
			pit = append(pit, 0)
		}

		testTop := fall(r, requiredHeight-1, pit, nextWind)
		top = util.MaxInt(top, testTop)
	}

	//printPit(pit, 0, 0)
	fmt.Println(top + incrTop)
}

func getPitSlice(pit []byte, top int) string {
	size := util.MinInt(100, top)

	slice := pit[top-size : top]
	return hex.EncodeToString(slice)
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
