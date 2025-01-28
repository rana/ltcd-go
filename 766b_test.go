package main

import "testing"

// Time Complexity:
// Space Complexity:
func isToeplitzMatrix766b(mtx [][]int) bool {
	return false
}

func TestIsToeplitzMatrix766b(t *testing.T) {
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
			got := isToeplitzMatrix766b(tc.mtx)
			if got != tc.want {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
