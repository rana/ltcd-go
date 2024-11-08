package main

import "testing"

// Time complexity: O(n), n is the length of the string. We traverse the string once.
// Space complexity: O(n), We store
// https://claude.ai/chat/4f9fa457-120d-406b-9926-7a54697b1b16
func isValid(s string) bool {
	// Valid Parentheses
	// Given a string s.
	// s contains (){}[]
	// Determine whether string meets parenthesis conditions.
	// Return true if conditions met.
	// Conditions:
	//	* Brackets are paired with the dame type.
	//	* Bracket ordering maintained (e.g., open then close)
	// Use a stack.
	// Push close brackets on to the stack to validate rules.

	// Check input min edge case.
	if len(s) <= 1 {
		return false
	}

	// Initialize stack.
	var stk []rune

	// Iterate through each character.
	for _, chr := range s {
		// Push for open brackets.
		// Pop for close brackets.
		switch chr {
		case '(':
			// Push closing paren
			stk = append(stk, ')')
		case '[':
			// Push closing square bracket.
			stk = append(stk, ']')
		case '{':
			// Push closing
			stk = append(stk, '}')
		default:
			// Check for valid conditions.
			stkLen := len(stk)
			if stkLen == 0 || chr != stk[stkLen-1] {
				return false
			}
			// Pop last bracket.
			stk = stk[:stkLen-1]
		}
	}

	// Stack to be empty for valid brackets.
	return len(stk) == 0
}

// Test cases
func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want bool
	}{
		{
			name: "single pair",
			str:  "()",
			want: true,
		},
		{
			name: "multiple pairs",
			str:  "()[]{}",
			want: true,
		},
		{
			name: "nested pairs",
			str:  "([{}])",
			want: true,
		},
		{
			name: "mismatched brackets",
			str:  "(]",
			want: false,
		},
		{
			name: "unmatched opening",
			str:  "([",
			want: false,
		},
		{
			name: "unmatched closing",
			str:  ")]",
			want: false,
		},
		{
			name: "wrong order",
			str:  "([)]",
			want: false,
		},
		{
			name: "single character",
			str:  "[",
			want: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isValid(tc.str)
			if got != tc.want {
				t.Errorf("isValid(%q) = %v; want %v", tc.str, got, tc.want)
			}
		})
	}
}
