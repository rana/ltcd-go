package main

import "testing"

// Time complexity: O(log n), n is the length of the array. We performa a binary search for O(log n).
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/2a7c581e-75c1-459f-b34b-12bdd230227a
func findPeakElement162(nums []int) int {
	// Find Peak Element
	// Given an integer array nums.
	// Find a peak element.
	// Return index of a peak element.
	// Conditions:
	//	* Peak element is great than element on right and left.
	//	* Time complexity: O(log n)
	// Use binary search.

	// Check input min case.
	if len(nums) == 1 {
		return 0
	}

	// Initialize binary search bounds.
	lft, rht := 0, len(nums)-1

	// Binary search for peak element.
	for lft < rht {
		// Calculate middle index.
		mid := lft + (rht-lft)/2

		// Check whether mid is less than next.
		if nums[mid] < nums[mid+1] {
			// Peak on right side.
			lft = mid + 1
		} else {
			// Peak on left side.
			rht = mid
		}
	}

	return lft
}

// Tests
func TestFindPeakElement162(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int // Note: Some inputs may have multiple valid answers
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 1},
			want: 2,
		},
		{
			name: "Example 2 - multiple peaks",
			nums: []int{1, 2, 1, 3, 5, 6, 4},
			want: 5,
		},
		{
			name: "Single element",
			nums: []int{1},
			want: 0,
		},
		{
			name: "Two elements ascending",
			nums: []int{1, 2},
			want: 1,
		},
		{
			name: "Two elements descending",
			nums: []int{2, 1},
			want: 0,
		},
		{
			name: "Valley",
			nums: []int{3, 1, 2},
			want: 0, // or 2
		},
		{
			name: "All ascending",
			nums: []int{1, 2, 3, 4, 5},
			want: 4,
		},
		{
			name: "All descending",
			nums: []int{5, 4, 3, 2, 1},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findPeakElement162(tt.nums)
			// Special handling for cases with multiple valid answers
			if tt.name == "Example 2 - multiple peaks" {
				if got != 1 && got != 5 { // both 1 and 5 are valid peaks
					t.Errorf("findPeakElement() = %v, want either 1 or 5", got)
				}
				return
			}
			if tt.name == "Valley" {
				if got != 0 && got != 2 { // both 0 and 2 are valid peaks
					t.Errorf("findPeakElement() = %v, want either 0 or 2", got)
				}
				return
			}
			// For other cases, check exact match
			if got != tt.want {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
			// Verify the result is actually a peak
			if len(tt.nums) > 1 {
				idx := got
				if idx > 0 && tt.nums[idx] <= tt.nums[idx-1] {
					t.Errorf("Result %v is not greater than left neighbor", idx)
				}
				if idx < len(tt.nums)-1 && tt.nums[idx] <= tt.nums[idx+1] {
					t.Errorf("Result %v is not greater than right neighbor", idx)
				}
			}
		})
	}
}
