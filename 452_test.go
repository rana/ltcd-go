package main

import (
	"sort"
	"testing"
)

// Time complexity: O(n log n), n is the length of the invs array. We quicksort the array on for O(n log n). We traverse the array once for O(n).
// Space complexity: O(log n), quicksort recursive call use call stack space of O(log n).
// https://claude.ai/chat/a3cd11ad-5576-48e3-ab70-dc27df0be1cf
func findMinArrowShots452(invs [][]int) int {
	// Minimum Number of Arrows to Burst Ballons
	// Given an interval array invs.
	// An interval is a two-integer array.
	// An interval start is index 0, and interval end is index 1.
	// An interval represents a ballon width.
	// Shoot an arrow from bellow a ballon.
	// Determine the minimum number of arrows to pop all ballons.
	// Return the minimum number of arrows.
	// Sort intervals by interval end point.
	// Count arrow by transition between two non-overlapping ballons.
	// Use a local optimization "greedy" approach.

	// Sort intervals by end point for optimal arrow counting.
	sort.Slice(invs, func(i, j int) bool {
		return invs[i][1] < invs[j][1]
	})

	// Initialize first end point.
	end := invs[0][1]
	cnt := 1

	// Iterate through remaining ballons.
	for idx := 1; idx < len(invs); idx++ {
		// Check whether previous ballon doesn't overlap with current.
		if end < invs[idx][0] {
			// Increment arrow count
			cnt++
			// Move end to current ballon
			end = invs[idx][1]
		}
		// Otherwise, overlapping ballons use the same arrow.
	}

	return cnt
}

// Unit Tests
func TestFindMinArrowShots452(t *testing.T) {
	tests := []struct {
		name     string
		points   [][]int
		expected int
	}{
		{
			name: "Example 1 - Overlapping balloons",
			points: [][]int{
				{10, 16}, {2, 8}, {1, 6}, {7, 12},
			},
			expected: 2,
		},
		{
			name: "Example 2 - Non-overlapping balloons",
			points: [][]int{
				{1, 2}, {3, 4}, {5, 6}, {7, 8},
			},
			expected: 4,
		},
		{
			name: "Example 3 - Sequential overlapping",
			points: [][]int{
				{1, 2}, {2, 3}, {3, 4}, {4, 5},
			},
			expected: 2,
		},
		{
			name: "Single balloon",
			points: [][]int{
				{1, 2},
			},
			expected: 1,
		},
		{
			name: "All balloons overlap",
			points: [][]int{
				{1, 5}, {2, 4}, {3, 3}, {2, 4},
			},
			expected: 1,
		},
		{
			name: "Large coordinate values",
			points: [][]int{
				{-2147483646, -2147483645}, {2147483646, 2147483647},
			},
			expected: 2,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := findMinArrowShots452(tc.points)
			if got != tc.expected {
				t.Errorf("findMinArrowShots() = %v, want %v", got, tc.expected)
			}
		})
	}
}
