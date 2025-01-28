package main

import "testing"

// Time complexity: O(logn^2)
// Space complexity: O(1)
func isHappy(num int) bool {
	return false
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
