package main

import (
	"sort"
	"testing"
)

// Time complexity: O(n log n), n is the length of the intervals array. We sort the array for O(n log n), and merge56 in O(n).
// Space complexity: O(log n), quicksort stack allocation.
// https://claude.ai/chat/02696ead-d720-4891-8a4e-c208edd23f36
func merge56(intervals [][]int) [][]int {
	// Merge Intervals
	// Given and array of intervals.
	// An interval is a two-integer array; start time and end time.
	// Merge overlapping intervals.
	// Return an array of non-overlapping intervals.
	// Sort array by start time.
	// Overlapping interval is: first end >= second start.

	// Check input minimum edge case.
	if len(intervals) <= 1 {
		return intervals
	}

	// Sort intervals by start time.
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Initialize variables.
	var res [][]int
	curInt := intervals[0]

	// Iterate through intervals.
	for idx := 1; idx < len(intervals); idx++ {
		nxtInt := intervals[idx]

		// Check for interval overlap.
		// Overlap: first end >= second end
		if curInt[1] >= nxtInt[0] {
			// Update max end to current.
			if nxtInt[1] > curInt[1] {
				curInt[1] = nxtInt[1]
			}
		} else {
			// No overlap.
			// Add current interval.
			res = append(res, curInt)
			// Move to next interval.
			curInt = nxtInt
		}
	}

	// Add last interval.
	res = append(res, curInt)

	return res
}

// TestMerge56_ValidCases tests the merge function with valid inputs
func TestMerge56_ValidCases(t *testing.T) {
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
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := merge56(tc.input)
			if !areEqual56(got, tc.expected) {
				t.Errorf("merge() = %v, want %v", got, tc.expected)
			}
		})
	}
}

// TestMerge_EdgeCases tests the merge function with edge cases
func TestMerge56_EdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected [][]int
	}{
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
			if !areEqual56(got, tc.expected) {
				t.Errorf("merge() = %v, want %v", got, tc.expected)
			}
		})
	}
}

// areEqual56 compares two interval slices for equality
func areEqual56(a, b [][]int) bool {
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
