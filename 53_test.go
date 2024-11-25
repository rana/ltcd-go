package main

import "testing"

// Time complexity: O(n), we traverse each element of the array nums.
// Space complexity: O(1), constant additional space used.
// https://claude.ai/chat/ca699e85-6d6f-41a7-b584-3c2d4ef57fea
func maxSubArray53(nums []int) int {
	// Maximum Subarray
	// Given and integer array nums.
	// Find the max sum within subarrays.
	// Return the max sum.
	// Use Kadane's algorithm.

	// Initialize variables.
	runSum, maxSum := nums[0], nums[0]

	// Iterate through remaining numbers.
	for _, num := range nums[1:] {
		// Two choices:
		//	1. Choose to continue running sum (runSum+num)
		//	2. Choose to start new running sum (num)
		// Then select max of choice, and previous max.
		runSum = max(num, runSum+num)
		maxSum = max(maxSum, runSum)
	}

	return maxSum
}

// TestMaxSubArray53 runs test cases for maxSubArray function
func TestMaxSubArray53(t *testing.T) {
	cases := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1: Mixed positive and negative",
			nums:     []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected: 6,
		},
		{
			name:     "Example 2: Single element",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "Example 3: All positive numbers",
			nums:     []int{5, 4, -1, 7, 8},
			expected: 23,
		},
		{
			name:     "All negative numbers",
			nums:     []int{-2, -3, -1, -5},
			expected: -1,
		},
		{
			name:     "Zero and negative numbers",
			nums:     []int{-2, 0, -1},
			expected: 0,
		},
		{
			name:     "Single negative number",
			nums:     []int{-1},
			expected: -1,
		},
		{
			name:     "Alternating positive and negative",
			nums:     []int{1, -1, 1, -1, 1},
			expected: 1,
		},
		{
			name:     "Large positive sum in middle",
			nums:     []int{-1, -2, 10, 20, -3, -4},
			expected: 30,
		},
		{
			name:     "Maximum at end",
			nums:     []int{-2, -3, 4, 5},
			expected: 9,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := maxSubArray53(tc.nums)
			if got != tc.expected {
				t.Errorf("maxSubArray(%v) = %d; want %d", tc.nums, got, tc.expected)
			}
		})
	}
}
