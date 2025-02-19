package main

import (
	"reflect"
	"testing"
)

// https://chatgpt.com/c/671433e1-611c-8002-b721-a9ab2468fb0e
func rotate189(nums []int, k int) {
	// Rotate Array
	// Given an integer array nums.
	// Given a segment length k.
	// Rotate the array to the right by k.
	// The array wraps around.
	// Use three calls to a reverse function.
	// Reverse whole array, first segment, second segment.

	// Adjust k in case larger than nums.
	n := len(nums)
	k = k % n

	// Reverse whole array.
	reverse189(nums, 0, n-1)

	// Reverse first segment
	reverse189(nums, 0, k-1)

	// Reverse second segment
	reverse189(nums, k, n-1)
}

func reverse189(nums []int, lft, rht int) {
	// Reverse from left to right.
	for lft < rht {
		nums[lft], nums[rht] = nums[rht], nums[lft]
		lft++
		rht--
	}
}

func TestRotate189(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:     "Example 2",
			nums:     []int{-1, -100, 3, 99},
			k:        2,
			expected: []int{3, 99, -1, -100},
		},
		{
			name:     "K is zero",
			nums:     []int{1, 2, 3, 4, 5},
			k:        0,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "K equals array length",
			nums:     []int{1, 2, 3, 4, 5},
			k:        5,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "K larger than array length",
			nums:     []int{1, 2, 3, 4, 5},
			k:        12,
			expected: []int{4, 5, 1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate189(tt.nums, tt.k)
			if !reflect.DeepEqual(tt.nums, tt.expected) {
				t.Errorf("got %v, want %v", tt.nums, tt.expected)
			}
		})
	}
}
