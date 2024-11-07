package main

import (
	"reflect"
	"testing"
)

// Time complexity: O(n), n is the length of the array invs. We traverse the array once.
// Space complexity: O(n), a new return array is allocated.
// https://claude.ai/chat/3e26edd5-c825-468b-a46f-6fc6d6fd5b40
func insert57(invs [][]int, newInv []int) [][]int {
	// Insert Interval
	// Given an array of non-overlapping intervals.
	// An interval is a two-integer array: start index 0, end index 1.
	// Intervals are sorted in ascending order by interval start.
	// Insert new interval.
	// Return the intervals after insertion.
	// Conditions:
	//	* Sort ascending by interval start preserved.
	//	* Any overlapping intervals are merged.
	//	* A new array may be returned.

	// Initialize variables.
	var res [][]int
	idx := 0
	invLen := len(invs)

	// Phase 1: Add intervals which end prior to new interval start.
	for idx < invLen && invs[idx][1] < newInv[0] {
		res = append(res, invs[idx])
		idx++
	}

	// Phase 2: Merge any overlapping intervals into new interval.
	// Overlap: first end >= second start
	for idx < invLen && newInv[1] >= invs[idx][0] {
		// Merge into new interval.
		newInv[0] = min(newInv[0], invs[idx][0])
		newInv[1] = max(newInv[1], invs[idx][1])
		idx++
	}
	// Add new interval.
	res = append(res, newInv)

	// Phase 3: Add remaining intervals.
	for idx < invLen {
		res = append(res, invs[idx])
		idx++
	}

	return res
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Unit Tests
func TestInsert57(t *testing.T) {
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
			got := insert57(tc.intervals, tc.newInt)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("insert() = %v, want %v", got, tc.want)
			}
		})
	}
}
