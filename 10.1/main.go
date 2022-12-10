package main

import (
	"fmt"
	"strconv"
	"strings"

	"main/util"
)

func main() {
	var (
		val          = 1
		cycle        = 1
		nextKeyCycle = 20
		keyCycleIncr = 40
		score        = 0
	)

	onKeyCycle := func() {
		fmt.Printf("Key cycle %d with val %d => %d\n", nextKeyCycle, val, nextKeyCycle*val)

		score += nextKeyCycle * val
		nextKeyCycle += keyCycleIncr
	}

	for _, line := range util.ReadInputLines("./input.txt") {
		var (
			addVal   = 0
			addCycle = 1
			err      error
		)

		//fmt.Printf("Cycle %d: val = %d, input = %s\n", cycle, val, line)

		pair := strings.Split(line, " ")
		cmd := pair[0]

		if cmd == "addx" {
			addVal, err = strconv.Atoi(pair[1])
			util.Check(err)
			addCycle = 2
		}

		nextCycle := cycle + addCycle
		if nextKeyCycle < nextCycle {
			// Must be an addx that overshoots the cycle, so use current value
			onKeyCycle()
		}

		val += addVal

		if nextKeyCycle == nextCycle {
			// Could be a noop, or an addx that lands on the key cycle
			onKeyCycle()
		}

		cycle = nextCycle
	}

	fmt.Printf("%d\n", score)
}
