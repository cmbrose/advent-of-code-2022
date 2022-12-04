package util

import (
	"testing"
)

func TestIntersectInts(t *testing.T) {
	tests := []struct {
		a []int
		b []int
		e []int
	}{
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			[]int{1, 2, 3},
		},
		{
			[]int{1, 2, 3},
			[]int{3, 4, 5},
			[]int{3},
		},
		{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{},
		},
		{
			[]int{},
			[]int{4, 5, 6},
			[]int{},
		},
		{
			[]int{1, 2, 3},
			[]int{},
			[]int{},
		},
		{
			[]int{1},
			[]int{1},
			[]int{1},
		},
	}

	for _, test := range tests {
		actual := Intersect(test.a, test.b)

		if len(actual) != len(test.e) {
			t.Fatalf("Incorrect output length. Expected=%d, got=%d", len(test.e), len(actual))
		}

		for i, exp := range test.e {
			act := actual[i]
			if act != exp {
				t.Fatalf("Incorrect value at index %d. Expected=%v, got=%v", i, exp, act)
			}
		}
	}
}

func TestIntersectRunes(t *testing.T) {
	tests := []struct {
		a []rune
		b []rune
		e []rune
	}{
		{
			[]rune{'a', 'b', 'c'},
			[]rune{'b', 'c', 'e'},
			[]rune{'b', 'c'},
		},
	}

	for _, test := range tests {
		actual := Intersect(test.a, test.b)

		if len(actual) != len(test.e) {
			t.Fatalf("Incorrect output length. Expected=%d, got=%d", len(test.e), len(actual))
		}

		for i, exp := range test.e {
			act := actual[i]
			if act != exp {
				t.Fatalf("Incorrect value at index %d. Expected=%v, got=%v", i, exp, act)
			}
		}
	}
}

func TestExceptInts(t *testing.T) {
	tests := []struct {
		a []int
		b []int
		e []int
	}{
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			[]int{},
		},
		{
			[]int{1, 2, 3},
			[]int{3, 4, 5},
			[]int{1, 2},
		},
		{
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			[]int{1, 2, 3},
		},
		{
			[]int{},
			[]int{4, 5, 6},
			[]int{},
		},
		{
			[]int{1, 2, 3},
			[]int{},
			[]int{1, 2, 3},
		},
		{
			[]int{1},
			[]int{1},
			[]int{},
		},
		{
			[]int{1},
			[]int{2},
			[]int{1},
		},
	}

	for _, test := range tests {
		actual := Except(test.a, test.b)

		if len(actual) != len(test.e) {
			t.Fatalf("Incorrect output length. Expected=%v, got=%v", test.e, actual)
		}

		for i, exp := range test.e {
			act := actual[i]
			if act != exp {
				t.Fatalf("Incorrect value at index %d. Expected=%v, got=%v", i, exp, act)
			}
		}
	}
}

func TestExceptRunes(t *testing.T) {
	tests := []struct {
		a []rune
		b []rune
		e []rune
	}{
		{
			[]rune{'a', 'b', 'c'},
			[]rune{'b', 'c', 'e'},
			[]rune{'a'},
		},
		{
			[]rune{'a', 'b', 'd'},
			[]rune{'a', 'b'},
			[]rune{'d'},
		},
	}

	for _, test := range tests {
		actual := Except(test.a, test.b)

		if len(actual) != len(test.e) {
			t.Fatalf("Incorrect output length. Expected=%v, got=%v", test.e, actual)
		}

		for i, exp := range test.e {
			act := actual[i]
			if act != exp {
				t.Fatalf("Incorrect value at index %d. Expected=%v, got=%v", i, exp, act)
			}
		}
	}
}
