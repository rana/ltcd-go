package main

import (
	"reflect"
	"testing"
)

// Time complexity: O(n!), for computing all permutations.
// Space complexity: O(n), for current permutation, visited map, and stack depth.
// permute46 returns all possible permutations of given integers
// https://claude.ai/chat/2d11eb9a-1dc2-4d84-a827-999c58e73684
func permute46(nums []int) [][]int {
	// Permutations
	// Given an integer array nums.
	// Create all possible permutations.
	// Return the permutations.
	// Use backtracking and recursion.
	// Use a closure for recursion and state management.
	// Use a visited map.

	// Check input min edge case.
	numLen := len(nums)
	if numLen <= 1 {
		return [][]int{nums}
	}

	// Initialize variables.
	var res [][]int
	cur := make([]int, 0, numLen)
	vst := make(map[int]bool)

	// Create closure for backtracking.
	var bkt func()
	bkt = func() {
		// Check base case.
		// Check whether whole permutation created.
		if len(cur) == numLen {
			// Copy cur
			tmp := make([]int, numLen)
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		// Create permutations recursively.
		for _, num := range nums {
			// Check whether the number is already visited.
			if vst[num] {
				continue
			}

			// Use current number.
			vst[num] = true
			cur = append(cur, num)

			// Recursively create permutations.
			bkt()

			// Backtrack.
			cur = cur[:len(cur)-1]
			vst[num] = false
		}
	}

	// Start backtracking
	bkt()

	return res
}

// TestPermute46 validates the permute function
func TestPermute46(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{
			name:  "Example 1",
			input: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3}, {1, 3, 2},
				{2, 1, 3}, {2, 3, 1},
				{3, 1, 2}, {3, 2, 1},
			},
		},
		{
			name:     "Example 2",
			input:    []int{0, 1},
			expected: [][]int{{0, 1}, {1, 0}},
		},
		{
			name:     "Example 3",
			input:    []int{1},
			expected: [][]int{{1}},
		},
		{
			name:     "Empty array",
			input:    []int{},
			expected: [][]int{{}},
		},
		{
			name:     "Array with negative numbers",
			input:    []int{-1, 1},
			expected: [][]int{{-1, 1}, {1, -1}},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := permute46(tc.input)
			// Sort both got and expected for consistent comparison
			if len(got) != len(tc.expected) {
				t.Errorf("got length %v, want length %v", len(got), len(tc.expected))
				return
			}

			// Create map to check if all permutations are present
			expMap := make(map[string]bool)
			for _, perm := range tc.expected {
				expMap[arrayToString46(perm)] = true
			}

			// Verify each generated permutation is valid
			for _, perm := range got {
				if !expMap[arrayToString46(perm)] {
					t.Errorf("unexpected permutation %v", perm)
				}
			}
		})
	}
}

// arrayToString46 creates string representation of array for comparison
func arrayToString46(arr []int) string {
	return string(reflect.ValueOf(arr).String())
}
