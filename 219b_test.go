package main

import "testing"

// Time complexity: O(n), n is the length of the nums array. We traverse the array once.
// Space complexity: O(min(n, k)), up to k num-idx pairs are stored in a map.
func containsNearbyDuplicate(nums []int, k int) bool {
	// Contains Duplicate II
	// Given an integer array nums.
	// Given an integer k.
	// Determine whether a duplicate exists within
	// a window with length k.
	// Notice that the array is not sorted.
	// Use a num-idx map for window k.

	// Initialize num-idx map.
	numIdx := make(map[int]int)

	// Iterate through each number in nums.
	for idx, num := range nums {
		// Check whether we've seen this integer before.
		if prvIdx, exists := numIdx[num]; exists {
			// Check if window condition met.
			if idx-prvIdx <= k {
				return true
			}
		}

		// Store num-idx mapping
		numIdx[num] = idx
	}

	return false
}

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{
			name: "example 1: duplicate within range",
			nums: []int{1, 2, 3, 1},
			k:    3,
			want: true,
		},
		{
			name: "example 2: adjacent duplicates",
			nums: []int{1, 0, 1, 1},
			k:    1,
			want: true,
		},
		{
			name: "example 3: duplicates outside range",
			nums: []int{1, 2, 3, 1, 2, 3},
			k:    2,
			want: false,
		},
		{
			name: "single element",
			nums: []int{1},
			k:    1,
			want: false,
		},
		{
			name: "no duplicates",
			nums: []int{1, 2, 3, 4},
			k:    1,
			want: false,
		},
		{
			name: "k equals zero",
			nums: []int{1, 1, 1},
			k:    0,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.nums, tt.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
