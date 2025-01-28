package main

import "testing"

// Time complexity:
// Space complexity:
func containsNearbyDuplicate219b(nums []int, k int) bool {
	return false
}

func TestContainsNearbyDuplicate219b(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{
			name: "example 1: duplicate within range",
			nums: []int{1, 2, 3, 1},
			k:    3,
			want: true,
		},
		{
			name: "example 2: adjacent duplicates",
			nums: []int{1, 0, 1, 1},
			k:    1,
			want: true,
		},
		{
			name: "example 3: duplicates outside range",
			nums: []int{1, 2, 3, 1, 2, 3},
			k:    2,
			want: false,
		},
		{
			name: "single element",
			nums: []int{1},
			k:    1,
			want: false,
		},
		{
			name: "no duplicates",
			nums: []int{1, 2, 3, 4},
			k:    1,
			want: false,
		},
		{
			name: "k equals zero",
			nums: []int{1, 1, 1},
			k:    0,
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate219b(tt.nums, tt.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
