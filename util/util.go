package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
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

func ReadInputBlocks() [][]string {
	content, err := os.ReadFile("./input.txt")
	Check(err)

	blocks := strings.Split(string(content), "\n\n")

	return Map(blocks, func(block string) []string {
		return strings.Split(block, "\n")
	})
}

func Map[X, Y any](xArr []X, f func(X) Y) []Y {
	yArr := make([]Y, len(xArr))

	for i, x := range xArr {
		y := f(x)
		yArr[i] = y
	}

	return yArr
}

func AssertInt(str string) int {
	i, err := strconv.Atoi(str)
	Check(err)
	return i
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
func Intersect[T constraints.Ordered](a, b []T) []T {
	var res []T

	if len(a) == 0 || len(b) == 0 {
		return res
	}

	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		a := a[i]
		b := b[j]

		if a == b {
			res = append(res, a)
			i++
			j++
		} else if a > b { // a > b
			j++
		} else { // a < b
			i++
		}
	}

	return res
}

func IntersectAll[T constraints.Ordered](a ...[]T) []T {
	if len(a) == 0 {
		return nil
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
func Except[T constraints.Ordered](a, b []T) []T {
	if len(a) == 0 || len(b) == 0 {
		return a
	}

	var res []T

	i := 0
	j := 0
	for i < len(a) && j < len(b) {
		a := a[i]
		b := b[j]

		if a == b {
			i++
			j++
		} else if a > b { // a > b
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
func ExceptAll[T constraints.Ordered](a []T, b ...[]T) []T {
	if len(a) == 0 || len(b) == 0 {
		return a
	}

	cur := a

	for _, b := range b {
		cur = Except(cur, b)
	}

	return cur
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

func IsUppercase(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func IsLowercase(b byte) bool {
	return b >= 'a' && b <= 'z'
}

func IsLetter(b byte) bool {
	return IsUppercase(b) || IsLowercase(b)
}

func IsNumber(b byte) bool {
	return b >= '0' && b <= '9'
}
