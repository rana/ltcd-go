package main

import (
	"math"
	"testing"
)

// Time Complexity: O(n) where n is the length of the array. We iterate through the array once to calculate the initial sum, and then once more to slide the window.
// Space Complexity: O(1) as we only use a few variables regardless of the input size.
// https://console.anthropic.com/workbench/8bf025be-8a86-419d-b352-82c20987ae69
func findMaxAverage643(nums []int, k int) float64 {
	// Maximum Average Subarray I
	// Given an integer array nums.
	// Given an integer k.
	// Determine a max average for
	// a subarray of length k.
	// Return the max average.
	// Track the sum.

	// Calculate initial curSum of first k elements
	curSum := 0
	for i := range k {
		curSum += nums[i]
	}

	// Initialize maxSum with the first window's sum
	maxSum := curSum

	// Slide the window and update maxSum if needed
	for i := k; i < len(nums); i++ {
		// Add the new element and remove the element that's no longer in the window
		curSum = curSum + nums[i] - nums[i-k]
		if curSum > maxSum {
			maxSum = curSum
		}
	}

	// Return the maximum average
	return float64(maxSum) / float64(k)
}

func TestFindMaxAverage643(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected float64
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 12, -5, -6, 50, 3},
			k:        4,
			expected: 12.75,
		},
		{
			name:     "Example 2",
			nums:     []int{5},
			k:        1,
			expected: 5.0,
		},
		{
			name:     "All negative numbers",
			nums:     []int{-1, -2, -3, -4, -5},
			k:        3,
			expected: -2.0,
		},
		{
			name:     "Exactly k elements",
			nums:     []int{1, 2, 3, 4},
			k:        4,
			expected: 2.5,
		},
		{
			name:     "Maximum at the beginning",
			nums:     []int{10, 10, 10, 1, 1, 1},
			k:        3,
			expected: 10.0,
		},
		{
			name:     "Maximum at the end",
			nums:     []int{1, 1, 1, 10, 10, 10},
			k:        3,
			expected: 10.0,
		},
		{
			name:     "Larger numbers",
			nums:     []int{-10000, 10000, -10000, 10000},
			k:        2,
			expected: 0.0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := findMaxAverage643(test.nums, test.k)
			if math.Abs(result-test.expected) > 1e-5 {
				t.Errorf("got %v, want %v", result, test.expected)
			}
		})
	}
}
