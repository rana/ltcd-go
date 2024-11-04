package main

import (
	"reflect"
	"testing"
)

// Time complexity: O(n), n is the length of the nums array. We iterate through the array once.
// Space complexity: O(n), up to n number-index pairs stored in a map.
func twoSum(nums []int, tgt int) []int {
	// Two Sum
	// Given an integer array nums.
	// Given an integer tgt.
	// Find two numbers in nums which sum to tgt.
	// Return the number indexes.
	// Store a number to index mapping for seen values.
	// Check with the complement of the current number.

	// Initialize a complement map.
	cmp := make(map[int]int)

	// Iterate through each number.
	for idx, num := range nums {
		// Check whether complement already encountered.
		if prvIdx, exists := cmp[tgt-num]; exists {
			return []int{prvIdx, idx}
		}

		// Store current number -> index mapping.
		cmp[num] = idx
	}

	// No solution.
	return nil
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		tgt  int
		exp  []int
	}{
		{
			name: "basic case",
			nums: []int{2, 7, 11, 15},
			tgt:  9,
			exp:  []int{0, 1},
		},
		{
			name: "numbers not at start",
			nums: []int{3, 2, 4},
			tgt:  6,
			exp:  []int{1, 2},
		},
		{
			name: "same number twice",
			nums: []int{3, 3},
			tgt:  6,
			exp:  []int{0, 1},
		},
		{
			name: "negative numbers",
			nums: []int{-1, -2, -3, -4, -5},
			tgt:  -8,
			exp:  []int{2, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := twoSum(tc.nums, tc.tgt)
			if !reflect.DeepEqual(res, tc.exp) {
				t.Errorf("twoSum(%v, %d) = %v; want %v",
					tc.nums, tc.tgt, res, tc.exp)
			}
		})
	}
}
