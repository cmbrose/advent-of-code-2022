package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"main/util"
)

func main() {
	var top3 []int
	curr := 0

	for _, line := range util.ReadInputLines("./input.txt") {
		if line == "" {
			top3 = append(top3, curr)

			if len(top3) > 3 {
				sort.Ints(top3)
				top3 = top3[1:]
			}

			curr = 0
			continue
		}

		value, err := strconv.Atoi(strings.TrimSpace(line))
		util.Check(err)

		curr += value
	}

	fmt.Printf("%d\n", top3[0]+top3[1]+top3[2])
}
