package main

import (
	"math"
	"testing"
)

// Time complexity: O(n) where n is the array length
// Space complexity: O(k) where k is the number of distinct values in a window (at most 3 for this problem)
// https://console.anthropic.com/workbench/f47554b3-a7d4-4ce3-ab85-1139244f4ebd
func continuousSubarrays2762(nums []int) int64 {
	// Continuous Subarrays
	// Given an integer array nums.
	// Determine count of continuous subarrays.
	// Return count.
	// Conditions:
	// * 0 <= abs(nums[i1] - nums[i2]) <= 2
	// * Subarray is contiguous, non-empty.
	// * Subarray may be single element.
	// Use two-pointer sliding window technique.
	// Use a frequency counter in window.

	cnt, lft := int64(0), 0

	// Keep track of frequency of each element in frqWnd
	frqWnd := make(map[int]int)

	for rht, num := range nums {
		// Add current element to window
		frqWnd[num]++

		// Check and fix window validity
		for !isValid2762(frqWnd) {
			// Remove leftmost element
			frqWnd[nums[lft]]--
			if frqWnd[nums[lft]] == 0 {
				delete(frqWnd, nums[lft])
			}
			lft++
		}

		// Add count of valid subarrays ending at right
		cnt += int64(rht - lft + 1)
	}

	return cnt
}

// Helper function to check if max-min ≤ 2
func isValid2762(frqWnd map[int]int) bool {
	min, max := math.MaxInt32, math.MinInt32
	for num := range frqWnd {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return max-min <= 2
}

func TestContinuousSubarrays2762(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int64
	}{
		// Standard examples from problem statement
		{
			name:     "Example 1: [5,4,2,4]",
			nums:     []int{5, 4, 2, 4},
			expected: 8,
		},
		{
			name:     "Example 2: [1,2,3]",
			nums:     []int{1, 2, 3},
			expected: 6,
		},

		// Edge cases
		{
			name:     "Single element",
			nums:     []int{7},
			expected: 1,
		},
		{
			name:     "Two identical elements",
			nums:     []int{3, 3},
			expected: 3, // [3], [3], [3,3]
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
		},

		// Array with all identical elements
		{
			name:     "All same elements",
			nums:     []int{5, 5, 5, 5},
			expected: 10, // n*(n+1)/2 = 4*5/2 = 10
		},

		// Special cases with max difference = 2
		{
			name:     "Max difference exactly 2",
			nums:     []int{10, 12, 10, 12, 10},
			expected: 15, // All possible subarrays are valid
		},
		{
			name:     "Three distinct values with diff ≤ 2",
			nums:     []int{1, 2, 3, 2, 1},
			expected: 15, // All possible subarrays are valid
		},

		// Cases where window needs to shrink
		{
			name:     "Requires window shrinking",
			nums:     []int{1, 5, 4, 2, 4},
			expected: 9, // First two elements form discontinuous subarray
		},
		{
			name:     "Alternating pattern forcing small windows",
			nums:     []int{1, 5, 1, 5, 1},
			expected: 5, // Only single-element subarrays are valid
		},

		// Increasing/decreasing sequences
		{
			name:     "Increasing sequence",
			nums:     []int{1, 3, 5, 7, 9},
			expected: 9, // Only windows of size 1 and 2 are valid
		},
		{
			name:     "Decreasing sequence",
			nums:     []int{9, 7, 5, 3, 1},
			expected: 9, // Only windows of size 1 and 2 are valid
		},

		// Repeated patterns
		{
			name:     "Repeated pattern within range",
			nums:     []int{1, 2, 3, 1, 2, 3},
			expected: 21, // All subarrays are valid
		},

		// Large value differences
		{
			name:     "Large jumps between values",
			nums:     []int{100, 1000, 10000, 100000},
			expected: 4, // Only individual elements form valid subarrays
		},

		// Special patterns
		{
			name:     "Plateau pattern",
			nums:     []int{1, 3, 3, 3, 1},
			expected: 15, // All possible subarrays are valid
		},
		{
			name:     "Valley pattern",
			nums:     []int{5, 3, 3, 3, 5},
			expected: 15, // All possible subarrays are valid
		},

		// Near boundary cases
		{
			name:     "At boundary",
			nums:     []int{1, 3, 5, 7},
			expected: 7, // Only pairs and individual elements
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := continuousSubarrays2762(tc.nums)
			if got != tc.expected {
				t.Errorf("continuousSubarrays(%v) = %v, want %v", tc.nums, got, tc.expected)
			}
		})
	}
}
