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

// Opponent move -> Desired result -> Your Move
var resultMatrix [][]int = [][]int{
	// rock
	{
		scissors, // lose
		rock,     // draw
		paper,    // win
	},
	// paper
	{
		rock,     // lose
		paper,    // draw
		scissors, // win
	},
	// scissors
	{
		paper,    // lose
		scissors, // draw
		rock,     // win
	},
}

func inputToScore(opp, result string) int {
	oppIdx := int(opp[0] - 'A')
	resIdx := int(result[0] - 'X')

	you := resultMatrix[oppIdx][resIdx]

	//fmt.Printf("%s %s => %d, %d\n", opp, result, you, resIdx*3)
	return resIdx*3 + you
}

func main() {
	score := 0

	for _, line := range util.ReadInputLines("./input.txt") {
		pair := strings.Split(line, " ")

		opp := pair[0]
		res := pair[1]

		score += inputToScore(opp, res)
	}

	fmt.Printf("%d\n", score)
}
