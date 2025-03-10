package main

import (
	"testing"
)

// Time complexity: O(n), n is the length of the array.
// Space complexity: O(1), constant additional space used.
func jump45(jmps []int) int {
	// Jump Game II
	// Given an integer array jmps.
	// Determine min jump count to last index.
	// Return min jump count.
	// Conditions:
	// * Start at first index
	// * Element value is max jump from index
	// * May jump less than max index if helpful
	// Use local optimization "greedy" strategy.
	//
	if len(jmps) == 1 {
		return 0
	}
	jmpCnt, lstIdx := 0, len(jmps)-1
	curMaxJmp, nxtMaxJmp := 0, 0
	// curMaxJmp: furthest index we can reach from current position
	// nxtMaxJmp: furthest index we can reach in next jump

	// Iterate through array (except last element)
	for idx, jmp := range jmps {
		// Update furthest reachable position in next jump
		nxtMaxJmp = max(nxtMaxJmp, idx+jmp)

		// If we've reached the limit of our current jump
		if idx == curMaxJmp {
			// Make the jump
			jmpCnt++
			// Update our current reach to the furthest we found
			curMaxJmp = nxtMaxJmp

			// If we can already reach the end, no need to continue
			if curMaxJmp >= lstIdx {
				break
			}
		}
	}

	return jmpCnt
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
