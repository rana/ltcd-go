package main

import (
	"math"
	"testing"
)

// Time complexity:
// Space complexity:
func findMaxAverage643b(nums []int, k int) float64 {
	return 0.0
}

func TestFindMaxAverage643b(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected float64
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 12, -5, -6, 50, 3},
			k:        4,
			expected: 12.75,
		},
		{
			name:     "Example 2",
			nums:     []int{5},
			k:        1,
			expected: 5.0,
		},
		{
			name:     "All negative numbers",
			nums:     []int{-1, -2, -3, -4, -5},
			k:        3,
			expected: -2.0,
		},
		{
			name:     "Exactly k elements",
			nums:     []int{1, 2, 3, 4},
			k:        4,
			expected: 2.5,
		},
		{
			name:     "Maximum at the beginning",
			nums:     []int{10, 10, 10, 1, 1, 1},
			k:        3,
			expected: 10.0,
		},
		{
			name:     "Maximum at the end",
			nums:     []int{1, 1, 1, 10, 10, 10},
			k:        3,
			expected: 10.0,
		},
		{
			name:     "Larger numbers",
			nums:     []int{-10000, 10000, -10000, 10000},
			k:        2,
			expected: 0.0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := findMaxAverage643b(test.nums, test.k)
			if math.Abs(result-test.expected) > 1e-5 {
				t.Errorf("got %v, want %v", result, test.expected)
			}
		})
	}
}
