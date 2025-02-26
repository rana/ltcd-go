package main

import (
	"reflect"
	"testing"
)

func insert57b(invs [][]int, newInv []int) [][]int {
	return nil
}

func TestInsert57b(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		newInt    []int
		want      [][]int
	}{
		{
			name:      "Example 1",
			intervals: [][]int{{1, 3}, {6, 9}},
			newInt:    []int{2, 5},
			want:      [][]int{{1, 5}, {6, 9}},
		},
		{
			name:      "Example 2",
			intervals: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			newInt:    []int{4, 8},
			want:      [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name:      "Empty intervals",
			intervals: [][]int{},
			newInt:    []int{5, 7},
			want:      [][]int{{5, 7}},
		},
		{
			name:      "New interval at start",
			intervals: [][]int{{3, 5}, {6, 9}},
			newInt:    []int{1, 2},
			want:      [][]int{{1, 2}, {3, 5}, {6, 9}},
		},
		{
			name:      "New interval at end",
			intervals: [][]int{{1, 2}, {3, 5}},
			newInt:    []int{6, 8},
			want:      [][]int{{1, 2}, {3, 5}, {6, 8}},
		},
		{
			name:      "New interval covers all",
			intervals: [][]int{{3, 5}, {7, 9}, {11, 13}},
			newInt:    []int{2, 14},
			want:      [][]int{{2, 14}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := insert57b(tc.intervals, tc.newInt)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("insert() = %v, want %v", got, tc.want)
			}
		})
	}
}
