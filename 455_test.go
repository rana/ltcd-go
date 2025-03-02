package main

import (
	"sort"
	"testing"
)

// Time complexity: O(nlogn + mlogm), n is the length of the grd array. m is the length of the siz array.
// Space complexity: O(1), constant additional space used.
// https://console.anthropic.com/workbench/cdd93511-b623-4229-b335-d403294b8491
func findContentChildren455(grd []int, siz []int) int {
	// Assign Cookies
	// Given arrays grd and siz.
	// Conditions:
	// * grd[i] is greed factor
	// * siz[j] is cookie size
	// Determine maximum count of content childern.
	// Return maximum count.
	// Use a local optimization "greedy" approach.
	// Sort each array
	// Check for  size meeting greed factor minimum.

	sort.Ints(grd)
	sort.Ints(siz)

	cnt := 0
	grdIdx, sizIdx := 0, 0

	for grdIdx < len(grd) && sizIdx < len(siz) {
		// If current cookie can satisfy current child
		if siz[sizIdx] >= grd[grdIdx] {
			cnt++
			grdIdx++ // Move to next child
		}
		sizIdx++ // Always move to next cookie
	}

	return cnt
}

func TestFindContentChildren455(t *testing.T) {
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
			got := findContentChildren455(tc.grd, tc.siz)
			if got != tc.expected {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}
