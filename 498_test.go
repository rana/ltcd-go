package main

import (
	"reflect"
	"testing"
)

// Time Complexity: O(m*n) where m and n are the dimensions of the matrix.
// Space Complexity: O(1), constant additional space used.
// https://console.anthropic.com/workbench/b80d7c88-eb16-460a-82cf-ace4a7afc35f
func findDiagonalOrder498(mtx [][]int) []int {
	// Diagonal Traverse
	// Given an m x n matrix mtx.
	// Traverse matrix in diagonal order.
	// Return integer array of traversal.
	// Conditions:
	// * Traverse right-up
	// * Traverse left-down
	// * Repeat

	rows, cols := len(mtx), len(mtx[0])
	top, btm := 0, rows-1
	lft, rht := 0, cols-1
	upRight := true
	resLen := rows * cols
	res := make([]int, 0, resLen)
	row, col := 0, 0
	for len(res) < resLen {
		res = append(res, mtx[row][col])
		if upRight {
			// up-right
			if col == rht {
				row++
				upRight = false
			} else if row == top {
				col++
				upRight = false
			} else {
				row--
				col++
			}
		} else {
			// down-left
			if row == btm {
				col++
				upRight = true
			} else if col == lft {
				row++
				upRight = true
			} else {
				row++
				col--
			}
		}
	}
	return res
}

func TestFindDiagonalOrder498(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		want []int
	}{
		{
			name: "3x3 matrix",
			mat:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want: []int{1, 2, 4, 7, 5, 3, 6, 8, 9},
		},
		{
			name: "2x2 matrix",
			mat:  [][]int{{1, 2}, {3, 4}},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "1x1 matrix",
			mat:  [][]int{{1}},
			want: []int{1},
		},
		{
			name: "1x3 matrix",
			mat:  [][]int{{1, 2, 3}},
			want: []int{1, 2, 3},
		},
		{
			name: "3x1 matrix",
			mat:  [][]int{{1}, {2}, {3}},
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDiagonalOrder498(tt.mat)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDiagonalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
