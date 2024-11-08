package main

import (
	"strconv"
	"testing"
)

// Time complexity: O(n), n is the length of the tkns array. We traverse teh array once.
// Space complexity: O(n), up to n strings stored on a stack.
// https://claude.ai/chat/b9a0b6fa-fedf-4218-aabc-2a34e92afcec
func evalRPN(tkns []string) int {
	// Evaluate Reverse Polish Notation
	// Given a string array.
	// Each string is part of an arithmetic expression.
	// Evaluate the whole expression.
	// Return the result of the whole expression.
	// Conditions:
	//	* Operators are + - * /
	//	* Operand is an integer or an expression.
	//	* Only valid expressions provided.

	// Initialize a stack to hold intermediate results.
	var stk []int

	// Iterate through each token.
	for _, tkn := range tkns {
		switch tkn {
		case "+", "-", "*", "/":
			// Pop last two integers from stack in reverse order.
			rht := stk[len(stk)-1]
			lft := stk[len(stk)-2]
			stk = stk[:len(stk)-2]

			// Perform arithmetic operation.
			var res int
			switch tkn {
			case "+":
				res = lft + rht
			case "-":
				res = lft - rht
			case "*":
				res = lft * rht
			case "/":
				res = lft / rht
			}
			// Push intermediate result to stack.
			stk = append(stk, res)
		default:
			// Parse string to integer.
			num, _ := strconv.Atoi(tkn)
			// Push integer to stack.
			stk = append(stk, num)
		}
	}

	// Return final result.
	return stk[0]
}

// TestEvalRPN contains test cases for the evalRPN function
func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name string
		tkns []string
		want int
	}{
		{
			name: "Simple addition and multiplication",
			tkns: []string{"2", "1", "+", "3", "*"},
			want: 9,
		},
		{
			name: "Division and addition",
			tkns: []string{"4", "13", "5", "/", "+"},
			want: 6,
		},
		{
			name: "Complex expression",
			tkns: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want: 22,
		},
		{
			name: "Single number",
			tkns: []string{"42"},
			want: 42,
		},
		{
			name: "Simple subtraction",
			tkns: []string{"5", "3", "-"},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN(tt.tkns); got != tt.want {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}
