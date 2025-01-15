package main

import "testing"

// Time complexity: O(n), n is the length of the `hgts` array.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/3bf1bc41-ba7c-4f08-9d49-f472dd6bbd31
func maxArea11(hgts []int) int {
	// Container with Most Water
	// Given an integer array `hgts`.
	// The arary represents line heights.
	// Find two lines which maximize the area of a rectangle.
	// Return the maximum area.
	// Use a two-pointer technique.

	// Initialize variables.
	lft, rht := 0, len(hgts)-1
	maxArea := 0

	// Iterate until two pointers meet.
	for lft < rht {
		// Calculate width.
		wid := rht - lft

		// Calculate height.
		// Use the minimum of the two lines.
		hgt := min(hgts[lft], hgts[rht])

		// Calculate current area.
		curArea := wid * hgt

		// Calculate maximum area.
		if curArea > maxArea {
			maxArea = curArea
		}

		// Move pointers forward.
		// Move the pointer for the shorter line.
		// Looking for a longer line.
		if hgts[lft] < hgts[rht] {
			lft++
		} else {
			rht--
		}
	}

	// Return the maximum area.
	return maxArea
}

func TestMaxArea11(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "Example 1",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "Example 2",
			height:   []int{1, 1},
			expected: 1,
		},
		{
			name:     "All Same Height",
			height:   []int{5, 5, 5, 5},
			expected: 15,
		},
		{
			name:     "Increasing Heights",
			height:   []int{1, 2, 3, 4, 5},
			expected: 6,
		},
		{
			name:     "Empty Container",
			height:   []int{0, 0},
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := maxArea11(tc.height)
			if result != tc.expected {
				t.Errorf("maxArea(%v) = %d; want %d",
					tc.height, result, tc.expected)
			}
		})
	}
}
