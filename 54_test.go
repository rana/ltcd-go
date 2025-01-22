package main

import (
	"reflect"
	"testing"
)

// Time complexity: O(m * n), m and n are dimensions of the matrix. We traverse teh matrix once.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/ccd4871c-08d7-4372-a9c2-6ec51c6ca95c
func spiralOrder54(mtx [][]int) []int {
	// Spiral Matrix
	// Given a matrix `mtx`.
	// Traverse the matrix in spiral order.
	// Return a matrix in spiral order.

	// Check input minimum edge cases.
	if len(mtx) == 0 || len(mtx[0]) == 0 {
		return []int{}
	}

	// Initialize variables.
	top := 0
	btm := len(mtx) - 1
	lft := 0
	rht := len(mtx[0]) - 1
	res := make([]int, 0, len(mtx)*len(mtx[0]))

	// Iterate until all elements processed.
	for lft <= rht && top <= btm {
		// top: move right
		for col := lft; col <= rht; col++ {
			res = append(res, mtx[top][col])
		}
		top++

		// right: move down
		for row := top; row <= btm; row++ {
			res = append(res, mtx[row][rht])
		}
		rht--

		if top <= btm {
			// bottom: move left
			for col := rht; col >= lft; col-- {
				res = append(res, mtx[btm][col])
			}
			btm--
		}

		if lft <= rht {
			// left: move up
			for row := btm; row >= top; row-- {
				res = append(res, mtx[row][lft])
			}
			lft++
		}
	}

	return res
}

func TestSpiralOrder54(t *testing.T) {
	tests := []struct {
		name string
		mtx  [][]int
		want []int
	}{
		{
			name: "3x3 matrix",
			mtx:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			name: "3x4 matrix",
			mtx:  [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
		{
			name: "1x1 matrix",
			mtx:  [][]int{{1}},
			want: []int{1},
		},
		{
			name: "empty matrix",
			mtx:  [][]int{},
			want: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := spiralOrder54(tt.mtx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
