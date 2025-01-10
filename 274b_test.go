package main

import (
	"sort"
	"testing"
)

func hIndex274b(ctes []int) int {
	// H-Index
	// Given an integer array ctes.
	// Determine the H-Index.
	// Return the H-Index.
	// Conditions:
	// * H-Index is 1-based.
	// * H-Index element is great or equal to 1-based index.
	// * Select maximum
	// Use sort-descending.

	sort.Sort(sort.Reverse(sort.IntSlice(ctes)))
	for idx, cte := range ctes {
		if cte < idx+1 {
			return idx
		}
	}
	return len(ctes)
}

func TestHIndex274b(t *testing.T) {
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
			result := hIndex274b(tt.citations)
			if result != tt.expected {
				t.Errorf("hIndex(%v) = %d; want %d", tt.citations, result, tt.expected)
			}
		})
	}
}
