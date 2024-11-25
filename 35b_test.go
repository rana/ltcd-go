package main

import "testing"

// Time complexity: O(log n), n is the length of the array. We binary search which divides the array in half at each iteration.
// Space complexity: O(1), constant additional space used.
func searchInsert35b(nums []int, tgt int) int {
	// Search Insert Position
	// Given sort-ascending integer array nums.
	// Given an integer tgt.
	// Determine the index of tgt in nums; or, the insert index.
	// Return the index.
	// Use binary search.

	// Initialize binary search bounds.
	lft, rht := 0, len(nums)-1

	// Binary search while in bounds.
	for lft <= rht {
		// Calculate middle index.
		mid := lft + (rht-lft)/2

		// Search for target.
		switch {
		case nums[mid] == tgt:
			// Exact match.
			return mid
		case nums[mid] < tgt:
			// Target on right side.
			lft = mid + 1
		default:
			// Target on left side.
			rht = mid - 1
		}
	}

	// No match; insert index is lft
	return lft
}

func TestSearchInsert35b(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		tgt  int
		want int
	}{
		// Base cases from problem description
		{
			name: "Found target in middle",
			nums: []int{1, 3, 5, 6},
			tgt:  5,
			want: 2,
		},
		{
			name: "Insert target in middle",
			nums: []int{1, 3, 5, 6},
			tgt:  2,
			want: 1,
		},
		{
			name: "Insert target at end",
			nums: []int{1, 3, 5, 6},
			tgt:  7,
			want: 4,
		},

		// Edge cases
		{
			name: "Single element, target smaller",
			nums: []int{1},
			tgt:  0,
			want: 0,
		},
		{
			name: "Single element, target larger",
			nums: []int{1},
			tgt:  2,
			want: 1,
		},
		{
			name: "Single element, target equal",
			nums: []int{1},
			tgt:  1,
			want: 0,
		},
		{
			name: "Insert at start",
			nums: []int{2, 3, 4, 5},
			tgt:  1,
			want: 0,
		},
		{
			name: "All elements smaller than target",
			nums: []int{1, 2, 3, 4},
			tgt:  5,
			want: 4,
		},
		{
			name: "All elements larger than target",
			nums: []int{2, 3, 4, 5},
			tgt:  1,
			want: 0,
		},

		// Boundary cases
		{
			name: "Target at minimum constraint",
			nums: []int{-10000, 0, 10000},
			tgt:  -10000,
			want: 0,
		},
		{
			name: "Target at maximum constraint",
			nums: []int{-10000, 0, 10000},
			tgt:  10000,
			want: 2,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := searchInsert35b(tc.nums, tc.tgt)
			if got != tc.want {
				t.Errorf("searchInsert(%v, %d) = %d; want %d", tc.nums, tc.tgt, got, tc.want)
			}
		})
	}
}
