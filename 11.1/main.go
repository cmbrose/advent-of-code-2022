package main

import (
	"fmt"
	"main/util"
	"strings"
)

type monkey struct {
	items     []int
	update    func(int) int
	getTarget func(int) int
	observed  int
}

func parseMonkey(input []string) *monkey {
	if len(input) != 5 {
		panic("Wrong number of monkey config lines")
	}

	itemsStr := strings.Split(input[0], ": ")[1]
	items := util.Map(strings.Split(itemsStr, ", "), util.AssertInt)

	operationStr := strings.Split(input[1], "= ")[1]
	operationParts := strings.Split(operationStr, " ")
	update := func(worry int) int {
		pick := func(label string) int {
			if label == "old" {
				return worry
			}
			return util.AssertInt(label)
		}

		x := pick(operationParts[0])
		y := pick(operationParts[2])

		switch operationParts[1] {
		case "+":
			return x + y
		case "-":
			return x - y
		case "*":
			return x * y
		case "/":
			return x / y
		default:
			panic(fmt.Sprintf("Unexpected operand %q", operationParts[1]))
		}
	}

	divisorStr := strings.Split(input[2], "by ")[1]
	divisor := util.AssertInt(divisorStr)

	ifTrueStr := strings.Split(input[3], "monkey ")[1]
	ifTrue := util.AssertInt(ifTrueStr)

	ifFalseStr := strings.Split(input[4], "monkey ")[1]
	ifFalse := util.AssertInt(ifFalseStr)

	getTarget := func(worry int) int {
		if worry%divisor == 0 {
			return ifTrue
		}

		return ifFalse
	}

	return &monkey{
		items,
		update,
		getTarget,
		0,
	}
}

func printMonkeys(monkeys map[int]*monkey, monkeyOrder []int) {
	for _, id := range monkeyOrder {
		monkey := monkeys[id]
		fmt.Printf("Monkey %d: %v\n", id, monkey.items)
	}
}

func main() {
	monkeyOrder := []int{}
	monkeys := make(map[int]*monkey)

	for _, block := range util.ReadInputBlocks() {
		idStr := strings.TrimSuffix(strings.Split(block[0], " ")[1], ":")
		id := util.AssertInt(idStr)

		monkey := parseMonkey(block[1:])
		monkeys[id] = monkey

		monkeyOrder = append(monkeyOrder, id)
	}

	printMonkeys(monkeys, monkeyOrder)

	reps := 20

	for i := 0; i < reps; i += 1 {
		for _, id := range monkeyOrder {
			monkey := monkeys[id]
			monkey.observed += len(monkey.items)
			for _, worry := range monkey.items {
				newWorry := monkey.update(worry) / 3

				targetId := monkey.getTarget(newWorry)
				targetMonkey := monkeys[targetId]

				targetMonkey.items = append(targetMonkey.items, newWorry)
			}
			monkey.items = []int{}
		}

		fmt.Printf("Round %d:\n", i)
		printMonkeys(monkeys, monkeyOrder)
		fmt.Println()
	}

	var max1, max2 int
	for _, monkey := range monkeys {
		if monkey.observed > max1 {
			max2 = max1
			max1 = monkey.observed
		} else if monkey.observed > max2 {
			max2 = monkey.observed
		}
	}

	fmt.Printf("%d = %d * %d\n", max1*max2, max1, max2)
}
