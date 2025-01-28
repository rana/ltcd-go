package main

import (
	"reflect"
	"testing"
)

// Time complexity:
// Space complexity:
func transpose867b(mtx [][]int) [][]int {
	return nil
}

func TestTranspose867b(t *testing.T) {
	cases := []struct {
		name string
		mtx  [][]int
		want [][]int
	}{
		{
			name: "3x3 matrix",
			mtx:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			want: [][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}},
		},
		{
			name: "2x3 matrix",
			mtx:  [][]int{{1, 2, 3}, {4, 5, 6}},
			want: [][]int{{1, 4}, {2, 5}, {3, 6}},
		},
		{
			name: "1x1 matrix",
			mtx:  [][]int{{1}},
			want: [][]int{{1}},
		},
		{
			name: "1x3 matrix",
			mtx:  [][]int{{1, 2, 3}},
			want: [][]int{{1}, {2}, {3}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := transpose867b(tc.mtx)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("transpose() = %v, want %v", got, tc.want)
			}
		})
	}
}
