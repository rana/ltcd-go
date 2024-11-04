package main

import "testing"

// Time complexity: O((logn)^2). One iteration within sumOfSquares is logn. There are logn iterations within sumOfSquares. Using log base 10.
// Space complexity: O(1), constant additional space used.
func isHappy(num int) bool {
	// Happy Number
	// Given an integer n.
	// Reduce the number to 1; or, detect infinite cycle.
	// Return true if number equals 1.
	// Conditions:
	//	* Reduce number to sum of square of digits.
	//	* Repeat until equal to 1; or, infinite cycle.
	// Use a helper function to reduce number.
	// Use Floyd's Tortoise and Hare to detect cycle.

	// Initialize cycle detection variables.
	slo, fst := num, num

	for {
		slo = sumOfSquares(slo)
		fst = sumOfSquares(sumOfSquares(fst))
		if fst == 1 {
			return true
		}
		if fst == slo {
			return false
		}
	}
}

func sumOfSquares(num int) int {
	sum := 0
	for num > 0 {
		dig := num % 10
		sum += dig * dig
		num /= 10
	}
	return sum
}

func TestBothApproaches(t *testing.T) {
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
			result := isHappy(tc.input)
			if result != tc.expected {
				t.Errorf("For input %d, expected %v but got %v",
					tc.input, tc.expected, result)
			}
		})
	}
}
