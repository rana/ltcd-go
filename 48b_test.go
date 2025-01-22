package main

import (
	"reflect"
	"testing"
)

// Time complexity:
// Space complexity:
func rotate48b(mtx [][]int) {

}

func TestRotate48b(t *testing.T) {
	tests := []struct {
		name string
		mtx  [][]int
		want [][]int
	}{
		{
			name: "3x3 matrix",
			mtx:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			name: "4x4 matrix",
			mtx:  [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			want: [][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
		{
			name: "1x1 matrix",
			mtx:  [][]int{{1}},
			want: [][]int{{1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a deep copy of input matrix
			input := make([][]int, len(tt.mtx))
			for i := range tt.mtx {
				input[i] = make([]int, len(tt.mtx[i]))
				copy(input[i], tt.mtx[i])
			}

			// Call rotate function
			rotate48b(input)

			// Compare with expected result
			if !reflect.DeepEqual(input, tt.want) {
				t.Errorf("rotate() = %v, want %v", input, tt.want)
			}
		})
	}
}
