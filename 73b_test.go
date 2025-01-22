package main

import "testing"

func setZeroes73b(mtx [][]int) {

}

func TestSetZeroes73b(t *testing.T) {
	tests := []struct {
		name string
		mtx  [][]int
		want [][]int
	}{
		{
			name: "3x3 matrix with center zero",
			mtx:  [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			want: [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			name: "3x4 matrix with two zeros",
			mtx:  [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			want: [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
		{
			name: "1x1 matrix with zero",
			mtx:  [][]int{{0}},
			want: [][]int{{0}},
		},
		{
			name: "2x2 matrix with no zeros",
			mtx:  [][]int{{1, 1}, {1, 1}},
			want: [][]int{{1, 1}, {1, 1}},
		},
		{
			name: "2x3 matrix with zeros in first row and column",
			mtx:  [][]int{{0, 1, 2}, {3, 4, 5}},
			want: [][]int{{0, 0, 0}, {0, 4, 5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of input matrix
			mtx := make([][]int, len(tt.mtx))
			for i := range tt.mtx {
				mtx[i] = make([]int, len(tt.mtx[i]))
				copy(mtx[i], tt.mtx[i])
			}

			setZeroes73b(mtx)

			// Compare result with expected output
			for i := range tt.want {
				for j := range tt.want[i] {
					if mtx[i][j] != tt.want[i][j] {
						t.Errorf("At position [%d][%d], got %d, want %d",
							i, j, mtx[i][j], tt.want[i][j])
					}
				}
			}
		})
	}
}
