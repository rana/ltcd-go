package main

import "testing"

func longestConsecutive(nums []int) int {
	// Longest Consecutive Sequence
	// Given an unsorted integer array nums.
	// Determine the longest length of consecutive integers.
	// Return the longest length.
	// Conditions:
	//	* Time complexity: O(n)
	// Use a hash set to achieve O(n) time.
	// Search for sequence starts.

	// Initialize a hash set.
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	// Initialize max length.
	maxLen := 0

	// Iterate through nums looking for sequence starts.
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

		// Check whether to update max length.
		if curLen > maxLen {
			maxLen = curLen
		}
	}

	return maxLen
}

func TestLongestConsecutive(t *testing.T) {
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
			got := longestConsecutive(tc.nums)
			if got != tc.expected {
				t.Errorf("longestConsecutive(%v) = %v; want %v",
					tc.nums, got, tc.expected)
			}
		})
	}
}
