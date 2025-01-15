package main

import "testing"

// Time complexity:
// Space complexity:
func maxProfit122b(prcs []int) int {
	return 0
}

func TestMaxProfit122b(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "Ascending prices",
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			name:     "Descending prices",
			prices:   []int{5, 4, 3, 2, 1},
			expected: 0,
		},
		{
			name:     "Mixed prices",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 7,
		},
		{
			name:     "Single day",
			prices:   []int{5},
			expected: 0,
		},
		{
			name:     "Empty prices",
			prices:   []int{},
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit122b(tt.prices)
			if got != tt.expected {
				t.Errorf("maxProfit122() = %d, want %d", got, tt.expected)
			}
		})
	}
}
