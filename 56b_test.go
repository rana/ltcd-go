package main

import (
	"testing"
)

// Time complexity:
// Space complexity:
func merge56b(invs [][]int) [][]int {
	return nil
}

func TestMerge56b(t *testing.T) {
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
			got := merge56b(tc.input)
			if !areEqual(got, tc.expected) {
				t.Errorf("merge() = %v, want %v", got, tc.expected)
			}
		})
	}
}
