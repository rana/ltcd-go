package main

import (
	"reflect"
	"testing"
)

// Time complexity: O(log n), n is the length of the array. Binary search divides the array in half at each iteration.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/83efa608-dcca-44bb-8283-109858eb833a
func searchRange34(nums []int, tgt int) []int {
	// Find First and Last Position of Element in Sorted Array
	// Given a sort-asecnding integer array `nums`.
	// Given an integer target `tgt`.
	// Find the first index and last index of target.
	// Return the two indexes; or, [-1,-1].
	// Conditions:
	//	* Duplicates may exist.
	//	* Time complexity: O(log n)
	// Use a helper function which focuses on first index or last index.
	// Use binary search to achieve O(log n).

	// Check input min edge case.
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	// Search for first index.
	fst := binFnd34(nums, tgt, true)
	if fst == -1 {
		return []int{-1, -1}
	}

	// Search for last index.
	lst := binFnd34(nums, tgt, false)
	return []int{fst, lst}
}

// binFnd34 binary searches for target.
// Supports looking for first index or last index of target.
func binFnd34(nums []int, tgt int, fndFst bool) int {
	// Initialize binary search bounds.
	lft, rht := 0, len(nums)-1

	// Binary search while in bounds.
	for lft <= rht {
		// Calculate middle index.
		mid := lft + (rht-lft)/2

		// Check target.
		switch {
		case nums[mid] == tgt:
			// Target exact match possible.
			if fndFst {
				// Left side.
				// Search for first index of target.
				if mid == 0 || nums[mid-1] != tgt {
					// Exact match.
					return mid
				}
				// Search left side.
				rht = mid - 1
			} else {
				// Right side.
				// Search for last index of target.
				if mid == len(nums)-1 || nums[mid+1] != tgt {
					// Exact match.
					return mid
				}
				// Search right side.
				lft = mid + 1
			}
		case nums[mid] < tgt:
			// Target on right half.
			lft = mid + 1
		default:
			// Target on lft half.
			rht = mid - 1
		}
	}

	// No match
	return -1
}

func TestSearchRange34(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		tgt  int
		want []int
	}{
		{
			name: "Example 1: Found target with multiple occurrences",
			nums: []int{5, 7, 7, 8, 8, 10},
			tgt:  8,
			want: []int{3, 4},
		},
		{
			name: "Example 2: Target not found",
			nums: []int{5, 7, 7, 8, 8, 10},
			tgt:  6,
			want: []int{-1, -1},
		},
		{
			name: "Example 3: Empty array",
			nums: []int{},
			tgt:  0,
			want: []int{-1, -1},
		},
		{
			name: "Single element equals target",
			nums: []int{1},
			tgt:  1,
			want: []int{0, 0},
		},
		{
			name: "Target at array boundaries",
			nums: []int{1, 1, 1},
			tgt:  1,
			want: []int{0, 2},
		},
		{
			name: "Target not present in single element",
			nums: []int{1},
			tgt:  2,
			want: []int{-1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchRange34(tt.nums, tt.tgt)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
