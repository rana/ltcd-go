package main

import "testing"

// Time complexity: O(log n), n is the length of array nums. We binary search, dividing the array in half at each iteration.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/20613a4a-a323-4826-830a-83598f9318fe
func findMin153(nums []int) int {
	// Find Minimum in Roatated Sorted Array
	// Given a sort-ascending integer array.
	// Find the min element.
	// Return the min element.
	// Conditions:
	//	* Array is circularly rotated.
	//	* Time complexity: O(log n)
	// Use modified binary search.
	// 	* for lft < rht
	//	* rht = mid

	// Initialize binary search bounds.
	lft, rht := 0, len(nums)-1

	// Check non-rotated case.
	if nums[lft] < nums[rht] {
		return nums[lft]
	}

	// Binary search while in bounds.
	// Modified:
	//	* for lft < rht
	// 	* rht = mid
	for lft < rht {
		// Calculate middle index.
		mid := lft + (rht-lft)/2

		// Check side of min.
		if nums[mid] > nums[rht] {
			// Min on right side.
			// Search right.
			lft = mid + 1
		} else {
			// Min on left side.
			// Search left.
			rht = mid
		}
	}

	return nums[lft]
}

func TestFindMin153(t *testing.T) {
	cases := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1: Rotated 3 times",
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		{
			name: "Example 2: Rotated 4 times",
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
		{
			name: "Example 3: No rotation",
			nums: []int{11, 13, 15, 17},
			want: 11,
		},
		{
			name: "Single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "Two elements rotated",
			nums: []int{2, 1},
			want: 1,
		},
		{
			name: "Two elements not rotated",
			nums: []int{1, 2},
			want: 1,
		},
		{
			name: "Rotated once",
			nums: []int{2, 1},
			want: 1,
		},
		{
			name: "Large rotation",
			nums: []int{5, 6, 7, 8, 9, 1, 2, 3, 4},
			want: 1,
		},
		{
			name: "Negative numbers",
			nums: []int{2, 3, 4, -5, -4, -3, -2},
			want: -5,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := findMin153(tc.nums)
			if got != tc.want {
				t.Errorf("findMin(%v) = %v, want %v", tc.nums, got, tc.want)
			}
		})
	}
}
