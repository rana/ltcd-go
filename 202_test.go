package main

import (
	"testing"
)

// Time complexity: O((logn)^2). One iteration reducing function is logn. Multiple iterations of is logn.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/3e3bb71c-c6cb-44e2-9dbe-73052755d795
func isHappy202(num int) bool {
	// Happy Number
	// Reduce a number to 1; or, detect an infinite cycle.
	// Return true if the reduced number equals 1.
	// Conditions:
	//	* Reduce by summing the square of the digits.
	//	* Repeat until equal to 1; or, infinite cycle.
	// Use a helper function to sum the digits.
	// Use Floyd's Toroise and Hare algorithm to detect cycle.

	// Initialize slow and fast pointers for cycle detection.
	slo, fst := num, num

	// Reduce and detect cycle.
	for {
		slo = sumOfSquares202(slo)
		fst = sumOfSquares202(sumOfSquares202(fst))
		if fst == 1 {
			return true
		}
		if fst == slo {
			return false
		}
	}
}

func sumOfSquares202(num int) int {
	sum := 0
	for num > 0 {
		dig := num % 10
		sum += dig * dig
		num /= 10
	}
	return sum
}

// Test cases
func TestBothApproaches202(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{
			name:     "Example 1: Happy number 19",
			input:    19,
			expected: true,
		},
		{
			name:     "Example 2: Unhappy number 2",
			input:    2,
			expected: false,
		},
		{
			name:     "Edge case: Smallest happy number 1",
			input:    1,
			expected: true,
		},
		{
			name:     "Edge case: Large happy number 100",
			input:    100,
			expected: true,
		},
		{
			name:     "Common unhappy number 4",
			input:    4,
			expected: false,
		},
		{
			name:     "Happy number 7",
			input:    7,
			expected: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name+" (Floyd)", func(t *testing.T) {
			result := isHappy202(tc.input)
			if result != tc.expected {
				t.Errorf("For input %d, expected %v but got %v",
					tc.input, tc.expected, result)
			}
		})
	}
}
