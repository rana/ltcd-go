package main

import (
	"testing"
)

func jump45b(nums []int) int {
	// Jump Game II
	// Given an integer array nums.
	// Determine min number of jumps to last index.
	// Return min number of jumps.
	// Conditions:
	// * Array element is max jump length.
	// * Can jump less than max.
	// Use a local optimization "greedy" approach.

	// Check min edge case.
	if len(nums) <= 1 {
		return 0
	}

	jmpCnt := 0
	curMax, nxtMax := 0, 0

	for idx := 0; idx < len(nums); idx++ {
		// Calculate nxtMax
		nxtMax = max(nxtMax, idx+nums[idx])

		// Check completion of current max jump.
		if idx == curMax {
			curMax = nxtMax
			jmpCnt++

			// Check early return
			if curMax >= len(nums)-1 {
				break
			}
		}
	}

	return jmpCnt
}

func TestJump45b(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{2, 3, 1, 1, 4},
			want: 2,
		},
		{
			name: "Example 2",
			nums: []int{2, 3, 0, 1, 4},
			want: 2,
		},
		{
			name: "Single element",
			nums: []int{0},
			want: 0,
		},
		{
			name: "Two elements",
			nums: []int{1, 1},
			want: 1,
		},
		{
			name: "Maximum jumps",
			nums: []int{1, 1, 1, 1},
			want: 3,
		},
		{
			name: "Large jumps",
			nums: []int{5, 9, 3, 2, 1, 0, 2, 3, 3, 1, 0, 0},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := jump45b(tt.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
