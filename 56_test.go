package main

import (
	"sort"
	"testing"
)

// Time complexity: O(n log n), n is the length of the intervals array. We sort the array for O(n log n), and merge56 in O(n).
// Space complexity: O(log n), quicksort stack allocation.
// https://claude.ai/chat/02696ead-d720-4891-8a4e-c208edd23f36
func merge56(invs [][]int) [][]int {
	// Merge Intervals
	// Given and array of intervals.
	// An interval is a two-integer array; start time and end time.
	// Merge overlapping intervals.
	// Return an array of non-overlapping intervals.
	// Sort array by start time.
	// Overlapping interval is: first end >= second start.

	// Check input minimum edge case.
	if len(invs) <= 1 {
		return invs
	}

	// Sort intervals by start time.
	sort.Slice(invs, func(i, j int) bool {
		return invs[i][0] < invs[j][0]
	})

	// Initialize variables.
	var res [][]int
	curInv := invs[0]

	// Iterate through intervals.
	for idx := 1; idx < len(invs); idx++ {
		nxtInv := invs[idx]

		// Check for interval overlap.
		// Overlap: first end >= second end
		if curInv[1] >= nxtInv[0] {
			// Update max end to current.
			if nxtInv[1] > curInv[1] {
				curInv[1] = nxtInv[1]
			}
		} else {
			// No overlap.
			// Add current interval.
			res = append(res, curInv)
			// Move to next interval.
			curInv = nxtInv
		}
	}

	// Add last interval.
	res = append(res, curInv)

	return res
}

func TestMerge56(t *testing.T) {
	areEqual := func(a, b [][]int) bool {
		if len(a) != len(b) {
			return false
		}
		for i := range a {
			if a[i][0] != b[i][0] || a[i][1] != b[i][1] {
				return false
			}
		}
		return true
	}
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
		{
			name:     "Example 1 - Multiple overlapping intervals",
			input:    [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:     "Example 2 - Two touching intervals",
			input:    [][]int{{1, 4}, {4, 5}},
			expected: [][]int{{1, 5}},
		},
		{
			name:     "Single interval",
			input:    [][]int{{1, 4}},
			expected: [][]int{{1, 4}},
		},
		{
			name:     "No overlap",
			input:    [][]int{{1, 2}, {3, 4}, {5, 6}},
			expected: [][]int{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name:     "Complete overlap",
			input:    [][]int{{1, 6}, {2, 3}, {3, 4}},
			expected: [][]int{{1, 6}},
		},
		{
			name:     "Unsorted input",
			input:    [][]int{{3, 6}, {1, 2}, {8, 10}},
			expected: [][]int{{1, 2}, {3, 6}, {8, 10}},
		},
		{
			name:     "Empty input",
			input:    [][]int{},
			expected: [][]int{},
		},
		{
			name:     "Same intervals",
			input:    [][]int{{1, 1}, {1, 1}},
			expected: [][]int{{1, 1}},
		},
		{
			name:     "Zero length intervals",
			input:    [][]int{{1, 1}, {2, 2}, {3, 3}},
			expected: [][]int{{1, 1}, {2, 2}, {3, 3}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := merge56(tc.input)
			if !areEqual(got, tc.expected) {
				t.Errorf("merge() = %v, want %v", got, tc.expected)
			}
		})
	}
}
