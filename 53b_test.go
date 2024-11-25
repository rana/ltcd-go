package main

import "testing"

func maxSubArray53b(nums []int) int {
	// Maximum Subarray
	// Given an integer array nums.
	// Find the max sum of any subarray.
	// Return the max sum.
	// Use Kadane's algorithm.
	// Choice:
	//	* Continue current running sum.
	//	* Start new sum if it's greater.
	// Then select max of current max, and previous choice.

	// Initialize variables.
	curMax, glbMax := nums[0], nums[0]

	// Iterate through remaining numbers.
	for _, cur := range nums[1:] {
		curMax = max(cur, cur+curMax)
		glbMax = max(glbMax, curMax)
	}

	return glbMax
}

// TestMaxSubArray53b runs test cases for maxSubArray function
func TestMaxSubArray53b(t *testing.T) {
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
			got := maxSubArray53b(tc.nums)
			if got != tc.expected {
				t.Errorf("maxSubArray(%v) = %d; want %d", tc.nums, got, tc.expected)
			}
		})
	}
}
