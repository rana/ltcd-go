package main

import (
	"reflect"
	"testing"
)

// Time Complexity: O(m * n)
// Space Complexity: O(1), constant additional space used.
// https://console.anthropic.com/workbench/5126c006-e66e-42b3-af0f-7af2e76343d1
func matrixReshape566(mtx [][]int, rows int, cols int) [][]int {
	// Reshape the Matrix
	// Given an m x n matrix mtx.
	// Given integers rows and cols.
	// Reshape mtx into a new matrix with dimensions rows and cols.
	// Return the reshaped matrix.
	prvRows, prvCols := len(mtx), len(mtx[0])
	prvLen := prvRows * prvCols
	if prvLen != rows*cols {
		return mtx
	}
	res := make([][]int, rows)
	for idx := range res {
		res[idx] = make([]int, cols)
	}
	for idx := range prvLen {
		prvRow := idx / prvCols
		prvCol := idx % prvCols
		row := idx / cols
		col := idx % cols
		res[row][col] = mtx[prvRow][prvCol]
	}
	return res
}

func TestMatrixReshape566(t *testing.T) {
	tests := []struct {
		name     string
		mtx      [][]int
		rows     int
		cols     int
		expected [][]int
	}{
		{
			name:     "Example 1: 2x2 to 1x4",
			mtx:      [][]int{{1, 2}, {3, 4}},
			rows:     1,
			cols:     4,
			expected: [][]int{{1, 2, 3, 4}},
		},
		{
			name:     "Example 2: Invalid reshape",
			mtx:      [][]int{{1, 2}, {3, 4}},
			rows:     2,
			cols:     4,
			expected: [][]int{{1, 2}, {3, 4}},
		},
		{
			name:     "Single element matrix",
			mtx:      [][]int{{1}},
			rows:     1,
			cols:     1,
			expected: [][]int{{1}},
		},
		{
			name:     "2x2 to 4x1",
			mtx:      [][]int{{1, 2}, {3, 4}},
			rows:     4,
			cols:     1,
			expected: [][]int{{1}, {2}, {3}, {4}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := matrixReshape566(tc.mtx, tc.rows, tc.cols)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}
