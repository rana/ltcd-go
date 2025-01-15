package main

import "testing"

// Time complexity:
// Space complexity:
func maxArea11b(hgts []int) int {
	return 0
}

func TestMaxArea11b(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "Example 1",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "Example 2",
			height:   []int{1, 1},
			expected: 1,
		},
		{
			name:     "All Same Height",
			height:   []int{5, 5, 5, 5},
			expected: 15,
		},
		{
			name:     "Increasing Heights",
			height:   []int{1, 2, 3, 4, 5},
			expected: 6,
		},
		{
			name:     "Empty Container",
			height:   []int{0, 0},
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := maxArea11b(tc.height)
			if result != tc.expected {
				t.Errorf("maxArea(%v) = %d; want %d",
					tc.height, result, tc.expected)
			}
		})
	}
}
