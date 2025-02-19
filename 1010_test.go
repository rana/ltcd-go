package main

import "testing"

// Time Complexity: O(n), where n is the length of the input array
// Space Complexity: O(1) as we store at most 60 different remainders in the map
// https://console.anthropic.com/workbench/87f50531-2567-4f8f-af95-fa64cd10a862
func numPairsDivisibleBy60_1010(time []int) int {
	// Pairs of Songs with Total Durations Divisible by 60
	// Given an integer array of song durations.
	// Determine quantity of duration sums divisible by 60.
	// Use frequency counter of remainder.
	// Complement needs modulo check.

	// Map to store frequency of remainders
	rem := make(map[int]int)
	// Count of valid pairs
	cnt := 0
	for _, dur := range time {
		// Get remainder when divided by 60
		cur := dur % 60
		// Find complement (what we need to get to 60 or 0)
		cmp := (60 - cur) % 60
		// Add pairs with complementary remainder
		cnt += rem[cmp]
		// Update frequency map
		rem[cur]++
	}
	return cnt
}

func TestNumPairsDivisibleBy60_1010(t *testing.T) {
	tests := []struct {
		name string
		time []int
		want int
	}{
		{
			name: "Example 1",
			time: []int{30, 20, 150, 100, 40},
			want: 3,
		},
		{
			name: "Example 2",
			time: []int{60, 60, 60},
			want: 3,
		},
		{
			name: "Empty array",
			time: []int{},
			want: 0,
		},
		{
			name: "Single element",
			time: []int{60},
			want: 0,
		},
		{
			name: "No valid pairs",
			time: []int{61, 62, 63},
			want: 0,
		},
		{
			name: "All same remainder",
			time: []int{30, 90, 150, 210},
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numPairsDivisibleBy60_1010(tt.time); got != tt.want {
				t.Errorf("numPairsDivisibleBy60() = %v, want %v", got, tt.want)
			}
		})
	}
}
