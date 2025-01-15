package main

import (
	"reflect"
	"testing"
)

// Time complexity:
// Space complexity:
func rotate189b(nums []int, k int) {

}

func TestRotate189b(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{1, 2, 3, 4, 5, 6, 7},
			k:        3,
			expected: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			name:     "Example 2",
			nums:     []int{-1, -100, 3, 99},
			k:        2,
			expected: []int{3, 99, -1, -100},
		},
		{
			name:     "K is zero",
			nums:     []int{1, 2, 3, 4, 5},
			k:        0,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "K equals array length",
			nums:     []int{1, 2, 3, 4, 5},
			k:        5,
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "K larger than array length",
			nums:     []int{1, 2, 3, 4, 5},
			k:        12,
			expected: []int{4, 5, 1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate189b(tt.nums, tt.k)
			if !reflect.DeepEqual(tt.nums, tt.expected) {
				t.Errorf("got %v, want %v", tt.nums, tt.expected)
			}
		})
	}
}
