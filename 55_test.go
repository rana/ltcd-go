package main

import "testing"

// Time complexity: O(n), n is the length of the nums array.
// Space complexity: O(1), constant additional space used.
// https://console.anthropic.com/workbench/12bc0d21-11f2-4aff-837d-fc50a6bcc9c2
func canJump55(jmps []int) bool {
	// Jump Game
	// Given an integer array nums.
	// Determine whether you can reach last index.
	// Return true if condition met.
	// Conditions:
	// * Start at first index
	// * Element value is max jump value
	// * Can jump less than max if helpful
	// Use local optimization "greedy" approach.

	if len(jmps) == 1 {
		return true
	}
	maxJmp, lstIdx := 0, len(jmps)-1
	for idx, jmp := range jmps {
		if idx > maxJmp {
			return false
		}
		curJmp := idx + jmp
		if curJmp > maxJmp {
			maxJmp = curJmp
		}
		if maxJmp >= lstIdx {
			return true
		}
	}
	return true
}

func TestCanJump55(t *testing.T) {
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
			if got := canJump55(tt.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
