package main

import (
	"fmt"
	"sort"

	"main/util"
)

func itemToPriority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return int(item-'a') + 1
	} else if item >= 'A' && item <= 'Z' {
		return int(item-'A') + 27
	}

	panic(fmt.Sprintf("Unknown item %c", item))
}

func main() {
	res := 0

	idx := 0
	packs := make([][]interface{}, 3)

	for _, line := range util.ReadInputLines("./input.txt") {
		pack := []rune(line)

		sort.Slice(pack, func(i, j int) bool { return pack[i] < pack[j] })

		packs[idx] = util.RuneSliceToInterfaceSlice(pack)

		idx += 1

		if idx == 3 {
			diff := util.IntersectAll(packs...)
			item := diff[0].(rune)
			res += itemToPriority(item)

			idx = 0
		}
	}

	fmt.Printf("%d\n", res)
}
