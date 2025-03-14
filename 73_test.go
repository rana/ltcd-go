package main

import "testing"

// Time complexity: O(m * n), m and n are the matrix dimensions. We traverse the array twice. Once for determining zero presence. A second time to set zeros.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/8b895b48-7701-4bb0-a48a-da43537a00cf
func setZeroes73(mtx [][]int) {
	// Set Matrix Zeroes
	// Given an m x n integer matrix mtx.
	// Transform in-place.
	// Conditions:
	// * If element is zero, set row zero.
	// * If element is zero, set col zero.
	// Use first row/col to store element states.
	// Use two bools to store first row/col state.
	// Use two steps:
	// 1. Identify and store element zero state.
	// 2. Update non-first row/col element state to zero.

	// Get dimensions.
	rows := len(mtx)
	cols := len(mtx[0])

	// Use first row and column as markers.
	fstRowZero := false
	fstColZero := false

	// Detect whether the first column has any zeros.
	for row := range rows {
		if mtx[row][0] == 0 {
			fstColZero = true
			break
		}
	}

	// Detect whether the first row has any zeros.
	for col := range cols {
		if mtx[0][col] == 0 {
			fstRowZero = true
			break
		}
	}

	// Determine whether interior matrix has any zeros.
	for row := 1; row < rows; row++ {
		for col := 1; col < cols; col++ {
			if mtx[row][col] == 0 {
				mtx[row][0] = 0
				mtx[0][col] = 0
			}
		}
	}

	// Set interior matrix.
	for row := 1; row < rows; row++ {
		for col := 1; col < cols; col++ {
			if mtx[row][0] == 0 || mtx[0][col] == 0 {
				mtx[row][col] = 0
			}
		}
	}

	// Set first row zero if needed.
	if fstRowZero {
		for col := range cols {
			mtx[0][col] = 0
		}
	}

	// Set first column zero if needed.
	if fstColZero {
		for row := range rows {
			mtx[row][0] = 0
		}
	}
}

func TestSetZeroes73(t *testing.T) {
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

			setZeroes73(mtx)

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
