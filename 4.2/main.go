package main

import (
	"fmt"
	"strconv"
	"strings"

	"main/util"
)

type assignment struct {
	min, max int
}

func (a assignment) containsMinPointOf(b assignment) bool {
	return a.min <= b.min && a.max >= b.min
}

func parseInput(line string) (assignment, assignment) {
	pair := strings.Split(line, ",")
	if len(pair) != 2 {
		panic(fmt.Sprintf("Bad input format %q", line))
	}

	return parseAssignment(pair[0]), parseAssignment(pair[1])
}

func parseAssignment(str string) assignment {
	pair := strings.Split(str, "-")
	if len(pair) != 2 {
		panic(fmt.Sprintf("Bad assignment format %q", str))
	}

	min, err := strconv.Atoi(pair[0])
	util.Check(err)

	max, err := strconv.Atoi(pair[1])
	util.Check(err)

	return assignment{min, max}
}

func main() {
	cnt := 0

	for _, line := range util.ReadInputLines("./input.txt") {
		a, b := parseInput(line)

		if a.containsMinPointOf(b) || b.containsMinPointOf(a) {
			cnt += 1
		}
	}

	fmt.Printf("%d\n", cnt)
}
