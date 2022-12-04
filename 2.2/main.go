package main

import (
	"fmt"
	"strings"

	"main/util"
)

func inputToScore(opp, result string) int {
	oppVal := int(opp[0] - 'A')
	resVal := int(result[0] - 'X')

	// Opp | Res | Diff   | You
	// ----+-----+--------+----
	// 0 R | 0 L | 0      | 2 S
	// 0 R | 1 D | -1 (2) | 0 R
	// 0 R | 2 W | -2 (1) | 1 P
	// ----+-----+---------+----
	// 1 P | 0 L | 1      | 0 R
	// 1 P | 1 D | 0      | 1 P
	// 1 P | 2 W | -1 (2) | 2 S
	// ----+-----+--------+----
	// 2 S | 0 L | 2      | 1 P
	// 2 S | 1 D | 1      | 2 S
	// 2 S | 2 W | 0      | 0 R

	// You = (Diff + ResModifier) % 3

	// Diff = (Opp - Res + 3) % 3

	// ResModifier(lose) = +2
	// ResModifier(draw) => +1
	// ResModifier(win)  => +0
	// => ResModifier = 2 - Res

	// You = (((Opp - Res + 3) % 3) + (2 - Res)) % 3
	// You = (Opp - Res + 3 + 2 - Res) % 3
	// You = (Opp - 2*Res + 5) % 3

	you := (oppVal-2*resVal+5)%3 + 1 // +1 to re-align with the score of each move

	// fmt.Printf("%s %s => %d, %d\n", opp, result, you, resVal*3)
	return resVal*3 + you
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
