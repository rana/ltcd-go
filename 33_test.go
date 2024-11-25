package main

import "testing"

// Time complexity: O(log n), n is the length of array nums. We perform binary search33 which divides the array in half at each iteration.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/7dea12b7-890f-4f46-b18e-66c71bf457b9
func search33(nums []int, tgt int) int {
	// Search in Rotated Array
	// Given a sort-ascending integer array nums.
	// Given an integer tgt.
	// Find the index of tgt in nums.
	// Return the index; or, -1 if not found.
	// Conditions:
	//	* Array is circularly rotated some unknown number.
	//	* Time complexity: O(log n)
	// Use a modified binary search.
	// Binary search through half which is intact sorted array.

	// Check input min edge case.
	if len(nums) == 0 {
		return -1
	}

	// Initialize binary search bounds.
	lft, rht := 0, len(nums)-1

	// Binary search while in bounds.
	for lft <= rht {
		// Calculate middle index.
		mid := lft + (rht-lft)/2

		// Check exact match.
		if nums[mid] == tgt {
			return mid
		}

		// Check which half is intact sorted.
		if nums[lft] <= nums[mid] {
			// Left is intact sorted.
			// Search for target.
			if nums[lft] <= tgt && tgt < nums[mid] {
				// Target in left half.
				rht = mid - 1
			} else {
				// Target in right half.
				lft = mid + 1
			}
		} else {
			// Right half is intact sorted.
			// Search for target.
			if nums[mid] < tgt && tgt <= nums[rht] {
				// Target is in right half.
				lft = mid + 1
			} else {
				// Target is in left half.
				rht = mid - 1
			}
		}
	}

	return -1
}

// Test cases
func TestSearch33(t *testing.T) {
	cases := []struct {
		name     string
		nums     []int
		tgt      int
		expected int
	}{
		{
			name:     "Example 1: Target in right portion",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			tgt:      0,
			expected: 4,
		},
		{
			name:     "Example 2: Target not found",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			tgt:      3,
			expected: -1,
		},
		{
			name:     "Example 3: Single element array",
			nums:     []int{1},
			tgt:      0,
			expected: -1,
		},
		{
			name:     "Target in left portion",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			tgt:      5,
			expected: 1,
		},
		{
			name:     "Not rotated array",
			nums:     []int{1, 2, 3, 4, 5},
			tgt:      3,
			expected: 2,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			tgt:      5,
			expected: -1,
		},
		{
			name:     "Target at rotation point",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			tgt:      4,
			expected: 0,
		},
		{
			name:     "Target is smallest element",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			tgt:      0,
			expected: 4,
		},
		{
			name:     "Target is largest element",
			nums:     []int{4, 5, 6, 7, 0, 1, 2},
			tgt:      7,
			expected: 3,
		},
		{
			name:     "Two element rotated array",
			nums:     []int{3, 1},
			tgt:      1,
			expected: 1,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := search33(tc.nums, tc.tgt)
			if got != tc.expected {
				t.Errorf("search(%v, %d) = %d; want %d",
					tc.nums, tc.tgt, got, tc.expected)
			}
		})
	}
}
