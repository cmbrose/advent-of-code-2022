package main

import (
	"fmt"
	"main/util"
	"strings"
)

type moduloWorry map[int]int

type monkey struct {
	items []moduloWorry

	update func(int) int

	divisor int

	zeroTarget  int
	otherTarget int

	observed int
}

func parseMonkey(input []string) (*monkey, []int) {
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

	return &monkey{
		nil,
		update,
		divisor,
		ifTrue,
		ifFalse,
		0,
	}, items
}

func printMonkeys(monkeys map[int]*monkey, monkeyOrder []int) {
	for _, id := range monkeyOrder {
		monkey := monkeys[id]
		fmt.Printf("Monkey %d: %v observations\n", id, monkey.observed)
	}
}

func main() {
	monkeyOrder := []int{}
	monkeys := make(map[int]*monkey)

	type monkeyWorryPair struct {
		m *monkey
		w []int
	}

	pairs := []monkeyWorryPair{}

	for _, block := range util.ReadInputBlocks("./input.txt") {
		idStr := strings.TrimSuffix(strings.Split(block[0], " ")[1], ":")
		id := util.AssertInt(idStr)

		monkey, items := parseMonkey(block[1:])
		monkeys[id] = monkey

		pairs = append(pairs, monkeyWorryPair{monkey, items})

		monkeyOrder = append(monkeyOrder, id)
	}

	for _, p := range pairs {
		for _, worry := range p.w {
			modWorry := make(moduloWorry)
			for id, monkey := range monkeys {
				modWorry[id] = worry % monkey.divisor
				//fmt.Printf("[%d] %d => %d\n", id, worry, modWorry[id])
			}
			//fmt.Println()

			p.m.items = append(p.m.items, modWorry)
		}
	}

	reps := 10_000

	for i := 0; i < reps; i += 1 {
		for _, id := range monkeyOrder {
			monkey := monkeys[id]
			monkey.observed += len(monkey.items)

			for _, modWorries := range monkey.items {
				newWorries := make(moduloWorry)

				for _, updateId := range monkeyOrder {
					worry := modWorries[updateId]
					updateMonkey := monkeys[updateId]

					newWorry := monkey.update(worry)
					newWorries[updateId] = newWorry % updateMonkey.divisor

					//fmt.Printf("[%d]: [%d] %d => %d => %d\n", id, updateId, worry, newWorry, newWorries[updateId])
				}

				targetId := monkey.otherTarget
				if newWorries[id] == 0 {
					targetId = monkey.zeroTarget
				}
				//fmt.Printf("[%d] throws %d to [%d]\n", id, newWorries[id], targetId)

				targetMonkey := monkeys[targetId]
				targetMonkey.items = append(targetMonkey.items, newWorries)

				//fmt.Println()
			}
			monkey.items = nil
		}

		// fmt.Printf("Round %d:\n", i)
		// printMonkeys(monkeys, monkeyOrder)
		// fmt.Println()
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
