package main

import (
	"reflect"
	"testing"
)

// Time complexity: O(n), n is the length of the nums array. We traverse the array once.
// Space complexity: O(n), we store up to n numbers in a map.
// https://claude.ai/chat/e71b28e8-24ae-48a1-bfea-0bfc68949103
func twoSum1(nums []int, tgt int) []int {
	// Two Sum
	// Given an integer array nums.
	// Given an integer tgt.
	// Find two numbers which sum to tgt.
	// Return the indexes of the two numbers.
	// Use a map to track previously seen number -> index.

	// Initialize number->index map.
	cmp := make(map[int]int)

	// Iterate through each number.
	for idx, num := range nums {
		// Check whether the complement has already been seen.
		if prvIdx, ok := cmp[tgt-num]; ok {
			return []int{prvIdx, idx}
		}

		// Store current number and index.
		cmp[num] = idx
	}

	// No solution.
	return nil
}

// Unit Tests
func TestTwoSum1(t *testing.T) {
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
			res := twoSum1(tc.nums, tc.tgt)
			if !reflect.DeepEqual(res, tc.exp) {
				t.Errorf("twoSum(%v, %d) = %v; want %v",
					tc.nums, tc.tgt, res, tc.exp)
			}
		})
	}
}
