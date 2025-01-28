package main

import (
	"reflect"
	"testing"
)

// Time Complexity: O(n²), We visit each cell exactly once.
// Space Complexity: O(1), constant additional space used.
// https://console.anthropic.com/workbench/ed7d116b-20a3-4841-9229-df208dfeba1b
func generateMatrix59(n int) [][]int {
	// Sprial Matrix II
	// Given a positive integer n.
	// Create an n x n matrix.
	// Fill matrix with integers in spiral order.
	// Conditions:
	// * Start at 1
	// top: move right
	// rht: move down
	// btm: move left
	// lft: move up
	mtx := make([][]int, n)
	for idx := range mtx {
		mtx[idx] = make([]int, n)
	}
	top, btm := 0, len(mtx)-1
	lft, rht := 0, len(mtx[0])-1
	num, numLen := 1, n*n
	for num <= numLen {
		// top: move right
		for col := lft; col <= rht; col++ {
			mtx[top][col] = num
			num++
		}
		top++
		// rht: move down
		for row := top; row <= btm; row++ {
			mtx[row][rht] = num
			num++
		}
		rht--
		// btm: move left
		for col := rht; col >= lft; col-- {
			mtx[btm][col] = num
			num++
		}
		btm--
		// lft: move up
		for row := btm; row >= top; row-- {
			mtx[row][lft] = num
			num++
		}
		lft++
	}
	return mtx
}

func TestGenerateMatrix59(t *testing.T) {
	cases := []struct {
		name string
		n    int
		want [][]int
	}{
		{
			name: "3x3 matrix",
			n:    3,
			want: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			name: "1x1 matrix",
			n:    1,
			want: [][]int{{1}},
		},
		{
			name: "2x2 matrix",
			n:    2,
			want: [][]int{
				{1, 2},
				{4, 3},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := generateMatrix59(tc.n)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("generateMatrix(%v) = %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}
