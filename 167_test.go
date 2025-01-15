package main

import (
	"reflect"
	"testing"
)

// Time complexity: O(n), n is the length of the nums array.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/fe698065-3a01-42f6-874d-f655b19e3842
func twoSum167(nums []int, tgt int) []int {
	// Two Sum II - Input Array Is Sorted
	// Given a sort-ascending integer array `nums`.
	// Given an integer `tgt`.
	// Find two numbers in `nums` which equal `tgt`.
	// Return an array of indexes of the numbers in 1-based form.
	// Conditions:
	//	* Solve with O(1) time complexity.
	//	* One solution guaranteed.
	// Use a two-pointer technique which uses the sort ascending.

	// Initialize two pointers.
	lft, rht := 0, len(nums)-1

	// Iterate until the two pointers meet.
	for lft < rht {
		// Calculate the sum of the two current numbers.
		sum := nums[lft] + nums[rht]

		// Check for success condition.
		if sum == tgt {
			// Success condition found.
			// Return 1-based indexes of numbers.
			return []int{lft + 1, rht + 1}
		} else if sum < tgt {
			// Sum too small. Make next larger.
			lft++
		} else {
			// Sum too larger. Make next smaller.
			rht--
		}
	}

	// Unexpected.
	return []int{0, 0}
}

// Test cases
func TestTwoSum167(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Example 1: Standard case",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{1, 2},
		},
		{
			name:     "Example 2: Three numbers",
			nums:     []int{2, 3, 4},
			target:   6,
			expected: []int{1, 3},
		},
		{
			name:     "Example 3: Negative numbers",
			nums:     []int{-1, 0},
			target:   -1,
			expected: []int{1, 2},
		},
		{
			name:     "Test 4: Larger numbers",
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			target:   13,
			expected: []int{6, 7},
		},
		{
			name:     "Test 5: Adjacent numbers",
			nums:     []int{1, 2, 3, 4},
			target:   3,
			expected: []int{1, 2},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := twoSum167(tc.nums, tc.target)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("twoSum() = %v, want %v", got, tc.expected)
			}
		})
	}
}
