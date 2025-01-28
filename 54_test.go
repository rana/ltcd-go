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
	var res []int
	top, btm := 0, len(mtx)-1
	lft, rht := 0, len(mtx[0])-1
	mtxLen := len(mtx) * len(mtx[0])

	for len(res) < mtxLen {
		// top: move right
		for col := lft; col <= rht; col++ {
			res = append(res, mtx[top][col])
		}
		top++
		// rht: move down
		for row := top; row <= btm; row++ {
			res = append(res, mtx[row][rht])
		}
		rht--
		// btm: move left
		if top <= btm {
			for col := rht; col >= lft; col-- {
				res = append(res, mtx[btm][col])
			}
			btm--
		}
		// lft: move up
		if lft <= rht {
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
