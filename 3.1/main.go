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

	for _, line := range util.ReadInputLines("./input.txt") {
		left := []rune(line[:len(line)/2])
		right := []rune(line[len(line)/2:])

		sort.Slice(left, func(i, j int) bool { return left[i] < left[j] })
		sort.Slice(right, func(i, j int) bool { return right[i] < right[j] })

		diff := util.Intersect(
			util.RuneSliceToInterfaceSlice(left),
			util.RuneSliceToInterfaceSlice(right),
		)

		item := diff[0].(rune)

		res += itemToPriority(item)
	}

	fmt.Printf("%d\n", res)
}
