package main

import (
	"errors"
	"fmt"
	"io"
	"main/util"
	"strings"
)

type nestedListItem struct {
	isNested bool

	value  int
	nested nestedList
}

func (nli nestedListItem) String() string {
	if nli.isNested {
		return nli.nested.String()
	}

	return fmt.Sprintf("%d", nli.value)
}

type nestedList []nestedListItem

func (nl nestedList) String() string {
	var bldr strings.Builder

	bldr.WriteRune('[')
	for i, nli := range nl {
		if i > 0 {
			bldr.WriteString(",")
		}

		bldr.WriteString(nli.String())
	}
	bldr.WriteRune(']')

	return bldr.String()
}

func parseList(list string) (nestedList, error) {
	reader := strings.NewReader(list)

	r, _, err := reader.ReadRune()
	if err != nil {
		return nil, fmt.Errorf("reading char: %w", err)
	}

	if r != '[' {
		return nil, fmt.Errorf("invalid starting rune, got %c", r)
	}

	parsed, err := parseListRec(reader)
	if err != nil {
		return nil, err
	}

	if reader.Len() != 0 {
		return nil, errors.New("parseListRec returned with unread chars")
	}

	return parsed, nil
}

func parseListRec(reader *strings.Reader) (nestedList, error) {
	list := nestedList{}
	intBldr := 0
	hasInt := false

	for r, _, err := reader.ReadRune(); !errors.Is(err, io.EOF); r, _, err = reader.ReadRune() {
		// fmt.Printf("parseListRec state: %v, %d, %c\n", list, intBldr, r)

		if err != nil {
			return nil, fmt.Errorf("reading char: %w", err)
		}

		switch r {
		case ']', ',':
			// We should only hit these while reading an int or empty list
			if hasInt {
				item := nestedListItem{isNested: false, value: intBldr}
				list = append(list, item)
				intBldr = 0
				hasInt = false
			}

			if r == ']' {
				return list, nil
			}

		case ' ':
			continue

		case '[':
			nested, err := parseListRec(reader)
			if err != nil {
				return nil, err
			}

			item := nestedListItem{isNested: true, nested: nested}
			list = append(list, item)

			next, _, err := reader.ReadRune()
			if err != nil {
				return nil, fmt.Errorf("reading char: %w", err)
			}

			if next == ',' {
				// Nothing to do, this item is done
			} else if next == ']' {
				// The whole list is done
				return list, nil
			} else {
				return nil, fmt.Errorf("expected to read comma or close-brace, but was: %c", next)
			}

		default: // Assume a digit
			hasInt = true
			digit := int(r - '0')

			if digit < 0 || digit > 9 {
				return nil, fmt.Errorf("expected to digit, but was: %c", r)
			}

			intBldr = intBldr*10 + digit
		}
	}

	return nil, errors.New("reader reached end before finding end of list")
}

type result int

const (
	valid result = iota
	invalid
	inconclusive
)

func isValid(left, right nestedList) result {
	var lc, rc int

	for lc < len(left) && rc < len(right) {
		l, r := left[lc], right[rc]

		if !l.isNested && !r.isNested {
			if l.value < r.value {
				return valid
			}
			if l.value > r.value {
				return invalid
			}
		} else {
			ln := l.nested
			if !l.isNested {
				ln = nestedList{nestedListItem{isNested: false, value: l.value}}
			}

			rn := r.nested
			if !r.isNested {
				rn = nestedList{nestedListItem{isNested: false, value: r.value}}
			}

			res := isValid(ln, rn)
			if res != inconclusive {
				return res
			}
		}

		lc, rc = lc+1, rc+1
	}

	if lc == len(left) && rc == len(right) {
		return inconclusive
	}
	if lc == len(left) {
		return valid
	}
	return invalid
}

func main() {
	sum := 0

	for i, block := range util.ReadInputBlocks() {
		left, err := parseList(block[0])
		util.Check(err)

		right, err := parseList(block[1])
		util.Check(err)

		res := isValid(left, right)

		// fmt.Println(left.String())
		// fmt.Println(right.String())
		// fmt.Println(res)
		// fmt.Println()

		if res == inconclusive {
			panic("Inconslusive result")
		}

		if res == valid {
			sum += i + 1
		}
	}

	fmt.Println(sum)
}
