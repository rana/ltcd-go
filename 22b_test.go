package main

import (
	"reflect"
	"testing"
)

func generateParenthesis22b(n int) []string {
	// Generate Parentheses
	// Given an integer n.
	// n is a number of paren pairs.
	// Create all combinations of paren pairs.
	// Return the combinations.
	// Use backtracking and recursion.
	// Use a closure for recursion and state management.
	// 2*n characters in a combination.

	// Initialize variables.
	var res []string

	// Create a closure.
	var bck func(str string, opn, cls int)
	bck = func(str string, opn, cls int) {
		// Check base case.
		// Check whether whole combination formed.
		if len(str) == 2*n {
			res = append(res, str)
			return
		}

		// Recurse open bracket until n.
		if opn < n {
			bck(str+"(", opn+1, cls)
		}

		// Recurse close bracket until opn.
		if cls < opn {
			bck(str+")", opn, cls+1)
		}
	}

	bck("", 0, 0)

	return res
}

// TestGenerateParenthesis22b runs test cases for generateParenthesis
func TestGenerateParenthesis22b(t *testing.T) {
	cases := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "Example 1 - Three pairs",
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "Example 2 - One pair",
			n:    1,
			want: []string{"()"},
		},
		{
			name: "Two pairs",
			n:    2,
			want: []string{"(())", "()()"},
		},
		{
			name: "Edge case - Maximum input",
			n:    8,
			// Length check only due to large output size
			want: func() []string {
				res := generateParenthesis22b(8)
				if len(res) != 1430 { // Catalan number C(8)
					t.Errorf("Expected 1430 combinations for n=8")
				}
				return res
			}(),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := generateParenthesis22b(tc.n)
			if tc.n < 8 && !reflect.DeepEqual(got, tc.want) {
				t.Errorf("generateParenthesis(%v) = %v, want %v", tc.n, got, tc.want)
			}
			// For n=8, just verify length
			if tc.n == 8 && len(got) != 1430 {
				t.Errorf("generateParenthesis(8) returned %d combinations, want 1430", len(got))
			}
		})
	}
}
