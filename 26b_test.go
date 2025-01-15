package main

import (
	"reflect"
	"testing"
)

// Time complexity:
// Space complexity:
func removeDuplicates26b(nums []int) int {
	return 0
}

func TestRemoveDuplicates26b(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
		k        int
	}{
		// Valid test cases within constraints
		{nums: []int{1, 1, 2}, expected: []int{1, 2}, k: 2},
		{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, expected: []int{0, 1, 2, 3, 4}, k: 5},
		{nums: []int{-100, -100, -99, -99, 0, 0, 1, 2, 2}, expected: []int{-100, -99, 0, 1, 2}, k: 5},
		{nums: []int{1}, expected: []int{1}, k: 1},
	}

	for _, test := range tests {
		numsCopy := make([]int, len(test.nums))
		copy(numsCopy, test.nums) // Preserve original nums for error messages

		k := removeDuplicates26b(test.nums)
		if k != test.k {
			t.Errorf("removeDuplicates(%v) returned k=%d; expected k=%d", numsCopy, k, test.k)
		}
		if !reflect.DeepEqual(test.nums[:k], test.expected) {
			t.Errorf("After removeDuplicates(%v), nums=%v; expected nums=%v", numsCopy, test.nums[:k], test.expected)
		}
	}
}
