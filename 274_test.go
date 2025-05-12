package main

import (
	"sort"
	"testing"
)

// Time complexity: O(n log n), n is the length of array ctes. We sort the array once. We traverse the array for h-index determination once.
// Space complexity: O(1), constant additional space used.
func hIndex274(ctes []int) int {
	// H-Index
	// Given an integer array ctes.
	// Determine H-Index.
	// Return H-Index.
	// Conditions:
	// * H-Index: max value element value >= array index
	// * H-Index is 1-based
	// * Order descending by definition
	//  val: [9 7 6 2 1]
	// hidx: [1 2 3 4 5] (1-based index)

	sort.Sort(sort.Reverse(sort.IntSlice(ctes)))
	for idx, cte := range ctes {
		// hidx is 1-based
		hidx := idx + 1
		// Check for h-index condition.
		if cte < hidx {
			return idx
		}
	}
	// All elements meet h-index condition.
	return len(ctes)
}

func TestHIndex274(t *testing.T) {
	tests := []struct {
		name      string
		citations []int
		expected  int
	}{
		{
			name:      "Multiple citations with h-index 3",
			citations: []int{3, 0, 6, 1, 5},
			expected:  3,
		},
		{
			name:      "Three citations with h-index 1",
			citations: []int{1, 3, 1},
			expected:  1,
		},
		{
			name:      "Single zero citation",
			citations: []int{0},
			expected:  0,
		},
		{
			name:      "Single high citation",
			citations: []int{100},
			expected:  1,
		},
		{
			name:      "Multiple zero citations",
			citations: []int{0, 0, 0},
			expected:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := hIndex274(tt.citations)
			if result != tt.expected {
				t.Errorf("hIndex(%v) = %d; want %d", tt.citations, result, tt.expected)
			}
		})
	}
}
