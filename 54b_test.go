package main

import (
	"reflect"
	"testing"
)

// Time complexity:
// Space complexity:
func spiralOrder54b(mtx [][]int) []int {
	return nil
}

func TestSpiralOrder54b(t *testing.T) {
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
			got := spiralOrder54b(tt.mtx)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
