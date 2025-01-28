package main

import (
	"reflect"
	"testing"
)

// Time Complexity:
// Space Complexity:
func matrixReshape566b(mtx [][]int, rows int, cols int) [][]int {
	return nil
}

func TestMatrixReshape566b(t *testing.T) {
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
			got := matrixReshape566b(tc.mtx, tc.rows, tc.cols)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}
