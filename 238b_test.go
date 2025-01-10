package main

import (
	"reflect"
	"testing"
)

func productExceptSelf238b(nums []int) []int {
	// Product of Array Except Self
	// Given and integer array nums.
	// At each element, calculate the product of the whole array except self.
	// Return the array of answers.
	// Conditions:
	// * Time complexity O(n)
	// * Do not use division
	// Use a two-pass approach.
	// Left to right pass product of prefixes.
	// Right to left pass product of suffixes and add.

	res := make([]int, len(nums))
	res[0] = 1

	// Left to right pass of product prefixes.
	for idx := 1; idx < len(nums); idx++ {
		res[idx] = res[idx-1] * nums[idx-1]
	}

	// Right to left pass pf prodcut suffixes and add.
	rhtPrd := 1
	for idx := len(nums) - 1; idx >= 0; idx-- {
		res[idx] *= rhtPrd
		rhtPrd *= nums[idx]
	}

	return res
}

func TestProductExceptSelf238b(t *testing.T) {
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
			result := productExceptSelf238b(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("productExceptSelf() = %v, want %v", result, tt.expected)
			}
		})
	}
}
