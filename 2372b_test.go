package main

import (
	"reflect"
	"testing"
)

// Time Complexity:
// Space Complexity:
func largestLocal2372b(grd [][]int) [][]int {
	return nil
}

func TestLargestLocal2372b(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want [][]int
	}{
		{
			name: "Example 1",
			grid: [][]int{
				{9, 9, 8, 1},
				{5, 6, 2, 6},
				{8, 2, 6, 4},
				{6, 2, 2, 2},
			},
			want: [][]int{
				{9, 9},
				{8, 6},
			},
		},
		{
			name: "Example 2",
			grid: [][]int{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 2, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
			},
			want: [][]int{
				{2, 2, 2},
				{2, 2, 2},
				{2, 2, 2},
			},
		},
		{
			name: "Minimum size 3x3",
			grid: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			want: [][]int{
				{9},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestLocal2372b(tt.grid)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("largestLocal() = %v, want %v", got, tt.want)
			}
		})
	}
}
