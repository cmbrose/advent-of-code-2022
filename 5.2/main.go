package main

import (
	"fmt"
	"strconv"
	"strings"

	"main/util"
)

func printStacks(stacks [][]byte) {
	maxHeight := 0
	for _, stack := range stacks {
		maxHeight = util.MaxInt(maxHeight, len(stack))
	}

	var bldr strings.Builder
	for i := 0; i < maxHeight; i += 1 {
		for _, stack := range stacks {
			if len(stack) < maxHeight-i {
				bldr.WriteString("    ")
			} else {
				bldr.WriteString(fmt.Sprintf("[%c] ", stack[maxHeight-i-1]))
			}
		}
		bldr.WriteByte('\n')
	}

	fmt.Println(bldr.String())
}

func main() {
	var stacks [][]byte

	i := 0
	lines := util.ReadInputLines("./input.txt")

	for line := lines[i]; !util.IsNumber(line[1]); line = lines[i] {
		if len(line)/4 >= len(stacks) {
			stacks = append(stacks, make([][]byte, len(line)/4-len(stacks)+1)...)
		}

		for j := 0; j < len(line); j += 4 {
			idx := j / 4
			crate := line[j+1]

			if crate == ' ' {
				continue
			}

			stacks[idx] = append([]byte{crate}, stacks[idx]...)
		}

		i += 1
	}

	printStacks(stacks)

	// Skip the stack index line and blank
	i += 2

	for _, line := range lines[i:] {
		parts := strings.Split(line, " ")
		if len(parts) != 6 {
			panic(fmt.Sprintf("Incorrectlu formatted line: %q", line))
		}

		amount, err := strconv.Atoi(parts[1])
		util.Check(err)

		from, err := strconv.Atoi(parts[3])
		util.Check(err)
		from -= 1 // 0 based

		to, err := strconv.Atoi(parts[5])
		util.Check(err)
		to -= 1 // 0 based

		height := len(stacks[from])
		stacks[to] = append(stacks[to], stacks[from][height-amount:]...)
		stacks[from] = stacks[from][:height-amount]

		fmt.Println(line)
		printStacks(stacks)
	}

	for _, stack := range stacks {
		fmt.Printf("%c", stack[len(stack)-1])
	}

	fmt.Println()
}
