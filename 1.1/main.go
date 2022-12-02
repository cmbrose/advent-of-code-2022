package main

import (
	"fmt"
	"strconv"
	"strings"

	"main/util"
)

func main() {
	max := 0
	curr := 0

	for _, line := range util.ReadInputLines("./input.txt") {
		if line == "" {
			max = util.MaxInt(max, curr)
			curr = 0
			continue
		}

		value, err := strconv.Atoi(strings.TrimSpace(line))
		util.Check(err)

		curr += value
	}

	fmt.Printf("%d", max)
}
