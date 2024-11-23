package main

import (
	"reflect"
	"testing"
)

// Time complexity: O(n!), n is the length of array nums.
// Space complexity: O(n), for recursive stack.
func permute46b(nums []int) [][]int {
	// Permutations
	// Given a distinct integer array nums.
	// Create all possible permutations of nums.
	// Return permutations.
	// Use backtracking, recursion, and early pruning.
	// Use a closure for recursion and state management.
	// Use a map for tracking visited numbers.

	// Check input min edge case.
	numLen := len(nums)
	if numLen <= 1 {
		return [][]int{nums}
	}

	// Initialize variables.
	var res [][]int
	cur := make([]int, 0, numLen)
	vst := make(map[int]bool)

	// Create a closure for recursion and state management.
	var bckPrm func()
	bckPrm = func() {
		// Check base case.
		// Check whole permutation length numLen created.
		if len(cur) == numLen {
			// Create copy of cur.
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res = append(res, tmp)
			return
		}

		// Recurse through each number.
		for _, num := range nums {
			// Check early pruning
			if vst[num] {
				continue
			}

			// Use current number.
			vst[num] = true
			cur = append(cur, num)

			// Recurse.
			bckPrm()

			// Backtrack current number.
			cur = cur[:len(cur)-1]
			vst[num] = false
		}
	}

	bckPrm()

	return res
}

// TestPermute46b validates the permute function
func TestPermute46b(t *testing.T) {
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
			got := permute46b(tc.input)
			// Sort both got and expected for consistent comparison
			if len(got) != len(tc.expected) {
				t.Errorf("got length %v, want length %v", len(got), len(tc.expected))
				return
			}

			// Create map to check if all permutations are present
			expMap := make(map[string]bool)
			for _, perm := range tc.expected {
				expMap[arrayToString46b(perm)] = true
			}

			// Verify each generated permutation is valid
			for _, perm := range got {
				if !expMap[arrayToString46b(perm)] {
					t.Errorf("unexpected permutation %v", perm)
				}
			}
		})
	}
}

// arrayToString46b creates string representation of array for comparison
func arrayToString46b(arr []int) string {
	return string(reflect.ValueOf(arr).String())
}
