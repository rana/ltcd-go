package main

import (
	"reflect"
	"testing"
)

// Time Complexity: O(n * n) where n is the grid dimension. For each position in the result matrix, we examine a fixed 3×3 window.
// Space Complexity: O(1), constant additional space used.
// https://console.anthropic.com/workbench/3af3ccfe-cfdb-49cc-8504-44bd8702e476
func largestLocal2372(grd [][]int) [][]int {
	// Largest Local Values in a Matrix
	// Given an n x n integer matrix grd.
	// Generate max local matrix.
	// Return max local matrix.
	// Conditions:
	// * Size (n-2)(n-2)
	// * Max is largest in 3 x 3 around element

	n := len(grd)
	resLen := n - 2
	res := make([][]int, resLen)
	for idx := range resLen {
		res[idx] = make([]int, resLen)
	}
	const diam = 3
	for row := range resLen {
		for col := range resLen {
			max := 0
			for i := row; i < row+diam; i++ {
				for j := col; j < col+diam; j++ {
					if grd[i][j] > max {
						max = grd[i][j]
					}
				}
			}
			res[row][col] = max
		}
	}
	return res
}

func TestLargestLocal2372(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want [][]int
	}{
		{
			name: "Example 1",
			grid: [][]int{
				{9, 9, 8, 1},
				{5, 6, 2, 6},
				{8, 2, 6, 4},
				{6, 2, 2, 2},
			},
			want: [][]int{
				{9, 9},
				{8, 6},
			},
		},
		{
			name: "Example 2",
			grid: [][]int{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 2, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
			},
			want: [][]int{
				{2, 2, 2},
				{2, 2, 2},
				{2, 2, 2},
			},
		},
		{
			name: "Minimum size 3x3",
			grid: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{9},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestLocal2372(tt.grid)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestLocal() = %v, want %v", got, tt.want)
			}
		})
	}
}
