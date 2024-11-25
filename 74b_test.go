package main

import "testing"

// Time complexity: O(log(m * n)), m and n are dimensions of the matrix. We perform binary search dividing the matrix in half at each iteration.
// Space complexity: O(1), constant additional space used.
func searchMatrix74b(mtx [][]int, tgt int) bool {
	// Search a 2D Matrix
	// Given an m x n integer matrix `mtx`.
	// Given an integer target `tgt`.
	// Determine whether `tgt` is in `mtx`.
	// Return true if conditions met.
	// Conditions:
	//	* Each row is sort-ascending.
	//	* Each first column is sort-ascending.
	//	* Time complexity: O(log(m * n))
	// Transform 2D coordinate to 1D.
	// Treat matrix as single sort-ascending array.
	// Use binary search.

	// Check input min edge case.
	if len(mtx) == 0 || len(mtx[0]) == 0 {
		return false
	}

	// Initialize variables.
	row, col := len(mtx), len(mtx[0])
	lft, rht := 0, row*col-1

	// Iterate through matrix as single array.
	for lft <= rht {
		// Calculate middle index.
		mid := lft + (rht-lft)/2

		// Transform 1D to 2D coordinate.
		curRow, curCol := mid/col, mid%col

		// Get current value.
		curVal := mtx[curRow][curCol]

		// Check for target.
		switch {
		case curVal == tgt:
			// Exact match.
			return true
		case curVal < tgt:
			// Target is on right side.
			lft = mid + 1
		default:
			// Target is on left side.
			rht = mid - 1
		}
	}

	// No match.
	return false
}

func TestSearchMatrix74b(t *testing.T) {
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
			got := searchMatrix74b(tt.matrix, tt.target)
			if got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
