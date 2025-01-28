package main

import "testing"

// Time Complexity: O(m * n)
// Space Complexity: O(1)
// https://console.anthropic.com/workbench/b484e06f-decc-4a35-8927-047ff7353a00
func isToeplitzMatrix766(mtx [][]int) bool {
	rows, cols := len(mtx), len(mtx[0])
	for row := 1; row < rows; row++ {
		for col := 1; col < cols; col++ {
			if mtx[row][col] != mtx[row-1][col-1] {
				return false
			}
		}
	}
	return true
}

func TestIsToeplitzMatrix766(t *testing.T) {
	cases := []struct {
		name string
		mtx  [][]int
		want bool
	}{
		{
			name: "Example 1 - Valid Toeplitz",
			mtx: [][]int{
				{1, 2, 3, 4},
				{5, 1, 2, 3},
				{9, 5, 1, 2},
			},
			want: true,
		},
		{
			name: "Example 2 - Invalid Toeplitz",
			mtx: [][]int{
				{1, 2},
				{2, 2},
			},
			want: false,
		},
		{
			name: "Single Row",
			mtx:  [][]int{{1, 2, 3}},
			want: true,
		},
		{
			name: "Single Column",
			mtx:  [][]int{{1}, {2}, {3}},
			want: true,
		},
		{
			name: "Single Element",
			mtx:  [][]int{{1}},
			want: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := isToeplitzMatrix766(tc.mtx)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
