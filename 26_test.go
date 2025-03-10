package main

import (
	"reflect"
	"testing"
)

// Time complexity: O(n), n is the length of the nums array. We traverse the array once.
// Space complexity: O(1), constant additional space used.
// https://chatgpt.com/c/6714082c-4dc4-8002-9e24-3f3622ee11e4
func removeDuplicates26(nums []int) int {
	// Remove Duplicates from Sorted Array
	// Given an integer array nums.
	// Sort array in-place with unique count constraint.
	// Determine processed list length.
	// Return list length.
	// Conditions:
	// * 1 unique element
	// * Maintain sort-ascending
	// Use two-pointer technique.
	// Use unique count 1.

	// Initialize the unique count.
	const unq_cnt = 1

	// Check input minimum edge case.
	if len(nums) <= unq_cnt {
		return len(nums)
	}

	// lft is the index of the next unique element.
	lft := unq_cnt

	// Iterate through the remaining elements.
	for rht := unq_cnt; rht < len(nums); rht++ {
		// Check whether unique condition met.
		// Skip over elements not meeting condition.
		if nums[rht] != nums[lft-unq_cnt] {
			// Move right element to the left.
			nums[lft] = nums[rht]
			lft++
		}
	}

	return lft
}

func TestRemoveDuplicates26(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
		k        int
	}{
		// Valid test cases within constraints
		{nums: []int{1, 1, 2}, expected: []int{1, 2}, k: 2},
		{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, expected: []int{0, 1, 2, 3, 4}, k: 5},
		{nums: []int{-100, -100, -99, -99, 0, 0, 1, 2, 2}, expected: []int{-100, -99, 0, 1, 2}, k: 5},
		{nums: []int{1}, expected: []int{1}, k: 1},
	}

	for _, test := range tests {
		numsCopy := make([]int, len(test.nums))
		copy(numsCopy, test.nums) // Preserve original nums for error messages

		k := removeDuplicates26(test.nums)
		if k != test.k {
			t.Errorf("removeDuplicates(%v) returned k=%d; expected k=%d", numsCopy, k, test.k)
		}
		if !reflect.DeepEqual(test.nums[:k], test.expected) {
			t.Errorf("After removeDuplicates(%v), nums=%v; expected nums=%v", numsCopy, test.nums[:k], test.expected)
		}
	}
}
