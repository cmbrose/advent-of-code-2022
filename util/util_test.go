package util

import (
	"testing"
)

func TestIntersect(t *testing.T) {
	tests := []struct {
		a []interface{}
		b []interface{}
		e []interface{}
	}{
		{
			[]interface{}{1, 2, 3},
			[]interface{}{1, 2, 3},
			[]interface{}{1, 2, 3},
		},
		{
			[]interface{}{1, 2, 3},
			[]interface{}{3, 4, 5},
			[]interface{}{3},
		},
		{
			[]interface{}{1, 2, 3},
			[]interface{}{4, 5, 6},
			[]interface{}{},
		},
		{
			[]interface{}{},
			[]interface{}{4, 5, 6},
			[]interface{}{},
		},
		{
			[]interface{}{1, 2, 3},
			[]interface{}{},
			[]interface{}{},
		},
		{
			[]interface{}{1},
			[]interface{}{1},
			[]interface{}{1},
		},
		{
			[]interface{}{'a', 'b', 'c'},
			[]interface{}{'b', 'c', 'e'},
			[]interface{}{'b', 'c'},
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
				// TODO - report char values better
				t.Fatalf("Incorrect value at index %d. Expected=%v, got=%v", i, exp, act)
			}
		}
	}
}

func TestExcept(t *testing.T) {
	tests := []struct {
		a []interface{}
		b []interface{}
		e []interface{}
	}{
		{
			[]interface{}{1, 2, 3},
			[]interface{}{1, 2, 3},
			[]interface{}{},
		},
		{
			[]interface{}{1, 2, 3},
			[]interface{}{3, 4, 5},
			[]interface{}{1, 2},
		},
		{
			[]interface{}{1, 2, 3},
			[]interface{}{4, 5, 6},
			[]interface{}{1, 2, 3},
		},
		{
			[]interface{}{},
			[]interface{}{4, 5, 6},
			[]interface{}{},
		},
		{
			[]interface{}{1, 2, 3},
			[]interface{}{},
			[]interface{}{1, 2, 3},
		},
		{
			[]interface{}{1},
			[]interface{}{1},
			[]interface{}{},
		},
		{
			[]interface{}{1},
			[]interface{}{2},
			[]interface{}{1},
		},
		{
			[]interface{}{'a', 'b', 'c'},
			[]interface{}{'b', 'c', 'e'},
			[]interface{}{'a'},
		},
		{
			[]interface{}{'a', 'b', 'd'},
			[]interface{}{'a', 'b'},
			[]interface{}{'d'},
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
				// TODO - report char values better
				t.Fatalf("Incorrect value at index %d. Expected=%v, got=%v", i, exp, act)
			}
		}
	}
}
