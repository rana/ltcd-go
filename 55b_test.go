package main

import "testing"

func canJump55b(nums []int) bool {
	// Jump Game
	// Given an integer array nums.
	// Determine whether you can reach last index.
	// Return true if condition met.
	// Conditions:
	// * Array element is a maximum jump length.
	// * May jump less than max if it helps.
	// Use a local optimization "greedy" approach.

	maxRch := 0

	for idx, num := range nums {
		// Check fail condition.
		if idx > maxRch {
			return false
		}

		// Update max reach.
		curRch := idx + num
		if curRch > maxRch {
			maxRch = curRch
		}

		// Check success condition.
		if maxRch >= len(nums)-1 {
			return true
		}
	}

	return true
}

func TestCanJump55b(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "Example 1: Can reach end",
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			name: "Example 2: Cannot reach end",
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		},
		{
			name: "Single element",
			nums: []int{0},
			want: true,
		},
		{
			name: "All zeros except first",
			nums: []int{2, 0, 0},
			want: true,
		},
		{
			name: "Cannot cross zeros",
			nums: []int{1, 0, 1, 0},
			want: false,
		},
		{
			name: "Maximum jumps",
			nums: []int{5, 4, 3, 2, 1, 0},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump55b(tt.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
