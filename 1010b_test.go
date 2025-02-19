package main

import "testing"

// Time Complexity:
// Space Complexity:
func numPairsDivisibleBy60_1010b(durs []int) int {
	return 0
}

func TestNumPairsDivisibleBy60_1010b(t *testing.T) {
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
			if got := numPairsDivisibleBy60_1010b(tt.time); got != tt.want {
				t.Errorf("numPairsDivisibleBy60() = %v, want %v", got, tt.want)
			}
		})
	}
}
