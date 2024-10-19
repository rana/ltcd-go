package main

import (
	"reflect"
	"testing"
)

// Time complexity: O(n), n is the length of the array nums. We traverse the array once.
// Space complexity: O(1), constant additional space used.
// https://chatgpt.com/c/6714151c-8950-8002-a100-a0b27371f372
func removeDuplicates2(nums []int) int {
	// Given a sort-ascending integer array nums.
	// Remove more than two duplicates in-place.
	// Maintain sort-ascending.
	// Use a two-pointer technique.

	// Define the unique count.
	const unq_cnt = 2

	// Check for input minimum edge case.
	if len(nums) <= unq_cnt {
		return len(nums)
	}

	// lft is the index of the next unique condition.
	lft := unq_cnt

	// Iterate through remaining elements of nums.
	for rht := unq_cnt; rht < len(nums); rht++ {
		// Check for unique condition met.
		// Skip elements not meeting condition.
		if nums[rht] != nums[lft-unq_cnt] {
			nums[lft] = nums[rht]
			lft++
		}
	}

	return lft
}

// TestRemoveDuplicates tests the removeDuplicates function with various test cases.
func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		wantLen int
		wantRes []int
	}{
		{
			name:    "Example 1",
			nums:    []int{1, 1, 1, 2, 2, 3},
			wantLen: 5,
			wantRes: []int{1, 1, 2, 2, 3},
		},
		{
			name:    "Example 2",
			nums:    []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			wantLen: 7,
			wantRes: []int{0, 0, 1, 1, 2, 3, 3},
		},
		{
			name:    "No duplicates",
			nums:    []int{1, 2, 3, 4, 5},
			wantLen: 5,
			wantRes: []int{1, 2, 3, 4, 5},
		},
		{
			name:    "All duplicates",
			nums:    []int{1, 1, 1, 1, 1},
			wantLen: 2,
			wantRes: []int{1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := make([]int, len(tt.nums))
			copy(numsCopy, tt.nums)
			gotLen := removeDuplicates2(numsCopy)
			if gotLen != tt.wantLen {
				t.Errorf("got length %d, want %d", gotLen, tt.wantLen)
			}
			if !reflect.DeepEqual(numsCopy[:gotLen], tt.wantRes) {
				t.Errorf("got nums %v, want %v", numsCopy[:gotLen], tt.wantRes)
			}
		})
	}
}
