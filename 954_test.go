package main

import (
	"sort"
	"testing"
)

// Time Complexity: O(n log n)
// Space Complexity: O(n)
// https://console.anthropic.com/workbench/dfdd0b01-9d52-45d9-9e8b-f5c16210eec1
func canReorderDoubled954(arr []int) bool {
	// Count frequencies
	frq := make(map[int]int)
	for _, num := range arr {
		frq[num]++
	}

	// Handle zero separately
	if frq[0]%2 != 0 {
		return false
	}

	// Get unique numbers and sort by absolute value
	nums := make([]int, 0, len(frq))
	for num := range frq {
		nums = append(nums, num)
	}
	sort.Slice(nums, func(i, j int) bool {
		return abs(nums[i]) < abs(nums[j])
	})

	// Try to pair each number with its double
	for _, num := range nums {
		if frq[num] > frq[2*num] {
			return false
		}
		frq[2*num] -= frq[num]
	}

	return true
}

func TestCanReorderDoubled954(t *testing.T) {
	cases := []struct {
		name     string
		arr      []int
		expected bool
	}{
		{
			name:     "Example 1",
			arr:      []int{3, 1, 3, 6},
			expected: false,
		},
		{
			name:     "Example 2",
			arr:      []int{2, 1, 2, 6},
			expected: false,
		},
		{
			name:     "Example 3",
			arr:      []int{4, -2, 2, -4},
			expected: true,
		},
		{
			name:     "Zero case",
			arr:      []int{0, 0},
			expected: true,
		},
		{
			name:     "Odd number of zeros",
			arr:      []int{0, 0, 0},
			expected: false,
		},
		{
			name:     "All same numbers",
			arr:      []int{2, 2, 2, 2},
			expected: false,
		},
		{
			name:     "Mixed positive negative",
			arr:      []int{-8, 4, -4, 8},
			expected: true,
		},
		{
			name:     "Empty array",
			arr:      []int{},
			expected: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := canReorderDoubled954(tc.arr)
			if got != tc.expected {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}
