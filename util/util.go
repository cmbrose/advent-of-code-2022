package util

import (
	"fmt"
	"os"
	"strings"
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadInputLines(path string) []string {
	content, err := os.ReadFile("./input.txt")
	Check(err)

	return strings.Split(string(content), "\n")
}

func ParseBitString(str string) int {
	val := 0
	for _, c := range str {
		val <<= 1
		if c == '1' {
			val++
		}
	}

	return val
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func MinInt(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func MaxInt(x, y int) int {
	if x < y {
		return y
	}

	return x
}

// Must be pre-sorted!
func Intersect(a, b []interface{}) []interface{} {
	if len(a) == 0 || len(b) == 0 {
		return []interface{}{}
	}

	res := []interface{}{}

	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		a := a[i]
		b := b[j]

		comp := Compare(a, b)
		if comp == 0 {
			res = append(res, a)
			i++
			j++
		} else if comp > 0 { // a > b
			j++
		} else { // a < b
			i++
		}
	}

	return res
}

func IntersectAll(a ...[]interface{}) []interface{} {
	if len(a) == 0 {
		return []interface{}{}
	}

	cur := a[0]

	if len(a) == 1 {
		return cur
	}

	for _, b := range a[1:] {
		cur = Intersect(cur, b)
	}

	return cur
}

// Must be pre-sorted!
func Except(a, b []interface{}) []interface{} {
	if len(a) == 0 || len(b) == 0 {
		return a
	}

	res := []interface{}{}

	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		a := a[i]
		b := b[j]

		comp := Compare(a, b)
		if comp == 0 {
			i++
			j++
		} else if comp > 0 { // a > b
			j++
		} else { // a < b
			res = append(res, a)
			i++
		}
	}

	for i < len(a) {
		res = append(res, a[i])
		i++
	}

	return res
}

// Must be pre-sorted!
func ExceptAll(a []interface{}, b ...[]interface{}) []interface{} {
	if len(a) == 0 || len(b) == 0 {
		return a
	}

	cur := a

	for _, b := range b {
		cur = Except(cur, b)
	}

	return cur
}

func IntSliceToInterfaceSlice(a []int) []interface{} {
	res := make([]interface{}, len(a))

	for i, a := range a {
		res[i] = a
	}

	return res
}

func RuneSliceToInterfaceSlice(a []rune) []interface{} {
	res := make([]interface{}, len(a))

	for i, a := range a {
		res[i] = a
	}

	return res
}

func Compare(a, b interface{}) int {
	switch a := a.(type) {
	case int:
		b, ok := b.(int)
		if !ok {
			panic("type mismatch")
		}
		return a - b

	case rune:
		b, ok := b.(rune)
		if !ok {
			panic("type mismatch")
		}
		return int(a) - int(b)

	default:
		panic("Unhandled type")
	}
}

func ParseIntGrid() [][]int {
	grid := [][]int{}

	for _, line := range ReadInputLines("./input.txt") {
		row := []int{}
		for _, cell := range line {
			row = append(row, int(cell-'0'))
		}

		grid = append(grid, row)
	}

	return grid
}

func PrintIntGrid(grid [][]int) {
	rows := make([]string, len(grid))

	for i, row := range grid {
		rows[i] = ""
		for _, cell := range row {
			rows[i] += fmt.Sprintf("%d", cell)
		}
	}

	fmt.Println(strings.Join(rows, "\n"))
}

func HexToBinary(char byte) string {
	switch char {
	case '0':
		return "0000"
	case '1':
		return "0001"
	case '2':
		return "0010"
	case '3':
		return "0011"
	case '4':
		return "0100"
	case '5':
		return "0101"
	case '6':
		return "0110"
	case '7':
		return "0111"
	case '8':
		return "1000"
	case '9':
		return "1001"
	case 'A':
		return "1010"
	case 'B':
		return "1011"
	case 'C':
		return "1100"
	case 'D':
		return "1101"
	case 'E':
		return "1110"
	case 'F':
		return "1111"
	}

	panic(fmt.Sprintf("Unknown hex char: %c", char))
}
