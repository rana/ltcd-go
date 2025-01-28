package main

import (
	"reflect"
	"testing"
)

// Time Complexity:
// Space Complexity:
func findDiagonalOrder498b(mtx [][]int) []int {
	return nil
}

func TestFindDiagonalOrder498b(t *testing.T) {
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
			got := findDiagonalOrder498b(tt.mat)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDiagonalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
