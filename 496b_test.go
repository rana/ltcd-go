package main

import (
	"reflect"
	"testing"
)

// Time complexity:
// Space complexity:
func nextGreaterElement496b(nums1 []int, nums2 []int) []int {
	// Next Greater Element I

	return nil
}

func TestNextGreaterElement496b(t *testing.T) {
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
			got := nextGreaterElement496b(tc.nums1, tc.nums2)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("got %v, want %v", got, tc.expected)
			}
		})
	}
}
