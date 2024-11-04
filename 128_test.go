package main

import "testing"

// Time complexity: O(n), n is the length of the nums array. We traverse it twice. Once to create the hash set. A second time to look for sequences.
// Space complexity: O(n), each element is stored in a map.
// https://claude.ai/chat/6a23ec0f-d3dc-48dc-8eea-8562fff96f52
func longestConsecutive128(nums []int) int {
	// Longest Consecutive Sequence
	// Given an unsorted integer array nums.
	// Determine the longest length of consecutive integers.
	// Return the longest length.
	// Conditions:
	//	* Time complexity: O(n)
	// Use a hash set to store each integer.
	// Search for sequence starts.

	// Check for input minimum edge case.
	if len(nums) == 0 {
		return 0
	}

	// Create a hash set for each integer.
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	// Initialize the max length.
	maxLen := 0

	// Iterate through each integer looking for sequence starts.
	for _, num := range nums {
		// Skip non-sequence starts.
		if numSet[num-1] {
			continue
		}

		// Found a sequence start.
		curLen := 1
		curNum := num + 1

		// Measure the length of the sequence.
		for numSet[curNum] {
			curLen++
			curNum++
		}

		// Check whether to update the maximum length.
		if curLen > maxLen {
			maxLen = curLen
		}
	}

	return maxLen
}

// Test cases
func TestLongestConsecutive128(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Example 1: Basic sequence",
			nums:     []int{100, 4, 200, 1, 3, 2},
			expected: 4,
		},
		{
			name:     "Example 2: Longer sequence with duplicates",
			nums:     []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			expected: 9,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			expected: 0,
		},
		{
			name:     "Single element",
			nums:     []int{1},
			expected: 1,
		},
		{
			name:     "All same numbers",
			nums:     []int{1, 1, 1},
			expected: 1,
		},
		{
			name:     "No consecutive numbers",
			nums:     []int{2, 4, 6, 8},
			expected: 1,
		},
		{
			name:     "Negative numbers",
			nums:     []int{-3, -2, -1, 0, 1},
			expected: 5,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := longestConsecutive128(tc.nums)
			if got != tc.expected {
				t.Errorf("longestConsecutive(%v) = %v; want %v",
					tc.nums, got, tc.expected)
			}
		})
	}
}
