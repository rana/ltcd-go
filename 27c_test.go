package main

import (
	"reflect"
	"testing"
)

// Time complexity: O(n), n is the length of array nums.
// Space complexity: O(1), constant additional space used.
func removeElement27c(nums []int, val int) int {
	// Remove Element
	// Given an integer array nums.
	// Given an integer val.
	// Remove val from nums.
	// Return the number of valid values in nums.
	// Conditions:
	// * Remove in-place.
	// * Move non-val to end of nums.
	// Use a two-pointer technique.
	// Move valid values from right to left.

	// Initialize a variable.
	lft := 0

	// Iterate through each integer in nums.
	for rht := 0; rht < len(nums); rht++ {
		// Check whether current value is valid.
		if nums[rht] != val {
			nums[lft] = nums[rht]
			lft++
		}
	}

	return lft
}

func TestRemoveElement27c(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		val     int
		expLen  int
		expNums []int
	}{
		{
			name:    "Basic case 1",
			nums:    []int{3, 2, 2, 3},
			val:     3,
			expLen:  2,
			expNums: []int{2, 2},
		},
		{
			name:    "Basic case 2",
			nums:    []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:     2,
			expLen:  5,
			expNums: []int{0, 1, 3, 0, 4},
		},
		{
			name:    "Empty array",
			nums:    []int{},
			val:     0,
			expLen:  0,
			expNums: []int{},
		},
		{
			name:    "No match",
			nums:    []int{1, 2, 3, 4},
			val:     5,
			expLen:  4,
			expNums: []int{1, 2, 3, 4},
		},
		{
			name:    "All match",
			nums:    []int{1, 1, 1, 1},
			val:     1,
			expLen:  0,
			expNums: []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy of input slice for error reporting
			cpyNums := make([]int, len(tt.nums))
			copy(cpyNums, tt.nums)

			gotLen := removeElement27c(tt.nums, tt.val)

			if gotLen != tt.expLen {
				t.Errorf("removeElement(%v, %d) = %d; want %d",
					cpyNums, tt.val, gotLen, tt.expLen)
			}

			if !reflect.DeepEqual(tt.nums[:gotLen], tt.expNums) {
				t.Errorf("After removeElement(%v, %d), nums = %v; want %v",
					cpyNums, tt.val, tt.nums[:gotLen], tt.expNums)
			}
		})
	}
}
