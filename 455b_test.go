package main

import (
	"testing"
)

// Time complexity:
// Space complexity:
func findContentChildren455b(grd []int, siz []int) int {
	return 0
}

func TestFindContentChildren455b(t *testing.T) {
	cases := []struct {
		name     string
		grd      []int
		siz      []int
		expected int
	}{
		{
			name:     "Example 1",
			grd:      []int{1, 2, 3},
			siz:      []int{1, 1},
			expected: 1,
		},
		{
			name:     "Example 2",
			grd:      []int{1, 2},
			siz:      []int{1, 2, 3},
			expected: 2,
		},
		{
			name:     "No cookies",
			grd:      []int{1, 2},
			siz:      []int{},
			expected: 0,
		},
		{
			name:     "No children",
			grd:      []int{},
			siz:      []int{1, 2, 3},
			expected: 0,
		},
		{
			name:     "All children satisfied",
			grd:      []int{1, 2, 3},
			siz:      []int{1, 2, 3, 4},
			expected: 3,
		},
		{
			name:     "No children satisfied",
			grd:      []int{7, 8, 9},
			siz:      []int{1, 2, 3},
			expected: 0,
		},
		{
			name:     "Some children satisfied",
			grd:      []int{1, 3, 5, 7},
			siz:      []int{2, 4, 6},
			expected: 3,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := findContentChildren455b(tc.grd, tc.siz)
			if got != tc.expected {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}
