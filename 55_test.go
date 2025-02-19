package main

import "testing"

// Time complexity: O(n), n is the length of the nums array.
// Space complexity: O(1), constant additional space used.
// https://console.anthropic.com/workbench/12bc0d21-11f2-4aff-837d-fc50a6bcc9c2
func canJump55(nums []int) bool {
	maxRch := 0 // Maximum reachable position
	numLen := len(nums)

	for idx := range numLen {
		// If current position is beyond our reach, we can't proceed
		if idx > maxRch {
			return false
		}

		// Update maximum reachable position
		curRch := idx + nums[idx]
		if curRch > maxRch {
			maxRch = curRch
		}

		// If we can reach the last index, return early
		if maxRch >= numLen-1 {
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
