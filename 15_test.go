package main

import (
	"fmt"
	"sort"
	"testing"
)

// Time complexity: O(n^2), n is the length of the nums array. We sort the array for O(nlogn). We traverse the array with an outer loop and inner loop for O(n^2).
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/14e3edf4-cd8c-413b-968c-1e205d50773a
func threeSum(nums []int) [][]int {
	// 3Sum
	// Given an integer array nums.
	// Find all triplets which sum to zero.
	// Return an array of triplets.
	// Notice that the array contains duplicates.
	// Notice that the array is unsorted.
	// Reduce the 3sum to 2sum with sorting.
	// Use a two-pointer technique with sorting.
	// Pin the first number and apply a two-pointer technique.

	// Initialize a result array.
	var res [][]int

	// Check input minimum edge case.
	if len(nums) < 3 {
		return res
	}

	// Sort nums to enable two-pointer technique.
	sort.Ints(nums)

	// Iterate through first numbers.
	for idx := 0; idx < len(nums)-2; idx++ {
		// Check for first number duplicates.
		if idx > 0 && nums[idx] == nums[idx-1] {
			continue
		}

		// Initialize two pointers.
		lft, rht := idx+1, len(nums)-1

		// Iterate until two pointers meet.
		for lft < rht {
			// Calculate the sum of the three numbers.
			sum := nums[idx] + nums[lft] + nums[rht]

			// Check for the success condition.
			const tgt = 0
			if sum == tgt {
				// Store a solution.
				res = append(res, []int{nums[idx], nums[lft], nums[rht]})

				// Skip second number duplicates.
				for lft < rht && nums[lft] == nums[lft+1] {
					lft++
				}

				// Skip third number duplicates.
				for lft < rht && nums[rht] == nums[rht-1] {
					rht--
				}

				// Move pointers.
				lft++
				rht--
			} else if sum < tgt {
				// Sum too small, make next larger.
				lft++
			} else {
				rht--
			}
		}
	}

	// Return solutions.
	return res
}

// TestThreeSum runs test cases for the threeSum function
func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{
			name: "Example 1",
			nums: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "Example 2",
			nums: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "Example 3",
			nums: []int{0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			name: "Empty array",
			nums: []int{},
			want: [][]int{},
		},
		{
			name: "Array with duplicates",
			nums: []int{-2, 0, 0, 2, 2},
			want: [][]int{{-2, 0, 2}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			if !compareResults(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

// compareResults compares two slices of integer slices regardless of order
func compareResults(got, want [][]int) bool {
	if len(got) != len(want) {
		return false
	}

	// Create maps for comparison
	gotMap := make(map[string]bool)
	wantMap := make(map[string]bool)

	for _, triplet := range got {
		sort.Ints(triplet)
		key := fmt.Sprintf("%d,%d,%d", triplet[0], triplet[1], triplet[2])
		gotMap[key] = true
	}

	for _, triplet := range want {
		sort.Ints(triplet)
		key := fmt.Sprintf("%d,%d,%d", triplet[0], triplet[1], triplet[2])
		wantMap[key] = true
	}

	for k := range gotMap {
		if !wantMap[k] {
			return false
		}
	}

	return true
}
