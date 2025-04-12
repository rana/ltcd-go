package main

import (
	"testing"
)

// Split Array Equal Sum
// https://www.geeksforgeeks.org/split-array-two-equal-sum-subarrays/
//
// Split array into two subarrays with equal sum.
//
// Given an array of integers greater than zero, find if it is possible to split it in two subarrays (without reordering the elements), such that the sum of the two subarrays is the same. Return the split index; or -1.
//
// - Time Complexity: O(n), where n is the length of the input array. We traverse the array twice - once to calculate the total sum and once to find the split point.
// - Space Complexity: O(1), as we only use a constant amount of extra space regardless of input size.
//
// https://console.anthropic.com/workbench/12b735bf-4ea4-4c73-b769-b2b1ca5626e0
func splitArrayEqualSum1(nums []int) int {

	// Calculate total sum
	totSum := 0
	for _, num := range nums {
		totSum += num
	}

	// Iterate through array elements
	lftSum := 0
	for idx, num := range nums {
		lftSum += num
		rhtSum := totSum - lftSum

		// If left sum equals right sum, we found the split point
		if lftSum == rhtSum {
			return idx
		}
	}

	// No split point found
	return -1
}

func TestSplitArrayEqualSum1(t *testing.T) {
	cases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Example 1",
			input:    []int{1, 2, 3, 4, 5, 5},
			expected: 3,
		},
		{
			name:     "Example 2",
			input:    []int{4, 1, 2, 3},
			expected: 1,
		},
		{
			name:     "Example 3",
			input:    []int{4, 3, 2, 1},
			expected: -1,
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: -1,
		},
		{
			name:     "Single element",
			input:    []int{5},
			expected: -1,
		},
		{
			name:     "Two elements with equal sum",
			input:    []int{5, 5},
			expected: 0,
		},
		{
			name:     "All zeros",
			input:    []int{0, 0, 0, 0},
			expected: 0, // Any index works
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := splitArrayEqualSum1(tc.input)
			if got != tc.expected {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}
