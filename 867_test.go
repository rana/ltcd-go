package main

import (
	"reflect"
	"testing"
)

// Time complexity: O(m * n), we traverse mtx once.
// Space complexity: O(1), constant additional space used.
func transpose867(mtx [][]int) [][]int {
	// Transpose Matrix
	// Given an m x n integer matrix mtx.
	// Traspose matrix
	// Return new matrix with transpose867.
	rows, cols := len(mtx), len(mtx[0])
	res := make([][]int, cols)
	for idx := range res {
		res[idx] = make([]int, rows)
	}
	for row := range mtx {
		for col := range mtx[row] {
			res[col][row] = mtx[row][col]
		}
	}
	return res
}

func TestTranspose867(t *testing.T) {
	cases := []struct {
		name string
		mtx  [][]int
		want [][]int
	}{
		{
			name: "3x3 matrix",
			mtx:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want: [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
		},
		{
			name: "2x3 matrix",
			mtx:  [][]int{{1, 2, 3}, {4, 5, 6}},
			want: [][]int{{1, 4}, {2, 5}, {3, 6}},
		},
		{
			name: "1x1 matrix",
			mtx:  [][]int{{1}},
			want: [][]int{{1}},
		},
		{
			name: "1x3 matrix",
			mtx:  [][]int{{1, 2, 3}},
			want: [][]int{{1}, {2}, {3}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := transpose867(tc.mtx)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("transpose() = %v, want %v", got, tc.want)
			}
		})
	}
}
