package main

import "testing"

// Time complexity: O(n)
// Space complexity: O(n)
// https://console.anthropic.com/workbench/4afe9b4b-2acd-44a5-bf63-5a8f0ba83b44
func maxSubArrayLen325(nums []int, k int) int {
	// Map to store prefix sum and its earliest occurrence index
	sumIdx := make(map[int]int)
	sumIdx[0] = -1 // Initialize with 0 sum at index -1

	maxLen := 0
	runSum := 0

	for idx, num := range nums {
		runSum += num

		// Check if there's a prefix sum that, when subtracted from current sum, gives k
		// runSum - prevSum = k  ->  prevSum = runSum - k
		if prvIdx, exists := sumIdx[runSum-k]; exists {
			curLen := idx - prvIdx
			if curLen > maxLen {
				maxLen = curLen
			}
		}

		// Only store the earliest occurrence of a sum
		if _, exists := sumIdx[runSum]; !exists {
			sumIdx[runSum] = idx
		}
	}

	return maxLen
}

func TestMaxSubArrayLen(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, -1, 5, -2, 3},
			k:        3,
			expected: 4,
		},
		{
			name:     "Example 2",
			nums:     []int{-2, -1, 2, 1},
			k:        1,
			expected: 2,
		},
		{
			name:     "Single element equal to k",
			nums:     []int{3},
			k:        3,
			expected: 1,
		},
		{
			name:     "No subarray sums to k",
			nums:     []int{1, 2, 3},
			k:        7,
			expected: 0,
		},
		{
			name:     "Zero sum subarray",
			nums:     []int{1, -1, 0},
			k:        0,
			expected: 2,
		},
		{
			name:     "Entire array sums to k",
			nums:     []int{1, 2, 3, 4},
			k:        10,
			expected: 4,
		},
		{
			name:     "Negative numbers and negative k",
			nums:     []int{-5, 8, -14, 2, 4, 12},
			k:        -5,
			expected: 5,
		},
		{
			name:     "Large array with repetitive patterns",
			nums:     []int{1, 1, 1, 1, 1, -1, -1, -1, -1, -1},
			k:        0,
			expected: 10,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := maxSubArrayLen325(tc.nums, tc.k)
			if got != tc.expected {
				t.Errorf("maxSubArrayLen(%v, %d) = %d; want %d",
					tc.nums, tc.k, got, tc.expected)
			}
		})
	}
}
