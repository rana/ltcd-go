package main

import (
	"reflect"
	"testing"
)

func removeDuplicates26c(nums []int) int {
	// Remove Duplicate from Sorted Array
	// Given a sort-ascending integer array.
	// Remove duplicate elements.
	// Conditions:
	// * Remove duplicate elements in-place.
	// * Place duplicates at end of list
	// * One unique element for each value.
	// Return count of valid elements.
	// Use a two-pointer technique.
	// Move valid elements from right to left.

	// Define a unique element count.
	const unq_cnt = 1

	// Check min edge case.
	if len(nums) <= unq_cnt {
		return len(nums)
	}

	// Initialize variable.
	lft := unq_cnt

	// Iterate through nums.
	for rht := unq_cnt; rht < len(nums); rht++ {
		// Check for unique element after gap.
		if nums[rht] != nums[lft-unq_cnt] {
			nums[lft] = nums[rht]
			lft++
		}
	}

	return lft
}

func TestRemoveDuplicates26c(t *testing.T) {
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

		k := removeDuplicates26c(test.nums)
		if k != test.k {
			t.Errorf("removeDuplicates(%v) returned k=%d; expected k=%d", numsCopy, k, test.k)
		}
		if !reflect.DeepEqual(test.nums[:k], test.expected) {
			t.Errorf("After removeDuplicates(%v), nums=%v; expected nums=%v", numsCopy, test.nums[:k], test.expected)
		}
	}
}
