package main

import (
	"reflect"
	"testing"
)

// Time complexity:
// Space complexity:
func twoSum(nums []int, tgt int) []int {
	return nil
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		tgt  int
		exp  []int
	}{
		{
			name: "basic case",
			nums: []int{2, 7, 11, 15},
			tgt:  9,
			exp:  []int{0, 1},
		},
		{
			name: "numbers not at start",
			nums: []int{3, 2, 4},
			tgt:  6,
			exp:  []int{1, 2},
		},
		{
			name: "same number twice",
			nums: []int{3, 3},
			tgt:  6,
			exp:  []int{0, 1},
		},
		{
			name: "negative numbers",
			nums: []int{-1, -2, -3, -4, -5},
			tgt:  -8,
			exp:  []int{2, 4},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			res := twoSum(tc.nums, tc.tgt)
			if !reflect.DeepEqual(res, tc.exp) {
				t.Errorf("twoSum(%v, %d) = %v; want %v",
					tc.nums, tc.tgt, res, tc.exp)
			}
		})
	}
}
