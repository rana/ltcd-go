package main

import (
	"reflect"
	"testing"
)

func insert57b(invs [][]int, newInv []int) [][]int {
	// Insert Interval
	// Given an array of non-overlapping invervals.
	// An interval is a two-integer array; start at index 0, end at index 1.
	// Array is sorted ascending by interval start.
	// Insert a new interval.
	// Return an array of intervals after insertion.
	// Conditions:
	//	* Maintains sort ascending.
	//	* Maintain non-overlapping. Merge if necessary.
	//	* Can create a new array for processing and return.

	var res [][]int
	idx := 0
	invsLen := len(invs)

	// Phase 1: Add intervals which end before new interval starts.
	for idx < invsLen && invs[idx][1] < newInv[0] {
		res = append(res, invs[idx])
		idx++
	}

	// Phase 2: Merge into new interval as necessary.
	// Overlap: first end >= second start.
	for idx < invsLen && newInv[1] >= invs[idx][0] {
		// Expand new interval
		newInv[0] = min(newInv[0], invs[idx][0])
		newInv[1] = max(newInv[1], invs[idx][1])
		idx++
	}
	// Append new interval after merging.
	res = append(res, newInv)

	// Phase 3: Add remaining intervals.
	for idx < invsLen {
		res = append(res, invs[idx])
		idx++
	}

	return res
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
