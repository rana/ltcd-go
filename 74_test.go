package main

import "testing"

// Time complexity: O(logn(m * n)), m and n are dimensions of the matrix.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/b8c0cab4-0161-4902-b456-8eb0c250b185
func searchMatrix74(mtx [][]int, tgt int) bool {
	// Search a 2D Matrix
	// Given an m x n integer matrix mtx.
	// Given an integer tgt.
	// Determine whether tgt is in mtx.
	// Conditions:
	//	* Each matrix row is sort-ascending.
	//	* Each first column is sort-ascending.
	//	* Time complexity: O(log(m * n))
	// Treat 2D matrix as 1D sorted array.
	// Use 2D to 1D coordinate transform with column index.
	// Use binary search.

	// Check input min edge cases.
	if len(mtx) == 0 || len(mtx[0]) == 0 {
		return false
	}

	// Initialize variables.
	row, col := len(mtx), len(mtx[0])
	lft, rht := 0, row*col-1

	// Binary search while in bounds.
	for lft <= rht {
		// Calculate middle index.
		mid := lft + (rht-lft)/2

		// Transform 1D to 2D coordinates.
		curRow, curCol := mid/col, mid%col

		// Get current value in matrix.
		curVal := mtx[curRow][curCol]

		// Check for presence.
		if curVal == tgt {
			// Found exact match.
			return true
		}
		if curVal < tgt {
			// Target is in right half.
			lft = mid + 1
		} else {
			// Target is in left half.
			rht = mid - 1
		}
	}

	return false
}

// Unit tests
func TestSearchMatrix74(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{
			name: "Example 1: Found",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 3,
			want:   true,
		},
		{
			name: "Example 2: Not Found",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
			target: 13,
			want:   false,
		},
		{
			name:   "Empty Matrix",
			matrix: [][]int{},
			target: 1,
			want:   false,
		},
		{
			name: "Single Element: Found",
			matrix: [][]int{
				{1},
			},
			target: 1,
			want:   true,
		},
		{
			name: "Single Row",
			matrix: [][]int{
				{1, 3, 5, 7},
			},
			target: 5,
			want:   true,
		},
		{
			name: "Single Column",
			matrix: [][]int{
				{1},
				{3},
				{5},
			},
			target: 3,
			want:   true,
		},
		{
			name: "Target Smaller Than Any Element",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
			},
			target: 0,
			want:   false,
		},
		{
			name: "Target Larger Than Any Element",
			matrix: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
			},
			target: 100,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchMatrix74(tt.matrix, tt.target)
			if got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
