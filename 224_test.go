package main

import "testing"

// Time complexity: O(n), n is the length of the s string. We traverse the string once.
// Space complexity: O(d), d is the stack depth based on the number of nested parentheses.
// https://claude.ai/chat/a62bad61-3c59-4b0f-955d-b506101be96b
func calculate(s string) int {
	// Basic Calculator
	// Given a string s.
	// String s is an arithmetic expression.
	// Evaluate the whole expression.
	// Return an integer result of the whole expression.
	// Use a stack to store intermediate sign and result.
	// Build up an integer one character at a time.

	// Initialize variables.
	var stk []int
	sgn := 1
	num := 0

	// Initialize stack with base number 0 for first number.
	// Stack always contains a result.
	stk = append(stk, 0)

	// Iterate through each character.
	for idx := 0; idx < len(s); idx++ {
		switch s[idx] {
		// Build up number.
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num = num*10 + int(s[idx]-'0')

		// Apply arithmetic operator.
		// Reset sign and number.
		case '+', '-':
			// Apply arithmetic operator.
			stk[len(stk)-1] += sgn * num

			// Reset sign and number.
			sgn = 1
			if s[idx] == '-' {
				sgn = -1
			}
			num = 0

		// Prepare for inner expression.
		// Store sign and base number 0.
		// Reset sign.
		case '(':
			// Store sign and base number 0.
			stk = append(stk, sgn, 0)

			// Reset sign.
			sgn = 1

		// Finalize inner expression number.
		// Apply inner expression to outer expression.
		// Reset current sign and number.
		case ')':
			// Finalize inner expression number.
			stk[len(stk)-1] += sgn * num

			// Get inner sign and number, and pop.
			inrNum := stk[len(stk)-1]
			inrSgn := stk[len(stk)-2]
			stk = stk[:len(stk)-2]

			// Apply inner expression to outer expression.
			stk[len(stk)-1] += inrSgn * inrNum

			// Reset the current sign and number.
			sgn = 1
			num = 0
		}
	}

	// Apply current number and return.
	return stk[0] + sgn*num
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		name string
		expr string
		want int
	}{
		{"simple addition", "1 + 1", 2},
		{"addition and subtraction", " 2-1 + 2 ", 3},
		{"nested parentheses", "(1+(4+5+2)-3)+(6+8)", 23},
		{"single digit", "2", 2},
		{"multiple digits", "123", 123},
		{"nested subtraction", "2-(5-6)", 3},
		{"complex nested", "1-(2+3-(4+5))", 5},
		{"multiple spaces", "   3  +  2   ", 5},
		{"deep nesting", "(1+(2+(3+(4))))", 10},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := calculate(tc.expr)
			if got != tc.want {
				t.Errorf("calculate(%q) = %d; want %d", tc.expr, got, tc.want)
			}
		})
	}
}
