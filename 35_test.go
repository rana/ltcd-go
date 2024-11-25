package main

import "testing"

// Time complexity: O(log n), n is the length of the array nums. We perform a binary search.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/2473f157-cebf-4872-be93-123b97db4a21
func searchInsert35(nums []int, tgt int) int {
	// Search Insert Position
	// Given a sort-ascending integer array.
	// Given an integer tgt.
	// Determine the index of tgt in nums; or, it's insert index.
	// Conditions:
	//	* Time complexity: O(log n)
	// Use binary search.

	// Initialize binary search bounds.
	lft, rht := 0, len(nums)-1

	// Perform binary search through nums for tgt.
	for lft <= rht {
		// Calculate middle index.
		mid := lft + (rht-lft)/2

		// Check for tgt index.
		switch {
		case nums[mid] == tgt:
			// Found exact match.
			return mid
		case nums[mid] < tgt:
			// Target is in right half.
			lft = mid + 1
		default:
			// Target is in left half.
			rht = mid - 1
		}
	}

	// No exact match. Return lft as insert index.
	return lft
}

// Test cases
func TestSearchInsert35(t *testing.T) {
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
			got := searchInsert35(tc.nums, tc.tgt)
			if got != tc.want {
				t.Errorf("searchInsert(%v, %d) = %d; want %d", tc.nums, tc.tgt, got, tc.want)
			}
		})
	}
}
