package main

import (
	"reflect"
	"testing"
)

// Time complexity: O(4^(n/sqrt(n))), the catalan sequence.
// Space complexity: O(n), for recursion call stack.
// https://claude.ai/chat/a63edbcb-06cd-4758-ae20-b57e4bf4da7b
func generateParenthesis22(n int) []string {
	// Generate Parentheses
	// Given an integer n.
	// n is the number of paren pairs.
	// Create all possible combinations of parens.
	// Return combinations.
	// Use backtracking and recursion.
	// Use a closure for recursion and state management.

	// Initialize variables.
	var res []string

	// Create closuse for recursion and state management.
	var bck func(str string, opn, cls int)
	bck = func(str string, opn, cls int) {
		// Check base case.
		// Check whether combination is at final length 2*n.
		if len(str) == 2*n {
			// Found combination.
			res = append(res, str)
			return
		}

		// Add open paren when under n.
		if opn < n {
			bck(str+"(", opn+1, cls)
		}

		// Add close paren when under open count.
		if cls < opn {
			bck(str+")", opn, cls+1)
		}
	}

	// Backtrack with empty string.
	bck("", 0, 0)

	return res
}

// TestGenerateParenthesis22 runs test cases for generateParenthesis
func TestGenerateParenthesis22(t *testing.T) {
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
				res := generateParenthesis22(8)
				if len(res) != 1430 { // Catalan number C(8)
					t.Errorf("Expected 1430 combinations for n=8")
				}
				return res
			}(),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := generateParenthesis22(tc.n)
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
