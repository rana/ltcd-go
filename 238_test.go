package main

import (
	"reflect"
	"testing"
)

// Time complexity: O(n), n is the length of array nums. We traverse the array twice.
// Space complexity: O(1), constant additional space used.
// https://chatgpt.com/c/671577be-d428-8002-8153-388aa27e3226
func productExceptSelf238(nums []int) []int {
	// Product of Array Except Self
	// Given an integer array.
	// Compute the product of adjacent elements.
	//	- Exclude self element
	//	- No division operation
	//	- Time complexity O(1)
	// Use a two-pass approach.
	// Pass one, compute left prefixes.
	// Pass two, compute right suffixes.

	// Initialize variables.
	res := make([]int, len(nums))
	res[0] = 1

	// lft to rht: compute left prefix products and store.
	for idx := 1; idx < len(nums); idx++ {
		res[idx] = res[idx-1] * nums[idx-1]
	}

	// rht to lft: compute right suffix products and store.
	rhtPrd := 1
	for idx := len(nums) - 1; idx >= 0; idx-- {
		res[idx] *= rhtPrd
		rhtPrd *= nums[idx]
	}

	return res
}

func TestProductExceptSelf238(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Basic positive numbers",
			input:    []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			name:     "Mixed numbers with zero",
			input:    []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
		{
			name:     "All zeros",
			input:    []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		{
			name:     "All negative numbers",
			input:    []int{-1, -2, -3, -4},
			expected: []int{-24, -12, -8, -6},
		},
		{
			name:     "Single element",
			input:    []int{5},
			expected: []int{1},
		},
		{
			name:     "Two elements",
			input:    []int{2, 3},
			expected: []int{3, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := productExceptSelf238(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("productExceptSelf() = %v, want %v", result, tt.expected)
			}
		})
	}
}
