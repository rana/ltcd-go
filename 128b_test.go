package main

import "testing"

// Time complexity:
// Space complexity:
func longestConsecutive128b(nums []int) int {
	return 0
}

func TestLongestConsecutive128b(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1: Basic sequence",
			nums:     []int{100, 4, 200, 1, 3, 2},
			expected: 4,
		},
		{
			name:     "Example 2: Longer sequence with duplicates",
			nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			expected: 9,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "All same numbers",
			nums:     []int{1, 1, 1},
			expected: 1,
		},
		{
			name:     "No consecutive numbers",
			nums:     []int{2, 4, 6, 8},
			expected: 1,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-3, -2, -1, 0, 1},
			expected: 5,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := longestConsecutive128b(tc.nums)
			if got != tc.expected {
				t.Errorf("longestConsecutive(%v) = %v; want %v",
					tc.nums, got, tc.expected)
			}
		})
	}
}
