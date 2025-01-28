package main

import (
	"reflect"
	"testing"
)

// Time Complexity:
// Space Complexity:
func generateMatrix59b(n int) [][]int {
	return nil
}

func TestGenerateMatrix59b(t *testing.T) {
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
			got := generateMatrix59b(tc.n)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("generateMatrix(%v) = %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}
