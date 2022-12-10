package main

import (
	"fmt"
	"strconv"
	"strings"

	"main/util"
)

func main() {
	var (
		val        = 1
		cursor     = 0
		lineLength = 40
	)

	incrCursor := func() {
		if util.AbsInt(val-cursor) < 2 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}

		cursor += 1
		cursor %= lineLength

		if cursor == 0 {
			fmt.Println()
		}
	}

	for _, line := range util.ReadInputLines("./input.txt") {
		pair := strings.Split(line, " ")
		cmd := pair[0]

		incrCursor()

		if cmd == "addx" {
			incrCursor()

			addVal, err := strconv.Atoi(pair[1])
			util.Check(err)

			val += addVal
		}
	}
}
