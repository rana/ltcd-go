package main

import (
	"testing"
)

// Time complexity:
// Space complexity:
func maxProfit121b(prcs []int) int {
	return 0
}

func TestMaxProfit121b(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "Valid case - profit possible",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			name:     "Valid case - decreasing prices",
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			name:     "Valid case - increasing prices",
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "Edge case - single price",
			prices:   []int{5},
			expected: 0,
		},
		{
			name:     "Edge case - all same prices",
			prices:   []int{0, 0, 0, 0},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := maxProfit121b(tt.prices)
			if result != tt.expected {
				t.Errorf("maxProfit1() = %d, want %d", result, tt.expected)
			}
		})
	}
}
