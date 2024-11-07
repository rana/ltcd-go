package main

import (
	"sort"
	"testing"
)

func merge(invs [][]int) [][]int {
	// Merge Intervals
	// Given an array of intervals.
	// An interval is a two-integer array; start at index 0, end at index 1.
	// Merge all overlapping intervals.
	// Return an array of non-overlapping intervals.
	// Sort array by interval start.
	// Overlap: first end >= second start.

	// Check minimum edge case.
	if len(invs) <= 1 {
		return invs
	}

	// Sort intervals array by interval start.
	sort.Slice(invs, func(i, j int) bool {
		return invs[i][0] < invs[j][0]
	})

	// Initialize variables.
	var res [][]int
	curInv := invs[0]

	// Iterate through remaining intervals.
	for idx := 1; idx < len(invs); idx++ {
		// Get next interval.
		nxtInv := invs[idx]

		// Check interval overlap.
		if curInv[1] >= nxtInv[0] {
			// Expand current interval.
			// Update max end.
			if nxtInv[1] > curInv[1] {
				curInv[1] = nxtInv[1]
			}
		} else {
			// No overlap.
			// Add interval.
			res = append(res, curInv)
			// Advance to next interval.
			curInv = nxtInv
		}
	}

	// Append last interval.
	res = append(res, curInv)

	return res
}

func TestMerge_ValidCases(t *testing.T) {
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
			got := merge(tc.input)
			if !areEqual(got, tc.expected) {
				t.Errorf("merge() = %v, want %v", got, tc.expected)
			}
		})
	}
}

// TestMerge_EdgeCases tests the merge function with edge cases
func TestMerge_EdgeCases(t *testing.T) {
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
			got := merge(tc.input)
			if !areEqual(got, tc.expected) {
				t.Errorf("merge() = %v, want %v", got, tc.expected)
			}
		})
	}
}

// areEqual compares two interval slices for equality
func areEqual(a, b [][]int) bool {
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
