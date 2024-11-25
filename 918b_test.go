package main

import "testing"

func maxSubarraySumCircular918b(nums []int) int {
	// Maximum Sum Circular Subarray
	// Given a circular integer array nums.
	// Find the max sum of any subarray.
	// Use Kadane's algorithm, and modified algorithm.
	// Track Kadane's algorithm for non-circular part.
	// Track modified Kadane's algorithm for circular part.
	// 	total, glbMin with tot-glbMin
	// Get max of the two approaches.
	// Check special case when all negative; return glbMax.

	// Initialize variables.
	tot := nums[0]
	curMax, glbMax := nums[0], nums[0]
	curMin, glbMin := nums[0], nums[0]

	// Iterate through remaining numbers.
	for _, cur := range nums[1:] {
		tot += cur

		// Kadane's for linear approach.
		curMax = max(cur, cur+curMax)
		glbMax = max(glbMax, curMax)

		// Kadane's modified for circular approach.
		curMin = min(cur, cur+curMin)
		glbMin = min(glbMin, curMin)
	}

	// Check special case of all negative.
	if glbMax < 0 {
		return glbMax
	}

	// Return max of both approaches.
	return max(glbMax, tot-glbMin)
}

func TestMaxSubarraySumCircular918b(t *testing.T) {
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
			got := maxSubarraySumCircular918b(tc.nums)
			if got != tc.expected {
				t.Errorf("maxSubarraySumCircular(%v) = %v; want %v", tc.nums, got, tc.expected)
			}
		})
	}
}
