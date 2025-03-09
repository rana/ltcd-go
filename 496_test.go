package main

import (
	"reflect"
	"testing"
)

// Time complexity: O(n + m) where n is the length of nums1 and m is the length of nums2. We process each element in nums2 once and then look up each element in nums1
// Space complexity: O(m) for the stack and hash map, where m is the length of nums2 https://console.anthropic.com/workbench/9967550a-b29d-49c7-ac0d-f90829f3a1f5
func nextGreaterElement496(nums1 []int, nums2 []int) []int {
	// Map to store next greater element for each number in nums2
	nxtGrt := make(map[int]int)
	// Monotonic decreasing stack
	var stk []int

	// Process nums2 to find next greater element for each number
	for _, num := range nums2 {
		// Pop elements from stack that have found their next greater element
		for len(stk) > 0 && stk[len(stk)-1] < num {
			// Current num is the next greater element for the popped element
			nxtGrt[stk[len(stk)-1]] = num
			stk = stk[:len(stk)-1]
		}
		// Push current number to stack
		stk = append(stk, num)
	}

	// Build result array for nums1
	res := make([]int, len(nums1))
	for i, num := range nums1 {
		// Look up next greater element in map, default to -1 if not found
		if val, exists := nxtGrt[num]; exists {
			res[i] = val
		} else {
			res[i] = -1
		}
	}

	return res
}

func TestNextGreaterElement496(t *testing.T) {
	cases := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected []int
	}{
		{
			name:     "Example 1",
			nums1:    []int{4, 1, 2},
			nums2:    []int{1, 3, 4, 2},
			expected: []int{-1, 3, -1},
		},
		{
			name:     "Example 2",
			nums1:    []int{2, 4},
			nums2:    []int{1, 2, 3, 4},
			expected: []int{3, -1},
		},
		{
			name:     "All decreasing",
			nums1:    []int{3, 2, 1},
			nums2:    []int{3, 2, 1},
			expected: []int{-1, -1, -1},
		},
		{
			name:     "All increasing",
			nums1:    []int{1, 2, 3},
			nums2:    []int{1, 2, 3, 4},
			expected: []int{2, 3, 4},
		},
		{
			name:     "Single element",
			nums1:    []int{5},
			nums2:    []int{5},
			expected: []int{-1},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := nextGreaterElement496(tc.nums1, tc.nums2)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}
