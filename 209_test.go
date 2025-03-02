package main

import (
	"math"
	"testing"
)

// Time complexity: O(n), n is the length of the `nums` array.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/cc616183-8721-48b3-9d08-62da32e71d02
func minSubArrayLen209(tgt int, nums []int) int {
	// Minimum Size Subarray Sum
	// Given an array of positive integers `nums`.
	// Given a positive integer `tgt`.
	// Find a min length of a subarray meet a condition.
	// Condition:
	//	* Subarray sum great or equal to target.
	// Return min length; or, zero.
	// Use a sliding window two-pointer technique.

	// Initialize variables.
	minLen := math.MaxInt32
	sum, lft := 0, 0

	// Expand the right-side of the window.
	for rht := range nums {
		// Calculate the window sum.
		sum += nums[rht]

		// Contract the left-side of the window
		// while condition is true.
		for sum >= tgt {
			// Calculate the current length.
			curLen := rht - lft + 1

			// Determine min length.
			if curLen < minLen {
				minLen = curLen
			}

			// Contract the window.
			sum -= nums[lft]
			lft++
		}
	}

	// Check for no solution.
	if minLen == math.MaxInt32 {
		return 0
	}

	return minLen
}

func TestMinSubArrayLen209(t *testing.T) {
	tests := []struct {
		name     string
		tgt      int
		nums     []int
		expected int
	}{
		{
			name:     "Example 1",
			tgt:      7,
			nums:     []int{2, 3, 1, 2, 4, 3},
			expected: 2,
		},
		{
			name:     "Example 2",
			tgt:      4,
			nums:     []int{1, 4, 4},
			expected: 1,
		},
		{
			name:     "Example 3",
			tgt:      11,
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 1},
			expected: 0,
		},
		{
			name:     "Single element equal to target",
			tgt:      5,
			nums:     []int{5},
			expected: 1,
		},
		{
			name:     "Empty array",
			tgt:      5,
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "All elements sum less than target",
			tgt:      15,
			nums:     []int{1, 2, 3, 4},
			expected: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := minSubArrayLen209(tc.tgt, tc.nums)
			if got != tc.expected {
				t.Errorf("minSubArrayLen(%d, %v) = %d; want %d",
					tc.tgt, tc.nums, got, tc.expected)
			}
		})
	}
}
