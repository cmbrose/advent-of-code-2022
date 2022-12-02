package main

import (
	"fmt"
	"strings"

	"main/util"
)

var (
	rock     = 1
	paper    = 2
	scissors = 3
)

func resultScore(opp, you int) int {
	diff := opp - you
	switch diff {
	case 0:
		return 3
	case 1, -2:
		return 0
	case -1, 2:
		return 6
	}

	panic("Unexpected score diff")
}

func inputToMove(input string) int {
	switch input {
	case "A", "X":
		return rock
	case "B", "Y":
		return paper
	case "C", "Z":
		return scissors
	}

	panic("Unexpected input")
}

func main() {
	score := 0

	for _, line := range util.ReadInputLines("./input.txt") {
		pair := strings.Split(line, " ")

		opp := inputToMove(pair[0])
		you := inputToMove(pair[1])

		score += resultScore(opp, you) + you
	}

	fmt.Printf("%d\n", score)
}
