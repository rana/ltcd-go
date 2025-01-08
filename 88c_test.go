package main

import (
	"reflect"
	"testing"
)

// Time complexity: O(m + n), m is the length of array nums1. n is the length of array nums2.
// Space complexity: O(1), constant additional space used.
func merge88c(nums1 []int, m int, nums2 []int, n int) {
	// Merge Sorted Array
	// Given two sort-ascending integer arrays nums1 and nums2.
	// Merge arrays.
	// Conditions:
	// * Merge in-place
	// * Maintain sort-ascending
	// * nums1 has sufficient length of both arrays
	// Use a three-pointer technique.

	// Initialize variables.
	idx1, idx2 := m-1, n-1
	idx := m + n - 1

	// Iterate through each array.
	for idx1 >= 0 && idx2 >= 0 {
		// Maintain sort-ascending.
		if nums1[idx1] > nums2[idx2] {
			nums1[idx] = nums1[idx1]
			idx1--
		} else {
			nums1[idx] = nums2[idx2]
			idx2--
		}
		idx--
	}

	// Append any remaining nums2.
	for idx2 >= 0 {
		nums1[idx] = nums2[idx2]
		idx2--
		idx--
	}
}

func TestMerge88c(t *testing.T) {
	tests := []struct {
		name   string
		nums1  []int
		m      int
		nums2  []int
		n      int
		expect []int
	}{
		{
			name:   "both arrays non-empty",
			nums1:  []int{1, 2, 3, 0, 0, 0},
			m:      3,
			nums2:  []int{2, 5, 6},
			n:      3,
			expect: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:   "nums2 empty",
			nums1:  []int{1},
			m:      1,
			nums2:  []int{},
			n:      0,
			expect: []int{1},
		},
		{
			name:   "nums1 empty",
			nums1:  []int{0},
			m:      0,
			nums2:  []int{1},
			n:      1,
			expect: []int{1},
		},
		{
			name:   "duplicate elements",
			nums1:  []int{2, 2, 3, 0, 0, 0},
			m:      3,
			nums2:  []int{1, 2, 3},
			n:      3,
			expect: []int{1, 2, 2, 2, 3, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge88c(tt.nums1, tt.m, tt.nums2, tt.n)
			if !reflect.DeepEqual(tt.nums1, tt.expect) {
				t.Errorf("Expected %v, got %v", tt.expect, tt.nums1)
			}
		})
	}
}
