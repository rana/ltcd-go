package main

import (
	"reflect"
	"testing"
)

// Time complexity:
// Space complexity:
func twoSum167b(nums []int, tgt int) []int {
	return nil
}

// Test cases
func TestTwoSum167b(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "Example 1: Standard case",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{1, 2},
		},
		{
			name:     "Example 2: Three numbers",
			nums:     []int{2, 3, 4},
			target:   6,
			expected: []int{1, 3},
		},
		{
			name:     "Example 3: Negative numbers",
			nums:     []int{-1, 0},
			target:   -1,
			expected: []int{1, 2},
		},
		{
			name:     "Test 4: Larger numbers",
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			target:   13,
			expected: []int{6, 7},
		},
		{
			name:     "Test 5: Adjacent numbers",
			nums:     []int{1, 2, 3, 4},
			target:   3,
			expected: []int{1, 2},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := twoSum167b(tc.nums, tc.target)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("twoSum() = %v, want %v", got, tc.expected)
			}
		})
	}
}
