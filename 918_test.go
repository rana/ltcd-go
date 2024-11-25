package main

import "testing"

// Time complexity: O(n), we traverse each element of the array nums.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/76be3487-5f46-4f68-a997-4ebedb0ef340
func maxSubarraySumCircular918(nums []int) int {
	// Maximum Sum Circular Subarray
	// Given an integer array nums.
	// Find the max sum of subarrays.
	// Conditions:
	//	* Consider the array circular
	// Use Kadane's algorithm for non-circular.
	// Use total sum - min sum for circular.
	// Select max of the two approaches.
	// Check case of all negative, return global max.

	// Initialize variables.
	tot := nums[0]
	curMax, curMin := nums[0], nums[0]
	glbMax, glbMin := nums[0], nums[0]

	// Iterate through each remaining number.
	for _, cur := range nums[1:] {
		tot += cur

		// Update non-circular Kadane's algorithm.
		curMax = max(cur, cur+curMax)
		glbMax = max(glbMax, curMax)

		// Update circular Kadane's algorithm variant.
		curMin = min(cur, cur+curMin)
		glbMin = min(glbMin, curMin)
	}

	// Check all negative case.
	if glbMax < 0 {
		return glbMax
	}

	// Return max of approaches.
	return max(glbMax, tot-glbMin)
}

func TestMaxSubarraySumCircular918(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1: Simple positive max",
			nums:     []int{1, -2, 3, -2},
			expected: 3,
		},
		{
			name:     "Example 2: Circular sum",
			nums:     []int{5, -3, 5},
			expected: 10,
		},
		{
			name:     "Example 3: All negative",
			nums:     []int{-3, -2, -3},
			expected: -2,
		},
		{
			name:     "Single element",
			nums:     []int{5},
			expected: 5,
		},
		{
			name:     "All positive",
			nums:     []int{1, 2, 3, 4},
			expected: 10,
		},
		{
			name:     "Zero sum possible",
			nums:     []int{-5, 3, -1, 2, -4, 5},
			expected: 5,
		},
		{
			name:     "Large negative numbers",
			nums:     []int{-30000, -20000, -10000},
			expected: -10000,
		},
		{
			name:     "Alternating signs",
			nums:     []int{1, -1, 2, -2, 3, -3},
			expected: 3,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := maxSubarraySumCircular918(tc.nums)
			if got != tc.expected {
				t.Errorf("maxSubarraySumCircular(%v) = %v; want %v", tc.nums, got, tc.expected)
			}
		})
	}
}
