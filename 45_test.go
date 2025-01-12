package main

import (
	"testing"
)

// Time complexity: O(n), n is the length of the array.
// Space complexity: O(1), constant additional space used.
func jump45(nums []int) int {
	// Edge case: single element array needs no jumps
	if len(nums) <= 1 {
		return 0
	}

	// Initialize variables
	jmp := 0    // number of jumps
	curMax := 0 // furthest index we can reach from current position
	nxtMax := 0 // furthest index we can reach in next jump

	// Iterate through array (except last element)
	for idx := 0; idx < len(nums)-1; idx++ {
		// Update furthest reachable position in next jump
		nxtMax = max(nxtMax, idx+nums[idx])

		// If we've reached the limit of our current jump
		if idx == curMax {
			// Make the jump
			jmp++
			// Update our current reach to the furthest we found
			curMax = nxtMax

			// If we can already reach the end, no need to continue
			if curMax >= len(nums)-1 {
				break
			}
		}
	}

	return jmp
}

func TestJump45(t *testing.T) {
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
			if got := jump45(tt.nums); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
